package initialize

import (
	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/Youknow2509/go-ecommerce/pkg/logger"
)

// Initial Logger
func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
