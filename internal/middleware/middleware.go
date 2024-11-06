package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/tyagnii/gw-currency-wallet/internal/token"
	"net/http"
)

// Auth middleware for gin router. Provides authentication mechanism via JWT
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		t := c.GetHeader("Authorization")
		parsedToken, claims, err := token.ParseToken(t)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if !token.ValidateToken(parsedToken) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.AddParam("username", claims.Username)
		c.Next()
	}
}
