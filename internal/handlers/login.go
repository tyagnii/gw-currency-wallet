package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tyagnii/gw-currency-wallet/internal/db/models"
	"github.com/tyagnii/gw-currency-wallet/internal/token"
	"net/http"
)

// Login authorizes  adds new user account
//
//	@Summary      Authorize existing user
//	@Description  authorize users
//	@Tags         accounts
//	@Accept       json
//	@Produce      json
//	@Success      200
//	@Failure      400
//	@Router       /api/v1/login [post]
func (h *Handler) Login(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		h.sLogger.Errorf("Login error: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwtToken, err := token.NewToken(user)
	if err != nil {
		h.sLogger.Errorf("Login error: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.sLogger.Debugf("Login success")
	c.JSON(http.StatusOK, gin.H{"token": jwtToken})
}
