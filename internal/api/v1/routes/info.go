package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mrKrabsmr/infokeeper/internal/api/v1/controllers"
)

func InfoRoutes(g gin.RouterGroup) {
	g.GET("/", controllers.Get)
	g.POST("/", controllers.Post)
	g.PATCH("/", controllers.Patch)
	g.DELETE("/", controllers.Delete)

	g.GET("/count", controllers.GetCount)
}
