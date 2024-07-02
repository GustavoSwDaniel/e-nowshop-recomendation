package config

import (
	"os"
)

type Config struct {
	RabbitMqUrl string
	DatabaseUrl string
}

func LoadConfig() *Config {
	return &Config{
		RabbitMqUrl: os.Getenv("RABBITMQ_URL"),
		DatabaseUrl: os.Getenv("DATABASE_URL"),
	}
}
