package main

import (
	"context"
	"log"

	"github.com/lcnssantos/iothub/cmd/messageConsumer/config"
	"github.com/lcnssantos/iothub/internal/database"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	err := config.Validate()

	if err != nil {
		log.Fatalln("invalid environment credentials")
	}

	configuration := config.Get()

	mongoDriver := database.NewMongoDriver(configuration.MONGO_HOST, configuration.MONGO_PORT, configuration.MONGO_USER, configuration.MONGO_PASS, configuration.MONGO_DBNAME)

	mongoDb, err := mongoDriver.GetDatabase()

	if err != nil {
		return
	}

	one, err := mongoDb.Collection("messages").InsertOne(context.Background(), bson.M{"name": "Luciano Souza Santos", "email": "luciano.ssants@gmail.com", "password": "llwigp62"})
	if err != nil {
		return
	}

	log.Println(one)
}
