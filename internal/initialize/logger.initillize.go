package initialize

import (
	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/Youknow2509/go-ecommerce/pkg/logger"
)

func InitializeLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}