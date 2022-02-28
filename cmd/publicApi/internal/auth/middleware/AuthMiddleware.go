package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/lcnssantos/iothub/cmd/publicApi/internal/auth/service"
	http2 "github.com/lcnssantos/iothub/internal/http"
)

type AuthenticationMiddleware struct {
	authService *service.AuthService
}

func NewAuthenticationMiddleware(authService *service.AuthService) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{authService: authService}
}

func (m AuthenticationMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")

		if authorization == "" {
			http2.ThrowHttpException(w, http.StatusUnauthorized, "Missing JWT Token")
			return
		}

		bearerToken := strings.Replace(authorization, "Bearer ", "", 1)

		user, err := m.authService.GetByToken(bearerToken, r.Context())

		if err != nil {
			http2.ThrowHttpException(w, http.StatusUnauthorized, err.Error())
			return
		}

		if !user.Active {
			http2.ThrowHttpException(w, http.StatusForbidden, "User not active")
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
