package producer

import (
	"context"
	broker "landate/internals/broker"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func MessageProducer() {

	ch, bctx := broker.RabbitMQClient()

	err := ch.ExchangeDeclare(
		"payload", // name
		"fanout",  // type
		true,      // durable
		false,     // auto-deleted
		false,     // internal
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	ctx, cancel := context.WithTimeout(bctx, 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"logs", // exchange
		"",     // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(""),
		})

	failOnError(err, "Failed to publish a message")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
