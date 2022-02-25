package config

import (
	"log"
	"os"

	"github.com/go-playground/validator/v10"
)

type Configuration struct {
	DB_DRIVER    string `validate:"required"`
	DB_HOST      string `validate:"required"`
	DB_USER      string `validate:"required"`
	DB_PASS      string `validate:"required"`
	DB_PORT      string `validate:"required"`
	DB_NAME      string `validate:"required"`
	DB_POOL_SIZE string `validate:"required"`
}

func Get() *Configuration {
	return &Configuration{
		DB_DRIVER:    os.Getenv("DB_DRIVER"),
		DB_HOST:      os.Getenv("DB_HOST"),
		DB_USER:      os.Getenv("DB_USER"),
		DB_PASS:      os.Getenv("DB_PASS"),
		DB_PORT:      os.Getenv("DB_PORT"),
		DB_NAME:      os.Getenv("DB_PORT"),
		DB_POOL_SIZE: os.Getenv("DB_POOL_SIZE"),
	}
}

func Validate() {
	environmentConfiguration := Get()

	if err := validator.New().Struct(environmentConfiguration); err != nil {
		log.Panicln("Invalid Environment Variables")
	}
}
