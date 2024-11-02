package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tyagnii/gw-currency-wallet/internal/models"
	"github.com/tyagnii/gw-currency-wallet/pkg/token"
	"net/http"
)

// Login authorizes  adds new user account
//
//	@Summary      Authorize existing user
//	@Description  authorize users
//	@Tags         accounts
//	@Accept       json
//	@Produce      json
//	@Param
//	@Success      200  {object}  http.StatusOK
//	@Failure      400  {object}  http.StatusBadRequest
//	@Router       /api/v1/login [post]
func (h *Handler) Login(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: generate token
	jwtToken, err := token.NewToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": jwtToken})
}
