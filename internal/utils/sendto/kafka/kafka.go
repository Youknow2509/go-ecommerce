package sendto

import (
	"encoding/json"
	"time"

	"github.com/Youknow2509/go-ecommerce/internal/utils/sendto"
	"github.com/Youknow2509/go-ecommerce/pkg/utils/kafka/impl"
	"github.com/segmentio/kafka-go"
)

type kafka_send_mail struct {
}

// SendAPIEmailOTP implements sendto.ISendTo.
func (k *kafka_send_mail) SendAPIEmailOTP(to string, from string, otp string) error {
	panic("unimplemented")
}

// SendTemplateEmailOTP implements sendto.ISendTo.
func (k *kafka_send_mail) SendTemplateEmailOTP(to []string, from string, nameTemplate string, dataTemplate map[string]interface{}) error {
	panic("unimplemented")
}

// SendTextEmailOTP implements sendto.ISendTo.
func (k *kafka_send_mail) SendTextEmailOTP(to []string, from string, otp string) error {
	panic("unimplemented")
}

// SendKafkaEmailOTP implements ISendTo.
func (k *kafka_send_mail) SendKafkaEmailOTP(to string, from string, otp string) error {
	body := make(map[string]interface{})

	body["from"] = from
	body["to"] = to
	body["type"] = 1
	body["data"] = otp
	// requestBody
	requestBody, _ := json.Marshal(body)

	// create message in kafaka
	msg := kafka.Message{
		Key:   []byte("register"),
		Value: []byte(requestBody),
		Time:  time.Now(),
	}

	return impl.NewProducerSendMailOtpImpl().WriteMessages(msg)
}

// implement ISendTo interface and create
func NewKafkaSendTo() sendto.ISendTo {
	return &kafka_send_mail{}
}
