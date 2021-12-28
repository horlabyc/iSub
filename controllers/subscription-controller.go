package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	logic "github.com/horlabyc/iSub/logic"
	model "github.com/horlabyc/iSub/models"
)

func CreateSubscription() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newSub model.Subscription
		if err := c.ShouldBindBodyWith(&newSub, binding.JSON); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		userId := c.GetString("userId")
		sub := logic.CreateSub(newSub, userId)
		c.JSON(http.StatusOK, gin.H{
			"message": "Subscription created successfully",
			"data":    sub,
		})
	}
}
