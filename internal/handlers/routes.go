package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
)

func NewRouter() (*gin.Engine, error) {
	ctx := context.Background()
	h, err := NewHandler(ctx)
	if err != nil {
		return nil, err
	}
	r := gin.Default()
	r.POST("/api/v1/register", h.Register)
	r.POST("/api/v1/login", h.Login)
	r.GET("/api/v1/balance", h.GetBalance)
	r.POST("/api/v1/wallet/deposit", h.Deposit)
	r.POST("/api/v1/wallet/withdraw", h.Withdraw)
	r.GET("/api/v1/exchange/rates", h.GetRates)
	r.POST("/api/v1/exchange", h.Exchange)

	return r, nil
}
