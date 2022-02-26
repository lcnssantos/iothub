package main

import (
	"database/sql"
	repository2 "github.com/lcnssantos/iothub/cmd/iothub/internal/accounts/repository"
	service3 "github.com/lcnssantos/iothub/cmd/iothub/internal/accounts/service"

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
	accountRepository := repository2.NewAccountRepository(db)

	accountService := service3.NewAccountService(accountRepository)
	hashService := service2.NewHashService()
	userService := service2.NewUserService(userRepository, hashService, accountService)
	jwtService := service.NewJwtService()
	authService := service.NewAuthService(userService, hashService, jwtService)

	userController := controller.NewUserController(userService)
	meController := controller.NewMeController()

	authMiddleware := middleware.NewAuthenticationMiddleware(authService)

	router.BuildUserRouter(userController, r.PathPrefix("/user").Subrouter())
	router.BuildMeRouter(meController, r.PathPrefix("/me").Subrouter(), authMiddleware)
}

func BuildAuthModule(db *sql.DB, r *mux.Router) {
	userRepository := repository.NewUserRepository(db)
	accountRepository := repository2.NewAccountRepository(db)

	accountService := service3.NewAccountService(accountRepository)
	hashService := service2.NewHashService()
	userService := service2.NewUserService(userRepository, hashService, accountService)
	jwtService := service.NewJwtService()
	authService := service.NewAuthService(userService, hashService, jwtService)

	authController := controller2.NewAuthController(authService)
	router2.BuildRouter(authController, r)
}
