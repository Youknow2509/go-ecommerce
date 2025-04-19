package impl

import (
	"context"

	"github.com/Youknow2509/go-ecommerce/global"
	iKafka "github.com/Youknow2509/go-ecommerce/pkg/utils/kafka"
	kafka "github.com/segmentio/kafka-go"
)

var producer_send_mail_otp *kafka.Writer

type producer_send_mail_otp_impl struct {
}

// Close implements kafka.IKafkaProducer.
func (p *producer_send_mail_otp_impl) Close() error {
	panic("unimplemented")
}

// WriteMessages implements kafka.IKafkaProducer.
func (p *producer_send_mail_otp_impl) WriteMessages(messages ...kafka.Message) error {
	return producer_send_mail_otp.WriteMessages(context.Background(), messages...)
}

// new producer_send_mail_otp_impl
func NewProducerSendMailOtpImpl() iKafka.IKafkaProducer {
	producer_send_mail_otp = &kafka.Writer{
		Addr:     kafka.TCP(global.Config.Kafka.TcpHost),
		Topic:    global.Config.Kafka.TopicServiceSendMail,
		Balancer: &kafka.LeastBytes{},
	}

	return &producer_send_mail_otp_impl{}
}