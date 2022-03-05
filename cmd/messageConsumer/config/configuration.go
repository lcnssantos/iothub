package config

import (
	"os"

	"github.com/go-playground/validator/v10"
)

type Configuration struct {
	MONGO_HOST   string `validate:"required"`
	MONGO_USER   string `validate:"required"`
	MONGO_PASS   string `validate:"required"`
	MONGO_PORT   string `validate:"required"`
	MONGO_DBNAME string `validate:"required"`
	RMQ_HOST     string `validate:"required"`
	RMQ_PORT     string `validate:"required"`
	RMQ_USER     string `validate:"required"`
	RMQ_PASS     string `validate:"required"`
}

func Get() *Configuration {
	return &Configuration{
		MONGO_DBNAME: os.Getenv("MONGO_DBNAME"),
		MONGO_PASS:   os.Getenv("MONGO_PASS"),
		MONGO_HOST:   os.Getenv("MONGO_HOST"),
		MONGO_PORT:   os.Getenv("MONGO_PORT"),
		MONGO_USER:   os.Getenv("MONGO_USER"),
		RMQ_HOST:     os.Getenv("CONSUMER_RMQ_HOST"),
		RMQ_PASS:     os.Getenv("CONSUMER_RMQ_PASS"),
		RMQ_PORT:     os.Getenv("CONSUMER_RMQ_PORT"),
		RMQ_USER:     os.Getenv("CONSUMER_RMQ_USER"),
	}
}

func Validate() error {
	configuration := Get()
	err := validator.New().Struct(configuration)
	return err
}
