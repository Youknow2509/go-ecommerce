package kafka

import (
	k "github.com/segmentio/kafka-go"
)

// Interface for kafka producer
type IKafkaProducer interface {
	WriteMessages(messages ...k.Message) error
	Close() error
}
