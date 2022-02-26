package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/lcnssantos/gomusic/cmd/gomusic/internal/auth/service"
	http2 "github.com/lcnssantos/gomusic/internal/http"
)

type AuthenticationMiddleware struct {
	authService *service.AuthService
}

func NewAuthenticationMiddleware(authService *service.AuthService) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{authService: authService}
}

func (this AuthenticationMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")

		if authorization == "" {
			http2.ThrowHttpError(w, http.StatusUnauthorized, "Missing JWT Token")
			return
		}

		bearerToken := strings.Replace(authorization, "Bearer ", "", 1)

		user, err := this.authService.GetByToken(bearerToken, r.Context())

		if err != nil {
			http2.ThrowHttpError(w, http.StatusUnauthorized, err.Error())
			return
		}

		if !user.Active {
			http2.ThrowHttpError(w, http.StatusForbidden, "User not active")
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
