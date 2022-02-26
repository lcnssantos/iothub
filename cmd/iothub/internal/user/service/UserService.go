package service

import (
	"context"
	"errors"

	"github.com/lcnssantos/iothub/cmd/iothub/internal/user/dto"
	"github.com/lcnssantos/iothub/cmd/iothub/internal/user/repository"
)

type UserService struct {
	repository  *repository.UserRepository
	hashService *HashService
}

func (this UserService) Create(data dto.CreateUserDto, ctx context.Context) error {
	_, err := this.repository.FindOneByEmail(data.Email, ctx)

	if err == nil {
		return errors.New("Email already exist")
	}

	hash, err := this.hashService.Hash(data.Password)

	if err != nil {
		return err
	}

	data.Password = hash

	return this.repository.Create(data, ctx)
}

func (this UserService) FindOneByEmail(email string, ctx context.Context) (*dto.User, error) {
	return this.repository.FindOneByEmail(email, ctx)
}

func (this UserService) FindOneById(id uint64, ctx context.Context) (*dto.User, error) {
	return this.repository.FindOneById(id, ctx)
}

func (this UserService) List(ctx context.Context) ([]*dto.User, error) {
	return this.repository.List(ctx)
}

func NewUserService(repository *repository.UserRepository, hashService *HashService) *UserService {
	return &UserService{repository: repository, hashService: hashService}
}
