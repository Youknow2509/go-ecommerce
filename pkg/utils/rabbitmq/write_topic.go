package rabbitmq

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Write interface {
	WriteToTopic(topic string, routingKey string, data []byte) error
	// ....
}

type rabbitmqWriter struct {
	// RabbitMQ connection
	channel *amqp.Channel
}

func NewRabbitMQWriter(ch *amqp.Channel) Write {
	return &rabbitmqWriter{
		channel: ch,
	}
}

// WriteToTopic implements Write.
func (rb *rabbitmqWriter) WriteToTopic(topic string, routingKey string, data []byte) error {
	err := rb.channel.Publish(
		topic,
		"logger.info",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/json",
			Body:        data,
		},
	)
	if err != nil {
		fmt.Println("Failed to publish a message", err)
		return err
	}

	fmt.Printf("Message published to topic:: %s with routingkey:: logger.info\n", topic)

	return nil
}
