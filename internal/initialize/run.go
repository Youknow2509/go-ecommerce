package initialize

import (
	"fmt"
	"strconv"

	"github.com/Bot-SomeOne/go-ecommerce/global"
)

// Run all initialization
func Run() {
	// load configuration
	LoadConfig()
	fmt.Println("@@@@ Loaded configuration")

	// initialize logger
	InitLogger()
	global.Logger.Info("Logger initialized")

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