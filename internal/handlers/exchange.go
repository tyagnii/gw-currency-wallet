package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tyagnii/gw-currency-wallet/internal/models"
	"net/http"
)

// Exchange one currency to another with provided amount
//
//	@Summary      Exchanger endpoint
//	@Description  exchange one currency to anoter
//	@Tags         exchange
//	@Accept       json
//	@Produce      json
//	@Param
//	@Success      200  {object}  http.StatusOK
//	@Failure      400  {object}  http.StatusBadRequest
//	@Router       /api/v1/exchange [post]
func (h *Handler) Exchange(c *gin.Context) {
	var exchange models.ExchangeReq
	//var wallet models.Wallet

	if err := c.BindJSON(&exchange); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// TODO: get rates for exchange
	// 		check balance before withdraw
	// 		swithching between currencies ???

}
