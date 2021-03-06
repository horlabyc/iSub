package routes

import (
	controller "github.com/horlabyc/iSub/controllers"
	middleware "github.com/horlabyc/iSub/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router gin.IRouter) {
	RegisterAuthRoutes(router)
	RegisterUserRoutes(router)
	RegisterSubscriptionRoutes(router)
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

func RegisterSubscriptionRoutes(router gin.IRouter) {
	subscriptionRouter := router.Group("/subscriptions", middleware.Authenticate())
	subscriptionRouter.GET("/all", controller.GetSubscriptions())
	subscriptionRouter.GET("/:subId", controller.GetSubscription())
	subscriptionRouter.POST("/create", middleware.CreateSubscriptionValidator(), controller.CreateSubscription())
	subscriptionRouter.PUT("/:subId/activate", controller.ActivateSubscription())
	subscriptionRouter.PUT("/:subId/cancel", controller.CancelSubscription())
}
