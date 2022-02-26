package service

import (
	"context"
	"github.com/lcnssantos/iothub/cmd/iothub/internal/rmq"

	"github.com/lcnssantos/iothub/cmd/iothub/internal/accounts/dto"
	"github.com/lcnssantos/iothub/cmd/iothub/internal/accounts/repository"
)

type AccountService struct {
	accountRepository *repository.AccountRepository
	rmqClient         *rmq.RMQClient
}

func NewAccountService(accountRepository *repository.AccountRepository, rmqClient *rmq.RMQClient) *AccountService {
	return &AccountService{accountRepository: accountRepository, rmqClient: rmqClient}
}

func (this AccountService) CreateAccount(data *dto.CreateAccountRequest, ctx context.Context) error {
	err := this.accountRepository.CreateAccount(data, ctx)

	if err != nil {
		return err
	}

	err = this.rmqClient.CreateAccount(data.Login, data.Password)

	if err != nil {
		return err
	}

	return nil
}
