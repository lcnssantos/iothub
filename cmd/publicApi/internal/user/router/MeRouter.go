package router

import (
	"github.com/gorilla/mux"
	"github.com/lcnssantos/iothub/cmd/publicApi/internal/auth/middleware"
	"github.com/lcnssantos/iothub/cmd/publicApi/internal/user/controller"
)

func BuildMeRouter(controller *controller.MeController, router *mux.Router, middleware *middleware.AuthenticationMiddleware) {
	router.Use(middleware.Handler)
	router.Methods("GET").HandlerFunc(controller.GetMe)
}
