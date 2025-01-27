package prometheus

import (
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
	Router.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
