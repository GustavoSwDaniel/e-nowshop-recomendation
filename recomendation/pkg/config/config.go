package config

type Config struct {
	RabbitMqUrl string
	DatabaseUrl string
}

// //postgres:t8ltavvRLd@127.0.0.1:5432/postgres
func LoadConfig() *Config {
	return &Config{
		RabbitMqUrl: "amqp://admin:admin@localhost:5672/",
		DatabaseUrl: "postgres://postgres:t8ltavvRLd@127.0.0.1:5432/postgres?sslmode=disable",
	}
}
