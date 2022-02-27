package controller

import (
	"net/http"

	"github.com/lcnssantos/iothub/cmd/iothub/internal/auth/dto"
	"github.com/lcnssantos/iothub/cmd/iothub/internal/auth/service"
	http2 "github.com/lcnssantos/iothub/internal/http"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c AuthController) Auth(w http.ResponseWriter, r *http.Request) {
	authRequest := &dto.AuthRequest{}

	if err := http2.HandleValidationRequest(w, r, authRequest); err != nil {
		return
	}

	auth, err := c.authService.Auth(*authRequest, r.Context())

	if err != nil {
		http2.ThrowHttpException(w, http.StatusUnauthorized, err.Error())
		return
	}

	jwtToken, err := c.authService.CreateAccessToken(*auth)

	if err != nil {
		http2.ThrowHttpException(w, http.StatusInternalServerError, err.Error())
		return
	}

	refreshToken, err := c.authService.CreateRefreshToken(*auth)

	if err != nil {
		http2.ThrowHttpException(w, http.StatusInternalServerError, err.Error())
		return
	}

	http2.SendHttpResponse(w, http.StatusOK, &dto.AuthResponse{AccessToken: jwtToken, RefreshToken: refreshToken, Type: "bearer"})
}

func (c AuthController) Refresh(w http.ResponseWriter, r *http.Request) {
	request := &dto.RefreshRequest{}

	if err := http2.HandleValidationRequest(w, r, request); err != nil {
		return
	}

	token, refreshToken, err := c.authService.RefreshToken(request.RefreshToken, r.Context())

	if err != nil {
		http2.ThrowHttpException(w, http.StatusUnauthorized, err.Error())
		return
	}

	http2.SendHttpResponse(w, http.StatusOK, &dto.AuthResponse{AccessToken: token, RefreshToken: refreshToken, Type: "bearer"})
}
