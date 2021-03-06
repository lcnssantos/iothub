package service

import (
	"github.com/lcnssantos/iothub/cmd/publicApi/config"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenClaims struct {
	Id   uint64 `json:"id"`
	Kind string `json:"kind"`
	jwt.StandardClaims
}
type JwtService struct {
}

func NewJwtService() *JwtService {
	return &JwtService{}
}

func (s JwtService) Encode(id uint64, kind string, expirationTime int) (string, error) {
	tokenClaims := TokenClaims{
		id, kind,
		jwt.StandardClaims{
			Issuer:    "iothub",
			Subject:   "iothub",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Second * time.Duration(expirationTime)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	return token.SignedString([]byte(config.Get().JWT_KEY))
}

func (s JwtService) Decode(inputToken string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(inputToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Get().JWT_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*TokenClaims)

	return claims, nil
}
