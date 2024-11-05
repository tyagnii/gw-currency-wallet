package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tyagnii/gw-currency-wallet/internal/models"
	"github.com/tyagnii/gw-exchanger/gen/exchanger/v1"
	"net/http"
)

// todo: implement
func (h *Handler) GetRates(c *gin.Context) {
	var resp models.Rates

	e, err := h.eClient.GetExchangeRates(c, &exchanger.Empty{})
	if err != nil {
		// todo: get rates from cache
		fmt.Println(err.Error())
	}

	if e == nil {
		// todo: nil?
		fmt.Println(e, "exchange rates")
	}
	mapRates := e.GetRates()
	resp.Rates.USD = mapRates["USD"]
	resp.Rates.EUR = mapRates["EUR"]
	resp.Rates.RUB = mapRates["RUB"]

	c.JSON(http.StatusOK, resp)
}
