package controller

import (
	"net/http"

	logic "github.com/horlabyc/iSub/logic"

	model "github.com/horlabyc/iSub/models"
	"github.com/horlabyc/iSub/repository"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type UserHandler struct {
	repository *repository.UserRepository
}

func NewUserHandler(repository *repository.UserRepository) *UserHandler {
	return &UserHandler{
		repository: repository,
	}
}

func (handler UserHandler) GetAll(c *gin.Context) {
	var users []model.User = handler.repository.FindAll()
	c.JSON(200, gin.H{"status": "Ok", "data": users})
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
