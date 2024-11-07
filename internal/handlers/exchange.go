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
	var exchangeReq models.ExchangeReq
	var wallet models.Wallet
	var user models.User
	var err error

	user.Username = c.Param("username")

	if err := c.BindJSON(&exchangeReq); err != nil {
		h.sLogger.Errorf("Could not bind JSON: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Looking for rates in cache
	rate, yes := h.Cache.Get("rate")
	if yes {
		h.sLogger.Debugf("rate fetched from cache: %v", rate)

		exchangeReq.Rate = rate.(models.Currency)
	} else {
		h.sLogger.Debugf("Exchange: could not get rate from cache")
	}

	wallet, err = h.dbconn.GetWalletByUsername(c, user.Username)
	if err != nil {
		h.sLogger.Errorf("Exchange: Could not get user: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	wallet, err = h.dbconn.Exchange(c, wallet, exchangeReq)
	if err != nil {
		h.sLogger.Errorf("Exchange: Could not exchange: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, wallet)
}
