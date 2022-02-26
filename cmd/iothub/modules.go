package main

import (
	"database/sql"

	controller2 "github.com/lcnssantos/iothub/cmd/iothub/internal/auth/controller"
	"github.com/lcnssantos/iothub/cmd/iothub/internal/auth/middleware"
	router2 "github.com/lcnssantos/iothub/cmd/iothub/internal/auth/router"
	"github.com/lcnssantos/iothub/cmd/iothub/internal/auth/service"

	"github.com/gorilla/mux"
	"github.com/lcnssantos/iothub/cmd/iothub/internal/user/controller"
	"github.com/lcnssantos/iothub/cmd/iothub/internal/user/repository"
	"github.com/lcnssantos/iothub/cmd/iothub/internal/user/router"
	service2 "github.com/lcnssantos/iothub/cmd/iothub/internal/user/service"
)

func BuildUserModule(db *sql.DB, r *mux.Router) {
	userRepository := repository.NewUserRepository(db)
	hashService := service2.NewHashService()
	userService := service2.NewUserService(userRepository, hashService)
	userController := controller.NewUserController(userService)
	meController := controller.NewMeController()
	jwtService := service.NewJwtService()
	authService := service.NewAuthService(userService, hashService, jwtService)
	authMiddleware := middleware.NewAuthenticationMiddleware(authService)

	router.BuildUserRouter(userController, r.PathPrefix("/user").Subrouter())
	router.BuildMeRouter(meController, r.PathPrefix("/me").Subrouter(), authMiddleware)
}

func BuildAuthModule(db *sql.DB, r *mux.Router) {
	userRepository := repository.NewUserRepository(db)
	hashService := service2.NewHashService()
	userService := service2.NewUserService(userRepository, hashService)
	jwtService := service.NewJwtService()
	authService := service.NewAuthService(userService, hashService, jwtService)
	authController := controller2.NewAuthController(authService)
	router2.BuildRouter(authController, r)
}
