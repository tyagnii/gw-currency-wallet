package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tyagnii/gw-currency-wallet/internal/db/models"
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
		h.sLogger.Errorf("Register error: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if u == (models.User{}) {
		h.sLogger.Errorf("Register error: User is empty")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User is empty"})
		return
	}

	if err := h.dbconn.CreateUser(c, u); err != nil {
		h.sLogger.Errorf("Register error: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, registerError)
		return
	}

	h.sLogger.Debugf("Register user successful: %v", u)
	c.JSON(http.StatusCreated, gin.H{"message": "User successfully registered"})
}
