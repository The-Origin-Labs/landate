package events

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	BROKER_URL = "amqp://guest:guest@localhost:5672"
)

func RabbitClient() (*amqp.Connection, *amqp.Channel) {
	mqURL := os.Getenv("BROKER_URL")
	if mqURL == "" {
		BROKER_URL = mqURL
	}

	conn, err := amqp.Dial(BROKER_URL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	return conn, ch
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
