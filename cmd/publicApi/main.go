package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/lcnssantos/iothub/cmd/publicApi/config"

	"github.com/gorilla/mux"
	"github.com/lcnssantos/iothub/internal/middlewares"

	"github.com/lcnssantos/iothub/cmd/publicApi/modules"
	"github.com/lcnssantos/iothub/internal/database"
)

func BuildModules(db *sql.DB, router *mux.Router) {
	modules.BuildUserModule(db, router.PathPrefix("/").Subrouter())
	modules.BuildAuthModule(db, router.PathPrefix("/auth").Subrouter())
}

func BuildRouters() *mux.Router {
	router := mux.NewRouter().PathPrefix("/v1").Subrouter()
	router.Use(middlewares.NewJsonMiddleware().Handler)
	return router
}

func main() {
	err := config.Validate()

	if err != nil {
		log.Fatalln("Error to validate environment variables")
	}

	environmentConfiguration := config.Get()
	db, err := database.GetConnection()

	if err != nil {
		log.Fatalln("Error to connect to database")
	}

	defer db.Close()

	router := BuildRouters()
	router.Use(middlewares.NewLogMiddleware().Handler)

	BuildModules(db, router)

	http.Handle("/", router)
	http.ListenAndServe(fmt.Sprintf(":%s", environmentConfiguration.PORT), nil)
}
