package initialize

import (
	"github.com/Bot-SomeOne/go-ecommerce/global"
	"github.com/Bot-SomeOne/go-ecommerce/pkg/logger"
)

// Initial Logger
func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}