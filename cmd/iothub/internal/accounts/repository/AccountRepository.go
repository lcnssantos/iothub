package repository

import (
	"database/sql"

	"github.com/lcnssantos/iothub/cmd/iothub/internal/accounts/dto"
)

type AccountRepository struct {
	db *sql.DB
}

func (r AccountRepository) CreateAccount(data *dto.CreateAccountRequest, tx *sql.Tx) error {
	prepare, err := tx.Prepare("INSERT INTO accounts (login, password, userId) values (?, ?, (SELECT id from users WHERE email = ?))")

	if err != nil {
		return err
	}

	_, err = prepare.Exec(data.Login, data.Password, data.Email)

	if err != nil {
		return err
	}

	return nil

}

func NewAccountRepository(database *sql.DB) *AccountRepository {
	return &AccountRepository{db: database}
}
