package initialize

import (
	"fmt"

	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/Youknow2509/go-ecommerce/pkg/logger"
	"github.com/gin-gonic/gin"
)

// Run all initialization
func Run() *gin.Engine {
	// load configuration
	LoadConfig()
	fmt.Println("@@@ Loader configuration")	

	// connect to rabbit mq
	InitRabbitMQ()
	fmt.Println("RabbitMQ initialized")
	logger.ChannelRabbitMq = global.RabbitMQProducer_LOGGERDISCORD

	// initialize logger
	InitLogger()
	global.Logger.Info("Logger initialized")

	// initialize prometheus
	InitPrometheus()
	global.Logger.Info("Prometheus initialized")

	// connect to my sql
	InitMysql()
	global.Logger.Info("Mysql initialized")

	// innitialize sqlc
	InitMysqlC()
	global.Logger.Info("MysqlC initialized")

	// initialize service interface
	InitServiceInterface()
	global.Logger.Info("Service interface initialized")

	// connect to redis
	InitRedis()
	global.Logger.Info("Redis initialized")

	// connect to kafka
	InitKafka()
	global.Logger.Info("Kafka initialized")

	// connect to Router
	r := InitRouter()

	// run server
	// port := strconv.Itoa(global.Config.Server.Port)
	// r.Run(":" + port)

	return r
}
