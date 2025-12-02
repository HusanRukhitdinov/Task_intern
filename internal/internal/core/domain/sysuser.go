package domain

import "errors"

type SysuserCreateRequest struct {
	Name     string   `json:"name" binding:"required"`
	Phone    string   `json:"phone" binding:"required"`
	Password string   `json:"password" binding:"required,min=6"`
	Roles    []string `json:"roles" binding:"required,min=1"`
}

type SysuserCreateResponse struct {
	ID string `json:"id"`
}
var (
	ErrSysuserAlreadyExists = errors.New("sysuser already exists")
	ErrRoleNotFound         = errors.New("role not found")
)
