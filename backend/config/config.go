package config

import (
	"fmt"
	"log"
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

const (
	LimitMaxAccountNameChars     = 100
	LimitMinAccountNameChars     = 1
	LimitMaxTransactionNoteChars = 200
	LimitMaxCategoryNameChars    = 150
	LimitMinCategoryNameChars    = 1
)

type AppConfig struct {
	PostgresUrl   string `env:"POSTGRES_URL"`
	RabbitMqUrl   string `env:"RABBIT_MQ_URL"`
	ServerAddress string `env:"SERVER_ADDRESS"`
	ServerPort    int    `env:"SERVER_PORT"`
}

func Get() *AppConfig {
	fmt.Println(os.Getwd())
	environment := os.Getenv("APP_ENV")

	if environment == "dev" {
		if err := godotenv.Load("../.env"); err != nil {
			log.Fatal(err)
		}
	}
	cfg := &AppConfig{}

	if err := env.Parse(cfg); err != nil {
		panic(err)
	}

	return cfg
}
