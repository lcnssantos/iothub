package user

import "errors"

type Service struct {
	repository  *Repository
	hashService *HashService
}

func (this Service) Create(data CreateUserDto) error {
	_, err := this.repository.FindOneByEmail(data.Email)

	if err == nil {
		return errors.New("Email already exist")
	}

	hash, err := this.hashService.Hash(data.Password)

	if err != nil {
		return err
	}

	data.Password = hash

	return this.repository.Create(data)
}

func (this Service) FindOneByEmail(email string) (*User, error) {
	return this.repository.FindOneByEmail(email)
}

func (this Service) FindOneById(uid string) (*User, error) {
	return this.repository.FindOneById(uid)
}

func NewService(repository *Repository, hashService *HashService) *Service {
	return &Service{repository: repository, hashService: hashService}
}
