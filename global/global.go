package global

import (
	"database/sql"

	"github.com/Youknow2509/go-ecommerce/pkg/logger"
	"github.com/Youknow2509/go-ecommerce/pkg/setting"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
	Rdb    *redis.Client
	Mdbc   *sql.DB
)
