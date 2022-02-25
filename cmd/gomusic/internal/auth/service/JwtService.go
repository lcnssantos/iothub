package service

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/lcnssantos/gomusic/config"
)

type JwtService struct {
}

func NewJwtService() *JwtService {
	return &JwtService{}
}

func (this JwtService) Encode(id string, kind string, expirationTime int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":  id,
		"kind": kind,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Second * time.Duration(expirationTime)).Unix(),
	})

	return token.SignedString([]byte(config.Get().JWT_KEY))
}
