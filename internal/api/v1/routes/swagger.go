package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SwaggerRoutes(e gin.Engine) {
	e.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
