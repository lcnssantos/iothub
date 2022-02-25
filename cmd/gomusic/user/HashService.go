package user

import "golang.org/x/crypto/bcrypt"

type HashService struct {
}

func (this *HashService) Hash(input string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(input), 14)
	return string(password), err
}

func (this *HashService) Compare(hashed string, input string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input))
	return err == nil
}

func NewHashService() *HashService {
	return &HashService{}
}
