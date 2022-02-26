package main

import (
	"database/sql"

	controller2 "github.com/lcnssantos/gomusic/cmd/gomusic/internal/auth/controller"
	"github.com/lcnssantos/gomusic/cmd/gomusic/internal/auth/middleware"
	router2 "github.com/lcnssantos/gomusic/cmd/gomusic/internal/auth/router"
	"github.com/lcnssantos/gomusic/cmd/gomusic/internal/auth/service"

	"github.com/gorilla/mux"
	"github.com/lcnssantos/gomusic/cmd/gomusic/internal/user/controller"
	"github.com/lcnssantos/gomusic/cmd/gomusic/internal/user/repository"
	"github.com/lcnssantos/gomusic/cmd/gomusic/internal/user/router"
	service2 "github.com/lcnssantos/gomusic/cmd/gomusic/internal/user/service"
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
