package handlers

import (
	"fmt"

	"github.com/tyagnii/gw-currency-wallet/internal/models"
)

// withdrawCheck checks if there is enough balance to withdraw
// provided amount from wallet
func withdrawCheck(wallet models.Wallet, amount models.WithdrawReq) (bool, error) {
	switch amount.Currency {
	case "USD":
		if wallet.Balance.USD < amount.Amount {
			return false, fmt.Errorf("insufficient balance")
		}
	case "EUR":
		if wallet.Balance.EUR < amount.Amount {
			return false, fmt.Errorf("insufficient balance")
		}
	case "RUB":
		if wallet.Balance.RUB < amount.Amount {
			return false, fmt.Errorf("insufficient balance")
		}
	}

	return true, nil
}
