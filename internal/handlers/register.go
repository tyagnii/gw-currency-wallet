package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tyagnii/gw-currency-wallet/internal/models"
	"net/http"
)

var registerError = gin.H{"error": "Username or email already exists"}

// Register adds new user account
//
//	@Summary      Register user
//	@Description  register new users
//	@Tags         accounts
//	@Accept       json
//	@Produce      json
//	@Success      201
//	@Failure      400
//	@Router       /api/v1/register [post]
func (h *Handler) Register(c *gin.Context) {
	var u models.User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.dbconn.CreateUser(c, u); err != nil {
		c.JSON(http.StatusBadRequest, registerError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User successfully registered"})
}
