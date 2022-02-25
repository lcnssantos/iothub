package middlewares

import (
	"net/http"
)

type JsonMiddleware struct {
}

func NewJsonMiddleware() *JsonMiddleware {
	return &JsonMiddleware{}
}

func (this *JsonMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
