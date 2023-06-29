package broker

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

type MsgBroker struct {
	mbChannel *amqp091.Channel
}

func (mb *MsgBroker) RabbitMQClient() *MsgBroker {

	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("Failed Initializing Broker connection.")
	}

	channel, err := conn.Channel()
	if err != nil {
		fmt.Println("Error:", err)
	}

	defer channel.Close()
	// mb.mbChannel = channel
	return &MsgBroker{
		mbChannel: channel,
	}
}

func (mb *MsgBroker) RabbitProducer() {

	queue, err := mb.mbChannel.QueueDeclare(
		"TestMsgQueue",
		false,
		false,
		false,
		false,
		nil, // arguments
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(queue)

}
