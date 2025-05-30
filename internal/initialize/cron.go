package initialize

import (
	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/robfig/cron/v3"
)

func InitCronJob() {
	cr := cron.New()

	global.Cron = cr
}
