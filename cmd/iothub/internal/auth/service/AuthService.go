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

func (s AuthService) Auth(data dto.AuthRequest, ctx context.Context) (*dto2.User, error) {
	user, err := s.userService.FindOneByEmail(data.Email, ctx)

	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !user.Active {
		return nil, errors.New("User not active")
	}

	compare := s.hashService.Compare(user.Password, data.Password)

	if compare {
		return user, nil
	}

	return nil, errors.New("invalid credentials")
}

func (s AuthService) CreateRefreshToken(data dto2.User) (string, error) {
	return s.JwtService.Encode(data.Id, "refresh", 24*60*60)
}

func (s AuthService) CreateAccessToken(data dto2.User) (string, error) {
	return s.JwtService.Encode(data.Id, "token", 15*60)
}

func (s AuthService) GetByToken(token string, ctx context.Context) (*dto2.User, error) {
	claims, err := s.JwtService.Decode(token)

	if err != nil {
		return nil, err
	}

	if claims.Kind != "token" {
		return nil, errors.New("Invalid Token type")
	}

	return s.userService.FindOneById(claims.Id, ctx)
}

func (s AuthService) RefreshToken(refreshToken string, ctx context.Context) (string, string, error) {
	claims, err := s.JwtService.Decode(refreshToken)

	if err != nil {
		return "", "", err
	}

	if claims.Kind != "refresh" {
		return "", "", errors.New("Invalid Token type")
	}

	user, err := s.userService.FindOneById(claims.Id, ctx)

	if err != nil {
		return "", "", err
	}

	if !user.Active {
		return "", "", errors.New("User not active")
	}

	jwtToken, err := s.CreateAccessToken(*user)

	if err != nil {
		return "", "", errors.New("Error to create token")
	}

	refreshToken, err = s.CreateRefreshToken(*user)

	if err != nil {
		return "", "", errors.New("Error to create token")
	}

	return jwtToken, refreshToken, nil
}
