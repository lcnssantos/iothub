package router

import (
	"github.com/gorilla/mux"
	"github.com/lcnssantos/gomusic/cmd/gomusic/internal/auth/middleware"
	"github.com/lcnssantos/gomusic/cmd/gomusic/internal/user/controller"
)

func BuildMeRouter(controller *controller.MeController, router *mux.Router, middleware *middleware.AuthenticationMiddleware) {
	router.Use(middleware.Handler)
	router.Methods("GET").HandlerFunc(controller.GetMe)
}
