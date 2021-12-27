package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type SignupCto struct {
	Email           string `json:"email" validate:"required,email"`
	Username        string `json:"username" validate:"required"`
	Password        string `json:"password" validate:"required,min=6,max=32,alphanum"`
	ConfirmPassword string `json:"confirmPassword" validate:"eqfield=Password,required"`
}

type LoginCto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func RegisterValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser SignupCto
		if err := c.ShouldBindBodyWith(&newUser, binding.JSON); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
		}
		validationErr := validator.New().Struct(newUser)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}

func LoginValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user LoginCto
		if err := c.ShouldBindBodyWith(&user, binding.JSON); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
		}
		validationErr := validator.New().Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
