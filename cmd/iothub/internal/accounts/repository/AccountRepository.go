package repository

import (
	"context"
	"database/sql"

	"github.com/lcnssantos/iothub/cmd/iothub/internal/accounts/dto"
)

type AccountRepository struct {
	db *sql.DB
}

func (this AccountRepository) CreateAccount(data *dto.CreateAccountRequest, ctx context.Context) error {
	return nil
}

func NewAccountRepository(database *sql.DB) *AccountRepository {
	return &AccountRepository{db: database}
}
