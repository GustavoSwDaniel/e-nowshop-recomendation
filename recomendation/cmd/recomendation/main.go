package main

import (
	"fmt"
	"log"

	"github.com/GustavoSwDaniel/e-nowshop-recomendation/recomendation/internal/infrastructure/database"
	"github.com/GustavoSwDaniel/e-nowshop-recomendation/recomendation/internal/infrastructure/rabbitmq"
	"github.com/GustavoSwDaniel/e-nowshop-recomendation/recomendation/internal/products"
	"github.com/GustavoSwDaniel/e-nowshop-recomendation/recomendation/pkg/config"
)

func main() {
	configs := config.LoadConfig()
	conn, err := rabbitmq.Connect(configs.RabbitMqUrl)
	if err != nil {
		log.Fatalf("Deu erro aqui: %v", err)
	}
	fmt.Println("Conectado")
	defer conn.Close()
	databaseConn, err := database.ConnectionDatabase(configs.DatabaseUrl)
	if err != nil {
		log.Fatalf("Erro ao conectar no banco")
	}
	fmt.Println("Conectado ao banco de dados")

	productService := products.ServiceProducs{
		RepositoryProducts: &products.RepositoryProducts{
			Conn: databaseConn,
		},
	}
	productService.GetOrdersMetrics(54)
	consumer := rabbitmq.NewConsumer(conn)
	if err := consumer.Consumer("Testando", productService.GetOrdersMetrics); err != nil {
		log.Fatalf("Deus pau, %v", err)
	}

}
