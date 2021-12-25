package routes

import (
	controller "github.com/horlabyc/iSub/controllers"
	validations "github.com/horlabyc/iSub/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router gin.IRouter) {

	// userRepository := repository.NewUserRepository(db)
	// userHandler := controller.NewUserHandler()

	userRouter := router.Group("/users")
	// userRouter.GET("/", userHandler.GetAll)
	userRouter.POST("/register", validations.RegisterValidator(), controller.Register())
	userRouter.POST("/login", validations.LoginValidator(), controller.Login())
}

// func RegisterSubscriptionRoutes(router gin.IRouter, db *gorm.DB) {
// 	db.AutoMigrate(&database.Subscription{})

// 	subscriptionRepository := repository.NewSubscriptionRepository(db)
// 	subscriptionHandler := controller.NewUserHandler(subscriptionRepository)

// 	userRouter := router.Group("/subscriptions")
// 	userRouter.GET("/", subscriptionHandler.GetAll)
// }
