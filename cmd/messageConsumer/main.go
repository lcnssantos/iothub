package main

import (
	"context"
	"log"

	"github.com/lcnssantos/iothub/cmd/messageConsumer/mqtt"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/lcnssantos/iothub/cmd/messageConsumer/config"
	"github.com/lcnssantos/iothub/internal/database"
)

func getMongoDb() (*mongo.Database, error) {
	configuration := config.Get()

	mongoDriver := database.NewMongoDriver(configuration.MONGO_HOST, configuration.MONGO_PORT, configuration.MONGO_USER, configuration.MONGO_PASS, configuration.MONGO_DBNAME)
	mongoDb, err := mongoDriver.GetDatabase()

	if err != nil {
		return nil, err
	}

	return mongoDb, nil
}

func getRMQ() (*mqtt.MQTT, error) {
	configuration := config.Get()

	mqttClient := mqtt.NewMQTT(configuration.RMQ_HOST, configuration.RMQ_PORT, configuration.RMQ_USER, configuration.RMQ_PASS)
	err := mqttClient.Connect()

	return mqttClient, err
}

func handleMongo(done chan interface{}, db *mongo.Database) {
	for {
		message := <-done
		log.Println(message)
		db.Collection("messages").InsertOne(context.Background(), &message)
	}
}

func main() {
	err := config.Validate()

	if err != nil {
		log.Fatalln("invalid environment credentials")
	}

	db, err := getMongoDb()

	if err != nil {
		log.Println(err)
		log.Println("Error to connect to mongo")
	}

	defer db.Client().Disconnect(context.Background())

	mqttClient, err := getRMQ()

	if err != nil {
		log.Println("Error to connect to mqtt")
		return
	}

	defer mqttClient.Close()

	mqttChan := make(chan interface{})
	mongoChan := make(chan interface{})

	go mqttClient.Listen("messages", mqttChan)
	go handleMongo(mongoChan, db)

	for {
		newMessage := <-mqttChan
		mongoChan <- newMessage
	}

}
