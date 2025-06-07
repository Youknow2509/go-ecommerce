package mq

import (
	"context"
	"fmt"
)

// interface IKafkaService
type IKafkaService interface {
	Publish(ctx context.Context, topic string, key string, partition int32, payload []byte) error
}

// ################################################

var (
	vIKafkaService IKafkaService
)

// ################################################

// init kafka service
func InitKafkaService(kafkaService IKafkaService) {
	vIKafkaService = kafkaService
}

// GetKafkaService
func GetKafkaService() (IKafkaService, error) {
	if vIKafkaService == nil {
		return nil, fmt.Errorf("kafka service is not initialized")
	}
	return vIKafkaService, nil
}
