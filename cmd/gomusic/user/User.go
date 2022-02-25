package user

import "time"

type User struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"-"`
	Email     string    `json:"email"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateUserDto struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
	Email    string `json:"email" validate:"required,email"`
}
