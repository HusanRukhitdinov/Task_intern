package domain

import (
	"time"
)

type RoleCreateRequest struct {
	Name string `json:"name" binding:"required"`
}
type RoleResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type RoleListResponse struct {
	Roles []RoleResponse `json:"roles"`
	Count int64          `json:"count"`
}
