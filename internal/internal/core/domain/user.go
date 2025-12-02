package domain

import (
	"time"
)

type UserParams struct {
	ID    string `json:"-"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type SelectOneUser struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SelectManyUsers struct {
	Users []SelectOneUser `json:"users"`
	Count int64           `json:"count"`
}

type ResponseSuccess struct {
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
}

type ResponseError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type IDRequest struct {
	ID string `json:"id" binding:"required"`
}
