package routes

import (
	controller "github.com/horlabyc/iSub/controllers"
	validations "github.com/horlabyc/iSub/middlewares"
	model "github.com/horlabyc/iSub/models"

	"github.com/horlabyc/iSub/repository"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func RegisterUserRoutes(router gin.IRouter, db *gorm.DB) {
	db.AutoMigrate(&model.User{})

	userRepository := repository.NewUserRepository(db)
	userHandler := controller.NewUserHandler(userRepository)

	userRouter := router.Group("/users")
	userRouter.GET("/", userHandler.GetAll)
	userRouter.POST("/register", validations.RegisterValidator(), userHandler.Register)
}

// func RegisterSubscriptionRoutes(router gin.IRouter, db *gorm.DB) {
// 	db.AutoMigrate(&database.Subscription{})

// 	subscriptionRepository := repository.NewSubscriptionRepository(db)
// 	subscriptionHandler := controller.NewUserHandler(subscriptionRepository)

// 	userRouter := router.Group("/subscriptions")
// 	userRouter.GET("/", subscriptionHandler.GetAll)
// }
