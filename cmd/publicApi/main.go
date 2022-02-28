package main

import (
	"database/sql"
	"fmt"
	"github.com/lcnssantos/iothub/cmd/publicApi/config"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lcnssantos/iothub/internal/middlewares"

	"github.com/lcnssantos/iothub/internal/database"
)

func BuildModules(db *sql.DB, router *mux.Router) {
	BuildUserModule(db, router.PathPrefix("/").Subrouter())
	BuildAuthModule(db, router.PathPrefix("/auth").Subrouter())
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
	router.Use(middlewares.NewLogMiddleware().Handler)

	BuildModules(db, router)

	http.Handle("/", router)
	http.ListenAndServe(fmt.Sprintf(":%s", environmentConfiguration.PORT), nil)
}
