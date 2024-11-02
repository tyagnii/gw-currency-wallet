package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
)

// NewRouter create an instance of gin Router
func NewRouter() (*gin.Engine, error) {
	// Root context
	ctx := context.Background()

	// Init handlers
	h, err := NewHandler(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: Replace default router with custom one
	l := gin.Logger()
	r := gin.Default()
	r.Use(l)

	// Create an authrization group of endpoints
	// for and auth middleware
	authGroup := r.Group("/api/v1/")
	authGroup.Use(Auth())
	{
		r.GET("/api/v1/balance", h.GetBalance)
		r.POST("/api/v1/wallet/deposit", h.Deposit)
		r.POST("/api/v1/wallet/withdraw", h.Withdraw)
		r.GET("/api/v1/exchange/rates", h.GetRates)
		r.POST("/api/v1/exchange", h.Exchange)
	}

	r.POST("/api/v1/register", h.Register)
	r.POST("/api/v1/login", h.Login)

	return r, nil
}
