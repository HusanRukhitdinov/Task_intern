package domain

import (
	"errors"
	"time"
)

type SignupRequest struct {
	OTPConfirmationToken string `json:"otp_confirmation_token" binding:"required"`
	Email                string `json:"email"                binding:"required,email"`
	Password             string `json:"password"             binding:"required,min=6"`
	Name                 string `json:"name"                 binding:"required"`
}

type SignupResponse struct {
	Message string `json:"message"`
}
type LoginRequest struct {
	UserType string `json:"user_type" binding:"required,oneof=user sysuser"`
	Email    string `json:"email"    binding:"omitempty,email"`
	Phone    string `json:"phone"    binding:"omitempty"`
	Password string `json:"password" binding:"required"`
}
type OTPTokenClaims struct {
	OTPID string `json:"otp_id"`
	Exp   int64  `json:"exp"`
}

func (c *OTPTokenClaims) Valid() error {
	if time.Now().Unix() > c.Exp {
		return errors.New("token expired")
	}
	return nil
}

var (
	ErrUnavailableOtpConfirmation = errors.New("otp confirmation token not found or mismatched")
	ErrInvalidCredentials         = errors.New("invalid credentials")
)
