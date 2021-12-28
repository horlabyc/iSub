package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	helper "github.com/horlabyc/iSub/helpers"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		clientToken := authHeader[len(BEARER_SCHEMA)+1:]
		if clientToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No authorization header provided"})
			return
		}
		claims, err := helper.ValidateToken(clientToken)
		if err != "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not authorized, bad/invalid token"})
			return
		}
		c.Set("email", claims.Email)
		c.Set("userId", claims.UserId.Hex())
		c.Set("username", claims.Username)
		c.Next()
	}
}
