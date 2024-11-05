package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/tyagnii/gw-currency-wallet/internal/models"
	"net/http"
)

func withdrawCheck(wallet models.Wallet, amount models.WithdrawReq) (bool, error) {
	switch amount.Currency {
	case "USD":
		if wallet.Balance.USD < amount.Amount {
			return false, errors.New("insufficient balance")
		}
	case "EUR":
		if wallet.Balance.EUR < amount.Amount {
			return false, errors.New("insufficient balance")
		}
	case "RUB":
		if wallet.Balance.RUB < amount.Amount {
			return false, errors.New("insufficient balance")
		}
	}

	return true, nil
}

func (h *Handler) Withdraw(c *gin.Context) {
	var user = models.User{}
	var wallet = models.Wallet{}
	var amount = models.WithdrawReq{}

	user.Username = c.Param("username")

	wallet, err := h.dbconn.GetBalance(c, user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = c.BindJSON(&amount)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	yes, err := withdrawCheck(wallet, amount)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !yes {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "insufficient balance"})
		return
	}

	err = h.dbconn.Withdraw(c, wallet)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, wallet})
}
