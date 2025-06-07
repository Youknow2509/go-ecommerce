package global

import (
	"database/sql"

	"github.com/Youknow2509/go-ecommerce/pkg/logger"
	"github.com/Youknow2509/go-ecommerce/pkg/setting"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron/v3"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	Cron                           *cron.Cron
	Config                         setting.Config
	Logger                         *logger.LoggerZap
	Mdb                            *gorm.DB
	Rdb                            *redis.Client
	MongoClient                    *mongo.Client
	Mdbc                           *sql.DB
	RabbitMQProducer_LOGGERDISCORD *amqp.Channel
	Prometheus                     *setting.PrometheusSetting
	KafkaSendOtp                   *kafka.Writer
	
)
