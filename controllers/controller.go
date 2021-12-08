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

func (handler UserHandler) Register(c *gin.Context) {
	var newUser model.User
	if err := c.ShouldBindBodyWith(&newUser, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
	}
	u, token, refreshToken, err := logic.RegisterUser(handler.repository, newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "Okay",
		"message": "User registration is successful",
		"data": map[string]interface{}{
			"id":           u.ID,
			"email":        u.Email,
			"username":     u.Username,
			"createdAt":    u.CreatedAt,
			"token":        token,
			"refreshToken": refreshToken,
		}})
}
