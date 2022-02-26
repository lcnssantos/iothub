package controller

import (
	"net/http"

	"github.com/lcnssantos/gomusic/cmd/gomusic/internal/auth/dto"
	"github.com/lcnssantos/gomusic/cmd/gomusic/internal/auth/service"
	http2 "github.com/lcnssantos/gomusic/internal/http"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (this AuthController) Auth(w http.ResponseWriter, r *http.Request) {
	authRequest := &dto.AuthRequest{}

	if err := http2.HandleValidationRequest(w, r, authRequest); err != nil {
		return
	}

	auth, err := this.authService.Auth(*authRequest, r.Context())

	if err != nil {
		http2.ThrowHttpError(w, http.StatusUnauthorized, err.Error())
		return
	}

	jwtToken, err := this.authService.CreateAccessToken(*auth)

	if err != nil {
		http2.ThrowHttpError(w, http.StatusInternalServerError, err.Error())
		return
	}

	refreshToken, err := this.authService.CreateRefreshToken(*auth)

	if err != nil {
		http2.ThrowHttpError(w, http.StatusInternalServerError, err.Error())
		return
	}

	http2.SetResponse(w, http.StatusOK, &dto.AuthResponse{AccessToken: jwtToken, RefreshToken: refreshToken, Type: "bearer"})
}

func (this AuthController) Refresh(w http.ResponseWriter, r *http.Request) {
	request := &dto.RefreshRequest{}

	if err := http2.HandleValidationRequest(w, r, request); err != nil {
		return
	}

	token, refreshToken, err := this.authService.RefreshToken(request.RefreshToken, r.Context())

	if err != nil {
		http2.ThrowHttpError(w, http.StatusUnauthorized, err.Error())
		return
	}

	http2.SetResponse(w, http.StatusOK, &dto.AuthResponse{AccessToken: token, RefreshToken: refreshToken, Type: "bearer"})
}
