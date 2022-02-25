package user

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/lcnssantos/gomusic/cmd/gomusic/user/controller"
	"github.com/lcnssantos/gomusic/cmd/gomusic/user/repository"
	"github.com/lcnssantos/gomusic/cmd/gomusic/user/router"
	"github.com/lcnssantos/gomusic/cmd/gomusic/user/service"
)

func Build(db *sql.DB, r *mux.Router) {
	userRepository := repository.NewRepository(db)
	hashService := service.NewHashService()
	userService := service.NewService(userRepository, hashService)
	userController := controller.NewController(userService)
	router.BuildRouter(userController, r)
}
