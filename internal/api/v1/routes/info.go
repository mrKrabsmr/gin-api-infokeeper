package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mrKrabsmr/infokeeper/internal/api/v1/controllers"
	"github.com/mrKrabsmr/infokeeper/internal/app/dao"
	"github.com/mrKrabsmr/infokeeper/internal/app/services"
	"github.com/mrKrabsmr/infokeeper/pkg/db_connection"
)

func InfoRoutes(g gin.RouterGroup, conn *db_connection.PostgresDB) {
	db, err := conn.PostgreSQLConnection()
	if err != nil {
		panic("database connection error")
	}

	infoDAO := dao.NewInfoDAO(db)
	clientDAO := dao.NewClientDAO(db)

	serviceInfo := services.NewInfoService(infoDAO, clientDAO)

	controller := controllers.InfoController{Service: serviceInfo}

	g.GET("/", controller.Get)
	g.POST("/", controller.Post)
	g.PATCH("/", controller.Patch)
	g.DELETE("/", controller.Delete)

	g.GET("/count", controller.GetCount)
}
