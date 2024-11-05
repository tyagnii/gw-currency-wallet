package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tyagnii/gw-currency-wallet/internal/models"
	"net/http"
)

// Deposit wallet with provided amount
//
// @Summary      Deposit balance
// @Description  deposit user wallet
// @Tags         accounts
// @Param 		 Authorization header string true "JWT token"
// @Param		 amount body models.DepositReq true "Deposit query in json format"
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Router       /api/v1/wallet/deposit [post]
func (h *Handler) Deposit(c *gin.Context) {
	var dq models.DepositReq
	var w models.Wallet

	if err := c.ShouldBindJSON(&dq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	switch dq.Currency {
	case "USD":
		w.Balance.USD += dq.Amount
	case "EUR":
		w.Balance.EUR += dq.Amount
	case "RUB":
		w.Balance.RUB += dq.Amount
	}

	if err := h.dbconn.Deposit(c, w); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account topped up successfully", "new_balance": w.Balance})

}
