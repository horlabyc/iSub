package main

import (
	helper "github.com/horlabyc/iSub/helpers"
	"github.com/horlabyc/iSub/routes"

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

	apiRoute := router.Group("/api")
	routes.RegisterUserRoutes(apiRoute)
	// routes.RegisterSubscriptionRoutes(apiRoute, database.DB)
	router.Run(":" + port)
}
