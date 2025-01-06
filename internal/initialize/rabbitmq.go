package initialize

import (
	"fmt"

	"github.com/Youknow2509/go-ecommerce/global"
	amqp "github.com/rabbitmq/amqp091-go"
)

func InitRabbitMQ() {
	InitRabbitMQExchangeChannel()
}

// init rabbitmq exchange channel handle logging to discord
func InitRabbitMQExchangeChannel() {
	conn, err := amqp.Dial(global.Config.RabbitMQ.Url_RB_D)
	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ", err)
		return
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Failed to open a channel", err)
		return
	}

	name_topic := "logger_discord"
	// Declare a topic exchange
	err = ch.ExchangeDeclare(
		name_topic, // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		fmt.Println("Failed to declare an exchange", err)
	}

	// set global
	global.RabbitMQProducer_LOGGERDISCORD = ch
	fmt.Println("RabbitMQ Handler logger to Discord initialized")
}
