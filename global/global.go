package global

import (
	"github.com/Bot-SomeOne/go-ecommerce/pkg/logger"
	"github.com/Bot-SomeOne/go-ecommerce/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb   *gorm.DB
)