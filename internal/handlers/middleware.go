package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tyagnii/gw-currency-wallet/pkg/token"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		t := c.GetHeader("token")
		if flag, _ := token.ParseToken(t); !flag {
			c.JSON(http.StatusUnauthorized, gin.H{})
		}

	}
}
