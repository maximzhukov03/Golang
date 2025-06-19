package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct{
	DB_PASSWORD string
	DB_USER string
	DB_NAME string
	DB_PORT string
	DB_HOST string
}

func NewConfig() Config{
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка env")
	}

	return Config{
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_USER: os.Getenv("DB_USER"),
		DB_NAME: os.Getenv("DB_NAME"),
		DB_PORT: os.Getenv("DB_PORT"),
		DB_HOST: os.Getenv("DB_HOST"),

	}
}