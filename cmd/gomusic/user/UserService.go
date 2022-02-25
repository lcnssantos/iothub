package user

import (
	"context"
	"errors"
)

type Service struct {
	repository  *Repository
	hashService *HashService
}

func (this Service) Create(data CreateUserDto, ctx context.Context) error {
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

func (this Service) FindOneByEmail(email string, ctx context.Context) (*User, error) {
	return this.repository.FindOneByEmail(email, ctx)
}

func (this Service) FindOneById(uid string, ctx context.Context) (*User, error) {
	return this.repository.FindOneById(uid, ctx)
}

func NewService(repository *Repository, hashService *HashService) *Service {
	return &Service{repository: repository, hashService: hashService}
}
