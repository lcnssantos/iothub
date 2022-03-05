package modules

import (
	"database/sql"

	"github.com/lcnssantos/iothub/cmd/publicApi/config"

	repository2 "github.com/lcnssantos/iothub/cmd/publicApi/internal/accounts/repository"
	service3 "github.com/lcnssantos/iothub/cmd/publicApi/internal/accounts/service"
	controller2 "github.com/lcnssantos/iothub/cmd/publicApi/internal/auth/controller"
	"github.com/lcnssantos/iothub/cmd/publicApi/internal/auth/middleware"
	router2 "github.com/lcnssantos/iothub/cmd/publicApi/internal/auth/router"
	"github.com/lcnssantos/iothub/cmd/publicApi/internal/auth/service"
	"github.com/lcnssantos/iothub/cmd/publicApi/internal/rmq"

	"github.com/gorilla/mux"
	"github.com/lcnssantos/iothub/cmd/publicApi/internal/user/controller"
	"github.com/lcnssantos/iothub/cmd/publicApi/internal/user/repository"
	"github.com/lcnssantos/iothub/cmd/publicApi/internal/user/router"
	service2 "github.com/lcnssantos/iothub/cmd/publicApi/internal/user/service"
)

func BuildUserModule(db *sql.DB, r *mux.Router) {
	config := config.Get()

	userRepository := repository.NewUserRepository(db)
	accountRepository := repository2.NewAccountRepository(db)

	rmqClient := rmq.NewRMQClient(config.RMQ_API_URL, config.RMQ_USER, config.RMQ_PASS)

	accountService := service3.NewAccountService(accountRepository, rmqClient)
	hashService := service2.NewHashService()
	userService := service2.NewUserService(userRepository, hashService, accountService)
	jwtService := service.NewJwtService()
	authService := service.NewAuthService(userService, hashService, jwtService)

	userController := controller.NewUserController(userService)
	meController := controller.NewMeController()

	authMiddleware := middleware.NewAuthenticationMiddleware(authService)

	router.BuildUserRouter(userController, r.PathPrefix("/user").Subrouter(), authMiddleware)
	router.BuildMeRouter(meController, r.PathPrefix("/me").Subrouter(), authMiddleware)
}

func BuildAuthModule(db *sql.DB, r *mux.Router) {
	config := config.Get()

	userRepository := repository.NewUserRepository(db)
	accountRepository := repository2.NewAccountRepository(db)
	rmqClient := rmq.NewRMQClient(config.RMQ_API_URL, config.RMQ_USER, config.RMQ_PASS)

	accountService := service3.NewAccountService(accountRepository, rmqClient)
	hashService := service2.NewHashService()
	userService := service2.NewUserService(userRepository, hashService, accountService)
	jwtService := service.NewJwtService()
	authService := service.NewAuthService(userService, hashService, jwtService)

	authController := controller2.NewAuthController(authService)
	router2.BuildRouter(authController, r)
}
