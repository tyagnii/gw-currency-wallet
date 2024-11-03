package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tyagnii/gw-currency-wallet/pkg/token"
	"net/http"
)

// Auth middleware for gin router. Provides authentication mechanism via JWT
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		t := c.GetHeader("token")
		if flag, _ := token.ParseToken(t); !flag {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}
		c.Next()

	}
}
