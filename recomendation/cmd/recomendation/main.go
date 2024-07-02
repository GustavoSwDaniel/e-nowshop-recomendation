package main

import (
	"fmt"
	"log"

	"github.com/GustavoSwDaniel/e-nowshop-recomendation/recomendation/internal/infrastructure/database"
	"github.com/GustavoSwDaniel/e-nowshop-recomendation/recomendation/internal/infrastructure/rabbitmq"
	orderitens "github.com/GustavoSwDaniel/e-nowshop-recomendation/recomendation/internal/orderItens"
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
	orderItensService := orderitens.ServiceOrdersItens{
		RepositoryOrdersItems: &orderitens.RepositoryOrdersItems{
			Conn: databaseConn,
		},
		PorductService: &productService,
	}
	consumer := rabbitmq.NewConsumer(conn)
	if err := consumer.Consumer("Testando", orderItensService.GetOrdersMetrics); err != nil {
		log.Fatalf("Deus pau, %v", err)
	}

}
