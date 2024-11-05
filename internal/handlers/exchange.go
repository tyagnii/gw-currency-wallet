package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tyagnii/gw-currency-wallet/internal/models"
	"net/http"
)

// Exchange one currency to another with provided amount
//
//	@Summary      Exchanger endpoint
//	@Description  exchange one currency to another
//	@Tags         exchange
//	@Param 		 Authorization header string true "JWT token"
//	@Param		  amount body models.ExchangeReq true "Exchange query in json format"
//	@Accept       json
//	@Produce      json
//	@Success      200
//	@Failure      400
//	@Router       /api/v1/exchange [post]
func (h *Handler) Exchange(c *gin.Context) {
	var exchange models.ExchangeReq
	//var wallet models.Wallet

	if err := c.BindJSON(&exchange); err != nil {
		h.sLogger.Errorf("Could not bind JSON: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: get rates for exchange
	// 		check balance before withdraw
	// 		swithching between currencies ???

}
