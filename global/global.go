package global

import (
	"database/sql"

	"github.com/Youknow2509/go-ecommerce/pkg/logger"
	"github.com/Youknow2509/go-ecommerce/pkg/setting"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
	amqp "github.com/rabbitmq/amqp091-go"

)

var (
	Config        setting.Config
	Logger        *logger.LoggerZap
	Mdb           *gorm.DB
	Rdb           *redis.Client
	Mdbc          *sql.DB
	KafkaProducer *kafka.Writer
	RabbitMQProducer_LOGGERDISCORD *amqp.Channel
	Prometheus *setting.PrometheusSetting
)
