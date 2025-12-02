package services

import (
	"context"
	"intern/internal/core/domain"
	"intern/internal/core/repository/psql/sqlc"
)

func (s *Service) CreateRole(ctx context.Context, req *domain.RoleCreateRequest) error {
	params := sqlc.CreateRoleParams{
		Name:   req.Name,
		Status: "active",
	}
	return s.storage.CreateRole(ctx, params)
}

func (s *Service) GetRoleList(ctx context.Context) (*domain.RoleListResponse, error) {
	rows, err := s.storage.RoleList(ctx, "active")
	if err != nil {
		return nil, err
	}

	var resp domain.RoleListResponse
	for _, r := range rows {
		resp.Roles = append(resp.Roles, domain.RoleResponse{ID: r.ID, Name: r.Name, CreatedAt: r.CreatedAt.Time})
	}
	resp.Count = int64(len(resp.Roles))
	return &resp, nil
}
func (s *Service) UpdateRole(ctx context.Context, roleID string, name string) error {
	params := sqlc.UpdateRoleParams{ID: roleID, Name: name}
	return s.storage.UpdateRole(ctx, params)
}
