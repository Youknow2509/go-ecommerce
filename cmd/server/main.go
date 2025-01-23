package main

import (
	"strconv"

	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/Youknow2509/go-ecommerce/internal/initialize"
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swaggo/files" // swagger embed files
	_ "github.com/Youknow2509/go-ecommerce/cmd/swag/docs" // docs is generated by Swag CLI, you have to import it.
)

// @title           API Documentation Go Ecommerce
// @version         1.0.0
// @description     This is a sample server celler server.
// @termsOfService  github.com/Youknow2509/go-ecommerce/

// @contact.name   TEAM V
// @contact.url    github.com/Youknow2509/go-ecommerce/
// @contact.email  lytranvinh.work@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8082
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	r := initialize.Run()

	// use swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := strconv.Itoa(global.Config.Server.Port)
	r.Run(":" + port)
}
