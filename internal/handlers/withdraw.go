package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tyagnii/gw-currency-wallet/internal/models"
	"net/http"
)

// Withdraw wallet with provided amount
// @Summary      Withdraw amount
// @Description  withdraw provided amount from user wallet
// @Tags         accounts
// @Param 		 Authorization header string true "JWT token"
// @Param		 amount body models.WithdrawReq true "Withdraw query in json format"
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Router       /api/v1/wallet/withdraw [post]
func (h *Handler) Withdraw(c *gin.Context) {
	var user = models.User{}
	var wallet = models.Wallet{}
	var amount = models.WithdrawReq{}

	user.Username = c.Param("username")

	wallet, err := h.dbconn.GetBalance(c, user)
	if err != nil {
		h.sLogger.Errorf("Withdraw error: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = c.BindJSON(&amount)
	if err != nil {
		h.sLogger.Errorf("Withdraw error: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	yes, err := withdrawCheck(wallet, amount)
	if err != nil {
		h.sLogger.Errorf("Withdraw error: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !yes {
		h.sLogger.Errorf("Withdraw error: %v", fmt.Errorf("insufficient balance"))
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "insufficient balance"})
		return
	}

	err = h.dbconn.Withdraw(c, wallet)
	if err != nil {
		h.sLogger.Errorf("Withdraw error: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.sLogger.Infof("Withdraw success")
	c.JSON(http.StatusOK, wallet)
}
