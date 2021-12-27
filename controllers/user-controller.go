package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	logic "github.com/horlabyc/iSub/logic"
	"github.com/horlabyc/iSub/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	model "github.com/horlabyc/iSub/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userId")
		user, err := logic.GetUser(userId)
		fmt.Println("error", err == mongo.ErrNoDocuments)
		if err == mongo.ErrNoDocuments {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "User does not exist"})
			return
		}
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "Ok", "data": map[string]interface{}{
			"Id":        user.Id,
			"Email":     user.Email,
			"Username":  user.Username,
			"CreatedAt": user.CreatedAt,
		}})
	}
}

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser model.User
		if err := c.ShouldBindBodyWith(&newUser, binding.JSON); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
		}
		u, error := logic.RegisterUser(newUser)
		if error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "Okay",
			"message": "User registration is successful",
			"data": map[string]interface{}{
				"id":        u.Id,
				"email":     u.Email,
				"username":  u.Username,
				"createdAt": u.CreatedAt,
			}})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindBodyWith(&user, binding.JSON); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
		}
		u, token, refreshToken, error := logic.LoginUser(user)
		if error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "Okay",
			"message": "User login is successful",
			"data": map[string]interface{}{
				"id":           u.UserId,
				"email":        u.Email,
				"username":     u.Username,
				"token":        token,
				"refreshToken": refreshToken,
			},
		})
	}
}

func GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			page = 1
		}
		limit, err := strconv.Atoi(c.Query("limit"))
		if err != nil || limit < 1 {
			limit = 15
		}
		skip := (page - 1) * limit
		match := bson.D{{"$match", bson.D{{}}}}
		group := bson.D{{"$group", bson.D{
			{"_id", bson.D{{"_id", "null"}}},
			{"total", bson.D{{"$sum", 1}}},
			{"data", bson.D{{"$push", "$$ROOT"}}},
		}}}
		projectStage := bson.D{
			{"$project", bson.D{
				{"_id", 0},
				{"totalCount", 1},
				{"users", bson.D{{"$slice", []interface{}{"$data", skip, limit}}}},
			}},
		}
		result, err := repository.FindAllUsers(match, group, projectStage)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing users"})
		}
		log.Println(result)
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var allusers []bson.M
		if err = result.All(ctx, &allusers); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allusers[0])
	}
}
