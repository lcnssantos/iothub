package repository

import (
	"context"
	"database/sql"
	"github.com/lcnssantos/iothub/internal/database"

	"github.com/lcnssantos/iothub/cmd/iothub/internal/accounts/dto"
)

type AccountRepository struct {
	db *sql.DB
}

func (this AccountRepository) CreateAccount(data *dto.CreateAccountRequest, ctx context.Context) error {
	return database.ExecuteTransaction(ctx, this.db, func(tx *sql.Tx) error {
		prepare, err := tx.Prepare("INSERT INTO accounts (login, password, userId) values (?, ?, ?)")

		if err != nil {
			return err
		}

		_, err = prepare.Exec(data.Login, data.Password, data.UserId)

		if err != nil {
			return err
		}

		return nil
	})
}

func NewAccountRepository(database *sql.DB) *AccountRepository {
	return &AccountRepository{db: database}
}
