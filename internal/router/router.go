package router

import (
	"mangosteen/internal/controller"

	"mangosteen/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swag
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func New() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.Version = "1.0"

	r.GET("/api/v1/ping", controller.Ping)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
