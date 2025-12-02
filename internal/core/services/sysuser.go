package services

import (
	"context"
	"database/sql"
	"intern/internal/core/domain"
	"intern/internal/core/repository/psql/sqlc"
	"time"

	"gopkg.in/guregu/null.v4/zero"
)

func (s *Service) CreateSysuser(ctx context.Context, req *domain.SysuserCreateRequest, hashedPassword string, createdBy string) (*domain.SysuserCreateResponse, error) {
	existingSysusers, err := s.storage.GetSysuserByPhone(ctx, sqlc.GetSysuserByPhoneParams{
		Status: "active",
		Phone:  req.Phone,
	})
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if len(existingSysusers) > 0 {
		return nil, domain.ErrSysuserAlreadyExists
	}
	for _, roleID := range req.Roles {
		_, err := s.storage.GetRoleById(ctx, sqlc.GetRoleByIdParams{
			ID:     roleID,
			Status: "active",
		})
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, domain.ErrRoleNotFound
			}
			return nil, err
		}
	}

	sysuserID, err := s.storage.CreateSysuser(ctx, sqlc.CreateSysuserParams{
		Status:    "active",
		Name:      req.Name,
		Phone:     req.Phone,
		Password:  hashedPassword,
		CreatedAt: zero.TimeFrom(time.Now()),
		CreatedBy: zero.StringFrom(createdBy),
	})
	if err != nil {
		return nil, err
	}

	for _, roleID := range req.Roles {
		err = s.storage.CreateSysRoles(ctx, sqlc.CreateSysRolesParams{
			SysuserID: sysuserID,
			RoleID:    roleID,
		})
		if err != nil {
			return nil, err
		}
	}

	return &domain.SysuserCreateResponse{
		ID: sysuserID,
	}, nil
}
