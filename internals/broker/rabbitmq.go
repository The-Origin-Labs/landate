package broker

import (
	"context"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type MsgBroker struct {
	mbChannel *amqp091.Channel
	ctx       context.Context
}

func RabbitMQClient() (*amqp091.Channel, context.Context) {

	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	channel, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer channel.Close()

	ctx := context.Background()
	return channel, ctx
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
