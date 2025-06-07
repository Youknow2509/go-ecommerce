package initialize

import (
	"fmt"

	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/segmentio/kafka-go"
)

func InitKafka() {
	// Initialize the Kafka writer for sending OTPs
	InitializeKafkaSendOtp()
}

func InitializeKafkaSendOtp() {
	global.KafkaSendOtp = kafka.NewWriter(
		kafka.WriterConfig{
			Brokers:  []string{fmt.Sprintf("%s:%d", global.Config.KafkaSendOtp.Host, global.Config.KafkaSendOtp.Port)},
			Topic:    global.Config.KafkaSendOtp.Topic,
			Balancer: &kafka.LeastBytes{},
		},
	)	
}
