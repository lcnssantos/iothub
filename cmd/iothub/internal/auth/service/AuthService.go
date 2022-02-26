package service

import (
	"context"
	"errors"

	"github.com/lcnssantos/iothub/cmd/iothub/internal/auth/dto"
	dto2 "github.com/lcnssantos/iothub/cmd/iothub/internal/user/dto"
	"github.com/lcnssantos/iothub/cmd/iothub/internal/user/service"
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

	if !user.Active {
		return nil, errors.New("User not active")
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

func (this AuthService) GetByToken(token string, ctx context.Context) (*dto2.User, error) {
	claims, err := this.JwtService.Decode(token)

	if err != nil {
		return nil, err
	}

	if claims.Kind != "token" {
		return nil, errors.New("Invalid Token type")
	}

	return this.userService.FindOneById(claims.Uid, ctx)
}

func (this AuthService) RefreshToken(refreshToken string, ctx context.Context) (string, string, error) {
	claims, err := this.JwtService.Decode(refreshToken)

	if err != nil {
		return "", "", err
	}

	if claims.Kind != "refresh" {
		return "", "", errors.New("Invalid Token type")
	}

	user, err := this.userService.FindOneById(claims.Uid, ctx)

	if err != nil {
		return "", "", err
	}

	if !user.Active {
		return "", "", errors.New("User not active")
	}

	jwtToken, err := this.CreateAccessToken(*user)

	if err != nil {
		return "", "", errors.New("Error to create token")
	}

	refreshToken, err = this.CreateRefreshToken(*user)

	if err != nil {
		return "", "", errors.New("Error to create token")
	}

	return jwtToken, refreshToken, nil
}
