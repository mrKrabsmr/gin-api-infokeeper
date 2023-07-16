package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/mrKrabsmr/infokeeper/docs"
	"github.com/mrKrabsmr/infokeeper/internal/api/v1/routes"
	"log"
	"net/http"
	"os"
)

// @title Info-Keeper
// @version 1.0
// @description You can keep even the most intimate secrets
// @host localhost:5050
// @contact.name Info-Keeper API Support
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	route := gin.Default()

	route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Info-Keeper API!",
		})
	})

	v1 := route.Group("api/v1/info-keeper")
	{
		routes.InfoRoutes(*v1)
	}

	routes.SwaggerRoutes(*route)

	route.Run(os.Getenv("SERVER_URL"))
}
