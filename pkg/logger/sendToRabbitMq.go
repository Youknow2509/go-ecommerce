package logger

import (
	"github.com/Youknow2509/go-ecommerce/pkg/utils/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	ChannelRabbitMq *amqp.Channel
)

type RabbitMQLogger struct {
	producer rabbitmq.Write
}

// handle io.Writer in RabbitMQLogger
func (r *RabbitMQLogger) Write(p []byte) (n int, err error) {
	err = r.producer.WriteToTopic("logger_discord", "logger.info", p)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}
