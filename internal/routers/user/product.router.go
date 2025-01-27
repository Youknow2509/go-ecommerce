package user

import (
	"github.com/Youknow2509/go-ecommerce/internal/middlewares"
	"github.com/gin-gonic/gin"
)

type ProductRouter struct {
}

func (pr *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	// public router
	productRouterPublic := Router.Group("/product")
	productRouterPublic.Use(middlewares.PrometheusMiddleware())
	{
		productRouterPublic.GET("/search")
		productRouterPublic.GET("/detail/:id")
	}

	// private router

}
