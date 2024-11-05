package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tyagnii/gw-currency-wallet/internal/models"
	"net/http"
)

// GetBalance returns wallet balacce
//
//	@Summary      Shows wallet balance
//	@Description  shows user wallet balance
//	@Tags         accounts, wallets
//	@Param 		  Authorization header string true "JWT token"
//	@Accept       json
//	@Produce      json
//	@Success      200
//	@Failure      400
//	@Router       /api/v1/balance [get]
func (h *Handler) GetBalance(c *gin.Context) {
	var u = new(models.User)

	u.Username = c.Param("username")

	w, err := h.dbconn.GetBalance(c, *u)
	if err != nil {
		h.sLogger.Errorf("Error getting balance: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, w)
}
