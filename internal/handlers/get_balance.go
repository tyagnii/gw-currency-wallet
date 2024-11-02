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
//	@Accept       json
//	@Produce      json
//	@Param
//	@Success      200  {object}  http.StatusOK
//	@Failure      400  {object}  http.StatusBadRequest
//	@Router       /api/v1/balance [get]
func (h *Handler) GetBalance(c *gin.Context) {
	var u models.User

	u.Username = c.Param("username")

	w, err := h.dbconn.GetBalance(c, u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, w)
}
