package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lcnssantos/gomusic/cmd/gomusic/user"
	"github.com/lcnssantos/gomusic/internal/middlewares"

	"github.com/lcnssantos/gomusic/config"
	"github.com/lcnssantos/gomusic/internal/database"
)

func main() {
	config.Validate()
	db, err := database.GetConnection()

	if err != nil {
		log.Panicln("Error to connect to database")
	}

	defer db.Close()

	router := mux.NewRouter()
	router.Use(middlewares.NewJsonMiddleware().Handler)

	user.Build(db, router.PathPrefix("/v1/user").Subrouter())

	http.Handle("/", router)
	http.ListenAndServe(":5000", nil)
}
