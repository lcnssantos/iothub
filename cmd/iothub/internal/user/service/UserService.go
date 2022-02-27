package service

import (
	"context"
	"errors"

	"github.com/lcnssantos/iothub/internal/seed"

	dto2 "github.com/lcnssantos/iothub/cmd/iothub/internal/accounts/dto"
	"github.com/lcnssantos/iothub/cmd/iothub/internal/accounts/service"
	"github.com/lcnssantos/iothub/cmd/iothub/internal/user/dto"
	"github.com/lcnssantos/iothub/cmd/iothub/internal/user/repository"
)

type UserService struct {
	repository     *repository.UserRepository
	hashService    *HashService
	accountService *service.AccountService
}

func (s UserService) Create(data dto.CreateUserDto, ctx context.Context) error {
	_, err := s.repository.FindOneByEmail(data.Email, ctx)

	if err == nil {
		return errors.New("email already exist")
	}

	tx, err := s.repository.GetTransaction(ctx)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	hash, err := s.hashService.Hash(data.Password)

	if err != nil {
		return err
	}

	data.Password = hash

	if err := s.repository.Create(data, tx); err != nil {
		return err
	}

	err = s.accountService.CreateAccount(&dto2.CreateAccountRequest{Login: seed.String(24), Password: seed.String(24), Email: data.Email}, tx)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s UserService) FindOneByEmail(email string, ctx context.Context) (*dto.User, error) {
	return s.repository.FindOneByEmail(email, ctx)
}

func (s UserService) FindOneById(id uint64, ctx context.Context) (*dto.User, error) {
	return s.repository.FindOneById(id, ctx)
}

func (s UserService) List(ctx context.Context) ([]*dto.User, error) {
	return s.repository.List(ctx)
}

func NewUserService(repository *repository.UserRepository, hashService *HashService, accountService *service.AccountService) *UserService {
	return &UserService{repository: repository, hashService: hashService, accountService: accountService}
}
