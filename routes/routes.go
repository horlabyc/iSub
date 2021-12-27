package routes

import (
	controller "github.com/horlabyc/iSub/controllers"
	middleware "github.com/horlabyc/iSub/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router gin.IRouter) {
	RegisterAuthRoutes(router)
	RegisterUserRoutes(router)
}

func RegisterAuthRoutes(router gin.IRouter) {
	authRouter := router.Group("/auth")
	authRouter.POST("/register", middleware.RegisterValidator(), controller.Register())
	authRouter.POST("/login", middleware.LoginValidator(), controller.Login())
}

func RegisterUserRoutes(router gin.IRouter) {
	userRouter := router.Group("/users", middleware.Authenticate())
	userRouter.GET("/all", controller.GetAllUsers())
	userRouter.GET("/:userId", controller.GetUser())
}

// func RegisterSubscriptionRoutes(router gin.IRouter, db *gorm.DB) {
// 	db.AutoMigrate(&database.Subscription{})

// 	subscriptionRepository := repository.NewSubscriptionRepository(db)
// 	subscriptionHandler := controller.NewUserHandler(subscriptionRepository)

// 	userRouter := router.Group("/subscriptions")
// 	userRouter.GET("/", subscriptionHandler.GetAll)
// }
