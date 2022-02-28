package service

import (
	"database/sql"

	"github.com/lcnssantos/iothub/cmd/publicApi/internal/rmq"

	"github.com/lcnssantos/iothub/cmd/publicApi/internal/accounts/dto"
	"github.com/lcnssantos/iothub/cmd/publicApi/internal/accounts/repository"
)

type AccountService struct {
	accountRepository *repository.AccountRepository
	rmqClient         *rmq.Client
}

func NewAccountService(accountRepository *repository.AccountRepository, rmqClient *rmq.Client) *AccountService {
	return &AccountService{accountRepository: accountRepository, rmqClient: rmqClient}
}

func (s AccountService) CreateAccount(data *dto.CreateAccountRequest, tx *sql.Tx) error {
	err := s.accountRepository.CreateAccount(data, tx)

	if err != nil {
		return err
	}

	err = s.rmqClient.CreateAccount(data.Login, data.Password)

	if err != nil {
		return err
	}

	return nil
}
