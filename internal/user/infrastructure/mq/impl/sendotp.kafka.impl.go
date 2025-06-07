package impl

import (
	"context"

	"github.com/Youknow2509/go-ecommerce/internal/user/infrastructure/mq"
	"github.com/segmentio/kafka-go"
)

type SendOtpService struct {
	writer *kafka.Writer
}

// ########################################################

// Publish implements mq.IKafkaService.
func (s *SendOtpService) Publish(ctx context.Context, topic string, key string, partition int32, payload []byte) error {
	msg := kafka.Message{
		Key:       []byte(key),
		Partition: int(partition),
		Value:     payload,
	}

	return s.writer.WriteMessages(ctx, msg)
}

// ########################################################

// new and impl
func NewSendOtpService(writer *kafka.Writer) mq.IKafkaService {
	return &SendOtpService{
		writer: writer,
	}
}
