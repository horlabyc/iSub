package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	logic "github.com/horlabyc/iSub/logic"
	model "github.com/horlabyc/iSub/models"
	"go.mongodb.org/mongo-driver/mongo"
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

func ActivateSubscription() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.GetString("userId")
		subId := c.Param("subId")
		if subId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Subscription id is required"})
			c.Abort()
			return
		}
		sub, error := logic.ActivateSub(userId, subId)
		if error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Subscription activated successfully",
			"data":    sub,
		})
	}
}

func CancelSubscription() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.GetString("userId")
		subId := c.Param("subId")
		if subId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Subscription id is required"})
			c.Abort()
			return
		}
		sub, error := logic.CancelSub(userId, subId)
		if error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Subscription cancelled successfully",
			"data":    sub,
		})
	}
}

func GetSubscription() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.GetString("userId")
		subId := c.Param("subId")
		if subId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Subscription id is required"})
			c.Abort()
			return
		}
		sub, error := logic.GetSub(userId, subId)
		if error == mongo.ErrNoDocuments {
			c.JSON(http.StatusOK, gin.H{
				"message": "Get Subscription",
				"data":    map[string]interface{}{},
			})
			return
		}
		if error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Get Subscription",
			"data":    sub,
		})
	}
}

func GetSubscriptions() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.GetString("userId")
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			page = 1
		}
		limit, err := strconv.Atoi(c.Query("limit"))
		if err != nil || limit < 1 {
			limit = 15
		}
		skip := (page - 1) * limit
		sub, error := logic.GetAllSub(page, limit, skip, userId)
		if error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Get Subscriptions",
			"data":    sub,
		})
	}
}
