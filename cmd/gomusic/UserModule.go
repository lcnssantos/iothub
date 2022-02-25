package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/lcnssantos/gomusic/cmd/gomusic/internal/user/controller"
	"github.com/lcnssantos/gomusic/cmd/gomusic/internal/user/repository"
	"github.com/lcnssantos/gomusic/cmd/gomusic/internal/user/router"
	service2 "github.com/lcnssantos/gomusic/cmd/gomusic/internal/user/service"
)

func Build(db *sql.DB, r *mux.Router) {
	userRepository := repository.NewUserRepository(db)
	hashService := service2.NewHashService()
	userService := service2.NewUserService(userRepository, hashService)
	userController := controller.NewUserController(userService)
	router.BuildRouter(userController, r)
}
