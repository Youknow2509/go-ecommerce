package product

import (
	"github.com/Youknow2509/go-ecommerce/internal/controller/product"
	"github.com/gin-gonic/gin"
)

type ClothingRouter struct {
}

func (pr *ClothingRouter) InitClothingRouter(Router *gin.RouterGroup) {

	// public router
	ClothingRouterPublic := Router.Group("/clothing")
	// ClothingRouterPublic.Use(middlewares.AuthenMiddleware())
	// ClothingRouterPublic.Use(Limmited())
	// ClothingRouterPublic.Use(Authen())
	// ClothingRouterPublic.Use(Permission())
	{
		ClothingRouterPublic.POST("/create", product.Product.CreateClothing)
	}
}
