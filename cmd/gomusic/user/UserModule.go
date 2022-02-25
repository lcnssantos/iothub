package user

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/lcnssantos/gomusic/cmd/gomusic/user/internal/controller"
	"github.com/lcnssantos/gomusic/cmd/gomusic/user/internal/repository"
	"github.com/lcnssantos/gomusic/cmd/gomusic/user/internal/router"
	service2 "github.com/lcnssantos/gomusic/cmd/gomusic/user/internal/service"
)

func Build(db *sql.DB, r *mux.Router) {
	userRepository := repository.NewUserRepository(db)
	hashService := service2.NewHashService()
	userService := service2.NewUserService(userRepository, hashService)
	userController := controller.NewUserController(userService)
	router.BuildRouter(userController, r)
}
