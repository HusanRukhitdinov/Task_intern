package services

import (
	"context"
	"intern/internal/core/domain"
	"intern/internal/core/repository/psql/sqlc"
	"intern/pkg/logger"
	"intern/pkg/serialize"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/cast"
	"gopkg.in/guregu/null.v4/zero"
)

func (auth *Service) DeleteOneUser(ctx context.Context, arg *domain.IDRequest) error {
	err := auth.storage.DeleteOneUser(ctx, arg.ID)
	if err != nil {
		auth.log.Error("this error can used in deleted on user storage", logger.Error(err))
		return err
	}
	return nil
}

func (auth *Service) EditOnePassword(ctx context.Context, arg *domain.LoginRequest) error {
	err := auth.storage.EditOnePassword(ctx, sqlc.EditOnePasswordParams{
		PasswordHash: arg.HashPassword,
		Email:        arg.Email,
	})
	if err != nil {
		auth.log.Error("this error can used in edit password on user storage", logger.Error(err))

		return err
	}
	return nil
}
func (auth *Service) EditOneUser(ctx context.Context, arg *domain.UserParams) error {
	err := auth.storage.EditOneUser(ctx, sqlc.EditOneUserParams{
		Name:      arg.Name,
		Email:     arg.Email,
		UpdatedAt: zero.TimeFrom(time.Now()),
		ID:        arg.ID,
	})
	if err != nil {
		auth.log.Error("this error can used in edit user on user storage", logger.Error(err))

		return err
	}
	return nil
}
func (auth *Service) LoginOneUser(ctx context.Context, arg *domain.LoginRequest) (*domain.IDResponse, error) {
	resp, err := auth.storage.LoginOneUser(ctx, sqlc.LoginOneUserParams{
		Email:        arg.Email,
		PasswordHash: arg.HashPassword,
	})
	if err != nil {
		auth.log.Error("this error can used in login user on user storage", logger.Error(err))

		return nil, err
	}
	return &domain.IDResponse{ID: resp}, nil
}
func (auth *Service) RegisterUser(ctx context.Context, arg *domain.RegisterRequest) error {
	err := auth.storage.RegisterUser(ctx, sqlc.RegisterUserParams{
		ID:           uuid.NewString(),
		Email:        arg.Email,
		PasswordHash: arg.HashPassword,
		CreatedAt:    zero.TimeFrom(time.Now()),
	})
	if err != nil {
		auth.log.Error("this error can used in register user on user storage", logger.Error(err))

		return err
	}
	return nil
}
func (auth *Service) SelectManyUsers(ctx context.Context) (*domain.SelectManyUsers, error) {
	users, err := auth.storage.SelectManyUsers(ctx)
	if err != nil {
		auth.log.Error("this error can used in get users on user storage", logger.Error(err))

		return nil, err
	}
	var resp []domain.SelectOneUser
	err = serialize.MarshalUnMarshal(users, &resp)
	if err != nil {
		auth.log.Error("this error can be used on marshal to struct ", logger.Error(err))
		return nil, err
	}
	return &domain.SelectManyUsers{
		Users: resp,
		Count: int64(len(resp)),
	}, nil
}
func (auth *Service) SelectOneUser(ctx context.Context, arg *domain.IDRequest) (*domain.SelectOneUser, error) {
	user, err := auth.storage.SelectOneUser(ctx, arg.ID)
	if err != nil {
		auth.log.Error("this error can used in get user on user storage", logger.Error(err))

		return nil, err
	}
	return &domain.SelectOneUser{
		ID:        user.ID,
		Name:      cast.ToString(user.Name),
		Email:     user.Email,
		CreatedAt: cast.ToTime(user.CreatedAt),
		UpdatedAt: cast.ToTime(user.UpdatedAt),
	}, nil
}
func (auth *Service) SelectOneUserEmail(ctx context.Context, arg *domain.EmailRequest) (*domain.PasswordResp, error) {
	email, err := auth.storage.SelectOneUserEmail(ctx, arg.Email)
	if err != nil {
		auth.log.Error("this error can used in get user by email on user storage", logger.Error(err))

		return nil, err
	}
	return &domain.PasswordResp{
		Password: email.PasswordHash,
		ID:       email.ID,
	}, nil
}
