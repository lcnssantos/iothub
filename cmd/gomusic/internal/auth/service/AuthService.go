package service

import (
	"context"
	"errors"

	"github.com/lcnssantos/gomusic/cmd/gomusic/internal/auth/dto"
	dto2 "github.com/lcnssantos/gomusic/cmd/gomusic/internal/user/dto"
	"github.com/lcnssantos/gomusic/cmd/gomusic/internal/user/service"
)

type AuthService struct {
	userService *service.UserService
	hashService *service.HashService
	JwtService  *JwtService
}

func NewAuthService(userService *service.UserService, hashService *service.HashService, jwtService *JwtService) *AuthService {
	return &AuthService{userService: userService, hashService: hashService, JwtService: jwtService}
}

func (this AuthService) Auth(data dto.AuthRequest, ctx context.Context) (*dto2.User, error) {
	user, err := this.userService.FindOneByEmail(data.Email, ctx)

	if err != nil {
		return nil, err
	}

	compare := this.hashService.Compare(user.Password, data.Password)

	if compare {
		return user, nil
	}

	return nil, errors.New("invalid credentials")
}

func (this AuthService) CreateRefreshToken(data dto2.User) (string, error) {
	return this.JwtService.Encode(data.Id, "refresh", 24*60*60)
}

func (this AuthService) CreateAccessToken(data dto2.User) (string, error) {
	return this.JwtService.Encode(data.Id, "token", 15*60)
}
