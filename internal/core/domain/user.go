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
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SelectManyUsers struct {
	Users []SelectOneUser `json:"users"`
	Count int64           `json:"count"`
}
