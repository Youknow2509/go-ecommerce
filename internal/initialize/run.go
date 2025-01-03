package initialize

import (
	"strconv"
	"fmt"
	"github.com/Youknow2509/go-ecommerce/global"
)

// Run all initialization
func Run() {
	// load configuration
	LoadConfig()
	fmt.Println("@@@ Loader configuration")
	
	// initialize logger
	InitLogger()
	global.Logger.Info("Logger initialized")

	// connect to my sql
	InitMysql()
	global.Logger.Info("Mysql initialized")

	// innitialize sqlc
	InitMysqlC()
	global.Logger.Info("MysqlC initialized")

	// connect to redis
	InitRedis()
	global.Logger.Info("Redis initialized")

	// connect to kafka
	InitKafka()

	// TODO ...

	// connect to Router
	r := InitRouter()

	// run server
	port := strconv.Itoa(global.Config.Server.Port)
	r.Run(":" + port)
}
