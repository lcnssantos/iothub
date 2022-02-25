package user

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func Build(db *sql.DB, router *mux.Router) {
	repository := NewRepository(db)
	hashService := NewHashService()
	service := NewService(repository, hashService)
	controller := NewController(service)
	BuildRouter(controller, router)
}
