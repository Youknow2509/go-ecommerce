package cron

import (
	"github.com/Youknow2509/go-ecommerce/internal/controller/cron"
	"github.com/gin-gonic/gin"
)

type AdminRouter struct {

}

// init admin cron router
func (ar *AdminRouter) InitAdminCronRouter(Router *gin.RouterGroup) {
	// public router
	adminRouterPublic := Router.Group("/cron")
	{
		adminRouterPublic.POST("/inspace", cron.CronJob.InspaceCronJob)
		adminRouterPublic.POST("/start", cron.CronJob.StartCronJob)
		adminRouterPublic.POST("/stop", cron.CronJob.StopCronJob)
	}
}