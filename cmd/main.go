package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/mrKrabsmr/infokeeper/docs"
	"github.com/mrKrabsmr/infokeeper/internal/api/v1/routes"
	config "github.com/mrKrabsmr/infokeeper/internal/config"
	"github.com/mrKrabsmr/infokeeper/pkg/db_connection"
	"log"
	"net/http"
	"os"
	"strconv"
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

	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	config := &config.Config{
		Server:          os.Getenv("SERVER_URL"),
		Database:        os.Getenv("DB_SERVER_URL"),
		MaxConn:         maxConn,
		MaxIdleConn:     maxIdleConn,
		MaxLifetimeConn: maxLifetimeConn,
	}

	route := gin.Default()

	conn := db_connection.NewPGConnection(config)

	route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Info-Keeper API!",
		})
	})

	v1 := route.Group("api/v1/info-keeper")
	{
		routes.InfoRoutes(*v1, conn)
	}

	routes.SwaggerRoutes(*route)

	route.Run(config.Server)
}
