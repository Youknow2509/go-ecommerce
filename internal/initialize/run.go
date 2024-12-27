package initialize

import (
	"strconv"

	"github.com/Bot-SomeOne/go-ecommerce/global"
)

// Run all initialization
func Run() {
	// load configuration
	LoadConfig()

	// initialize logger
	InitLogger()

	// connect to my sql
	InitMysql()

	// connect to redis
	InitRedis()

	// connect to kafka
	InitKafka()

	// TODO ...

	// connect to Router
	r := InitRouter()

	// run server
	port := strconv.Itoa(global.Config.Server.Port)
	r.Run(":" + port)
}