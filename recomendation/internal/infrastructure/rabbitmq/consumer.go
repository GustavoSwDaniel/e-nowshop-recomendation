package rabbitmq

import (
	"encoding/json"
	"log"

	"github.com/GustavoSwDaniel/e-nowshop-recomendation/recomendation/internal/dto/rabbitmq"
	"github.com/streadway/amqp"
)

type Consumer struct {
	conn *amqp.Connection
}

func NewConsumer(conn *amqp.Connection) *Consumer {
	return &Consumer{conn: conn}
}

func (c *Consumer) Consumer(queueName string, function func(int64)) error {
	ch, err := c.conn.Channel()
	if err != nil {
		return err
	}

	defer ch.Close()
	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			dto := &rabbitmq.ProductMessageRabbitmq{}
			err := json.Unmarshal(d.Body, dto)
			if err != nil {
				log.Fatal("Error na deserialização")
			}
			log.Printf("Mensagem %+v", dto)
			function(dto.ProductId)
		}
	}()

	log.Printf("Esperando mesagem")
	<-forever
	return nil
}
