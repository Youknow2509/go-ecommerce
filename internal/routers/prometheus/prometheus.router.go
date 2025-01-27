package prometheus

import (
	"github.com/Youknow2509/go-ecommerce/internal/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type PrometheusRouter struct {
}

// init router
func (p *PrometheusRouter) InitRouter(Router *gin.Engine) {
	// prometheusRouter := Router.Group("/prometheus")
	// prometheusRouter.Use(middlewares.PrometheusMiddleware())
	// {
	// 	prometheusRouter.GET("/metrics", gin.WrapH(promhttp.Handler()))
	// }
	Router.GET("/metrics", middlewares.PrometheusMiddleware(), gin.WrapH(promhttp.Handler()))
}
