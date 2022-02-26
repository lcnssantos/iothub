package service

import (
	"context"

	"github.com/lcnssantos/iothub/cmd/iothub/internal/accounts/dto"
	"github.com/lcnssantos/iothub/cmd/iothub/internal/accounts/repository"
)

type AccountService struct {
	accountRepository *repository.AccountRepository
}

func NewAccountService(repository *repository.AccountRepository) *AccountService {
	return &AccountService{accountRepository: repository}
}

func (this AccountService) CreateAccount(data *dto.CreateAccountRequest, ctx context.Context) error {
	return this.accountRepository.CreateAccount(data, ctx)
}
