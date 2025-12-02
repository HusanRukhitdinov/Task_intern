package services

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"intern/internal/configs"
	"intern/internal/core/domain"
	"intern/internal/core/repository/psql/sqlc"
	token "intern/pkg/jwt"
	"intern/pkg/logger"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/guregu/null.v4/zero"
)

func (s *Service) SendOTP(ctx context.Context, req *domain.OTPCreateRequest) (*domain.OTPResponse, error) {
	code, err := token.Generate6DigitCode()
	if err != nil {
		return nil, err
	}

	params := sqlc.InsertOTPParams{
		Email:     req.Email,
		Code:      code,
		Status:    "unconfirmed",
		ExpiresAt: zero.TimeFrom(time.Now().Add(3 * time.Minute)),
	}

	id, err := s.storage.InsertOTP(ctx, params)
	if err != nil {
		return nil, err
	}

	err = s.emailSender.SendOTP(req.Email, code)
	if err != nil {
		s.log.Error("Failed to send OTP email", logger.Error(err))
	}

	s.log.Info("OTP Sent", logger.String("email", req.Email), logger.String("code", code))

	return &domain.OTPResponse{OTPID: id}, nil
}

func (s *Service) ConfirmOTP(ctx context.Context, req *domain.OTPConfirmRequest) (*domain.OTPConfirmationTokenResponse, error) {
	otp, err := s.storage.GetOTPByID(ctx, sqlc.GetOTPByIDParams{
		ID:     req.OTPID,
		Status: "unconfirmed",
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUnavailableOtpConfirmation
		}
		return nil, err
	}

	if otp.ExpiresAt.Valid && time.Now().After(otp.ExpiresAt.Time) {
		return nil, domain.ErrOTPConfirmationTokenExpired
	}

	if otp.Code != req.Code {
		return nil, domain.ErrIncorrectOTP
	}

	if err = s.storage.UpdateOTPStatus(ctx, sqlc.UpdateOTPStatusParams{
		Status: "confirmed",
		ID:     otp.ID,
	}); err != nil {
		return nil, err
	}

	otpToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &domain.OTPTokenClaims{
		OTPID: otp.ID,
		Exp:   time.Now().Add(15 * time.Minute).Unix(),
	})

	signedToken, err := otpToken.SignedString([]byte(configs.SignKey))
	if err != nil {
		return nil, err
	}

	return &domain.OTPConfirmationTokenResponse{Token: signedToken}, nil
}

func (s *Service) Signup(ctx context.Context, req *domain.SignupRequest) (*domain.SignupResponse, error) {
	claims, err := token.DecodeOTPToken(req.OTPConfirmationToken)
	if err != nil {
		return nil, err
	}

	if time.Now().Unix() > claims.Exp {
		return nil, domain.ErrOTPConfirmationTokenExpired
	}

	otpConfirmed, err := s.storage.GetOTPByID(ctx, sqlc.GetOTPByIDParams{
		ID:     claims.OTPID,
		Status: "confirmed",
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUnavailableOtpConfirmation
		}
		return nil, err
	}

	if otpConfirmed.Email != req.Email {
		return nil, domain.ErrUnavailableOtpConfirmation
	}

	_, err = s.storage.GetUserByEmail(ctx, sqlc.GetUserByEmailParams{
		Email:  req.Email,
		Status: "active",
	})
	if err == nil {
		return nil, domain.ErrUserAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	err = s.storage.InsertUser(ctx, sqlc.InsertUserParams{
		Email:    req.Email,
		Name:     req.Name,
		Password: string(hashedPassword),
		Status:   "active",
	})
	if err != nil {
		return nil, err
	}

	return &domain.SignupResponse{Message: "signup successful"}, nil
}

func (s *Service) Login(ctx context.Context, req *domain.LoginRequest) (*domain.TokenResponse, error) {
	var (
		id       string
		password string
		role     string
	)

	if req.UserType == "sysuser" {
		users, err := s.storage.GetSysuserByPhone(ctx, sqlc.GetSysuserByPhoneParams{
			Status: "active",
			Phone:  req.Phone,
		})
		if err != nil {
			return nil, err
		}
		if len(users) == 0 {
			return nil, domain.ErrInvalidCredentials
		}
		user := users[0]

		id = user.ID
		password = user.Password
		role = "sysuser"
	} else {
		user, err := s.storage.GetUserByEmail(ctx, sqlc.GetUserByEmailParams{
			Email:  req.Email,
			Status: "active",
		})
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, domain.ErrInvalidCredentials
			}
			return nil, err
		}
		id = user.ID
		password = user.Password
		role = "user"
	}

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(req.Password)); err != nil {
		return nil, domain.ErrInvalidCredentials
	}

	return token.GenerateJWTToken(&domain.TokenRequest{
		ID:   id,
		Role: role,
	}, s.log)
}
