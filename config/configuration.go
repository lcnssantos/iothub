package config

import (
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
	PORT         string `validate:"required"`
	JWT_KEY      string `validate:"required"`
	RMQ_HOST     string `validate:"required"`
	RMQ_PORT     string `validate:"required"`
	RMQ_USER     string `validate:"required"`
	RMQ_PASS     string `validate:"required"`
	RMQ_API_URL  string `validate:"required""`
}

func Get() *Configuration {
	return &Configuration{
		DB_DRIVER:    os.Getenv("DB_DRIVER"),
		DB_HOST:      os.Getenv("DB_HOST"),
		DB_USER:      os.Getenv("DB_USER"),
		DB_PASS:      os.Getenv("DB_PASS"),
		DB_PORT:      os.Getenv("DB_PORT"),
		DB_NAME:      os.Getenv("DB_NAME"),
		DB_POOL_SIZE: os.Getenv("DB_POOL_SIZE"),
		PORT:         os.Getenv("PORT"),
		JWT_KEY:      os.Getenv("JWT_KEY"),
		RMQ_HOST:     os.Getenv("RMQ_HOST"),
		RMQ_PASS:     os.Getenv("RMQ_PASS"),
		RMQ_PORT:     os.Getenv("RMQ_PORT"),
		RMQ_USER:     os.Getenv("RMQ_USER"),
		RMQ_API_URL:  os.Getenv("RMQ_API_URL"),
	}
}

func Validate() error {
	environmentConfiguration := Get()
	err := validator.New().Struct(environmentConfiguration)
	return err
}
