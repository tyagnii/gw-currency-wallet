package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tyagnii/gw-currency-wallet/internal/models"
	"net/http"
)

// Deposit wallet with provided amount
//
//	@Summary      Deposit balance
//	@Description  deposit user wallet
//	@Tags         accounts
//	@Accept       json
//	@Produce      json
//	@Param
//	@Success      200  {object}  http.StatusOK
//	@Failure      400  {object}  http.StatusBadRequest
//	@Router       /api/v1/wallet/deposit [post]
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

}
