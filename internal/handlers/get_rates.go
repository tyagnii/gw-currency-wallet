package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tyagnii/gw-currency-wallet/internal/models"
	"github.com/tyagnii/gw-exchanger/gen/exchanger/v1"
	"net/http"
)

// todo: implement
func (h *Handler) GetRates(c *gin.Context) {
	var resp models.Rates

	eResp, err := h.eClient.GetExchangeRates(c, &exchanger.Empty{})
	if err != nil {
		// todo: get rates from cache
		h.sLogger.Errorf("Could not get exchange rates: %s", err.Error())
	}

	if eResp == nil {
		h.sLogger.Errorf("No exchange rates found")
		return
	}

	mapRates := eResp.GetRates()
	resp.Rates.USD = mapRates["USD"]
	resp.Rates.EUR = mapRates["EUR"]
	resp.Rates.RUB = mapRates["RUB"]

	c.JSON(http.StatusOK, resp)
}
