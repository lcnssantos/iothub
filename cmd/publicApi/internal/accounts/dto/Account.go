package dto

import (
	"os/user"
	"time"
)

type Account struct {
	Id        uint64    `json:"id"`
	Vhost     string    `json:"vhost"`
	User      user.User `json:"user"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateAccountRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
