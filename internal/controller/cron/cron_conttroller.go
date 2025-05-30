package cron

import (
	"fmt"

	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/Youknow2509/go-ecommerce/internal/model"
	"github.com/Youknow2509/go-ecommerce/response"
	"github.com/gin-gonic/gin"
)

// manager controller Cron
var CronJob = new(cCronJob)

type cCronJob struct{}

// @Summary	  		Inspace cron job
// @Description  	Inspace cron job
// @Tags         	cron management
// @Accept       	json
// @Produce      	json
// @Success      	200  {object}  response.ResponseData
// @Failure      	500  {object}  response.ErrResponseData
// @Router       	/v1/cron/inspace [post]
func (p *cCronJob) InspaceCronJob(ctx *gin.Context) {

	data := make([]model.CronEntry, 0)
	for _, entry := range global.Cron.Entries() {
		data = append(data, model.CronEntry{
			ID:       string(entry.ID),
			Next:     entry.Next.String(),
			Prev:     entry.Prev.String(),
			Schedule: fmt.Sprintf("%T", entry.Schedule),
		})
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, data)
}

// @Summary	  		Start cron job
// @Description  	Start cron job
// @Tags         	cron management
// @Accept       	json
// @Produce      	json
// @Success      	200  {object}  response.ResponseData
// @Failure      	500  {object}  response.ErrResponseData
// @Router       	/v1/cron/start [post]
func (p *cCronJob) StartCronJob(ctx *gin.Context) {
	global.Cron.Start()
	global.Logger.Info("Cron job started")
	response.SuccessResponse(ctx, response.ErrCodeSuccess, "Cron job started successfully")
}

// @Summary	  		Stop cron job
// @Description  	Stop cron job
// @Tags         	cron management
// @Accept       	json
// @Produce      	json
// @Success      	200  {object}  response.ResponseData
// @Failure      	500  {object}  response.ErrResponseData
// @Router       	/v1/cron/stop [post]
func (p *cCronJob) StopCronJob(ctx *gin.Context) {
	global.Cron.Stop()
	global.Logger.Info("Cron job stopped")
	response.SuccessResponse(ctx, response.ErrCodeSuccess, "Cron job stopped successfully")
}