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
	// var wallet models.Wallet
	var user models.User

	user.Username = c.Param("username")

	if err := c.BindJSON(&exchange); err != nil {
		h.sLogger.Errorf("Could not bind JSON: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Looking for rates in cache
	rate, yes := h.Cache.Get("rate")
	if yes {
		h.sLogger.Debugf("rate fetched from cache: %v", rate)

	} else {
		h.sLogger.Debugf("could not get rate from cache")
	}

	// TODO: get rates for exchange
	// 		check balance before withdraw
	// 		swithching between currencies ???

}
