package main

import (
	"log"

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
}
