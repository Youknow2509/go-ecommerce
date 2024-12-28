package global

import (
	"github.com/Youknow2509/go-ecommerce/pkg/logger"
	"github.com/Youknow2509/go-ecommerce/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)
