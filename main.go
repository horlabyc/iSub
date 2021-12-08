package main

import (
	"github.com/horlabyc/iSub/database"
	helper "github.com/horlabyc/iSub/helpers"
	routes "github.com/horlabyc/iSub/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := helper.LoadEnv("PORT")
	if port == "" {
		port = "80"
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())
	database.ConnectDB()
	defer database.DB.Close()

	apiRoute := router.Group("/api")
	routes.RegisterUserRoutes(apiRoute, database.DB)
	// routes.RegisterSubscriptionRoutes(apiRoute, database.DB)
	router.Run(":" + port)
}
