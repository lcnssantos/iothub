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
	Vhost  string `json:"vhost"`
	UserId uint64 `json:"userId"`
}
