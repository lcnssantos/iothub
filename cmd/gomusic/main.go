package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lcnssantos/gomusic/cmd/gomusic/user"
	"github.com/lcnssantos/gomusic/internal/middlewares"

	"github.com/lcnssantos/gomusic/config"
	"github.com/lcnssantos/gomusic/internal/database"
)

func BuildModules(db *sql.DB, router *mux.Router) {
	user.Build(db, router.PathPrefix("/user").Subrouter())
}

func BuildRouters() *mux.Router {
	router := mux.NewRouter().PathPrefix("/v1").Subrouter()
	router.Use(middlewares.NewJsonMiddleware().Handler)
	return router
}

func main() {
	err := config.Validate()

	if err != nil {
		log.Panicln("Error to validate environment variables")
	}

	environmentConfiguration := config.Get()
	db, err := database.GetConnection()

	if err != nil {
		log.Panicln("Error to connect to database")
	}

	defer db.Close()

	router := BuildRouters()

	BuildModules(db, router)

	http.Handle("/", router)
	http.ListenAndServe(fmt.Sprintf(":%s", environmentConfiguration.PORT), nil)
}
