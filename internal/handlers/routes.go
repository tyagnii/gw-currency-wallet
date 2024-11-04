package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tyagnii/gw-currency-wallet/internal/middleware"
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
	r.POST("/api/v1/register", h.Register)
	r.POST("/api/v1/login", h.Login)
	r.Use(l)

	// Create an authrization group of endpoints
	// for and auth middleware
	authGroup := r.Group("/api/v1").Use(middleware.Auth())
	authGroup.GET("/balance", h.GetBalance)
	authGroup.POST("/wallet/deposit", h.Deposit)
	authGroup.POST("/wallet/withdraw", h.Withdraw)
	authGroup.GET("/exchange/rates", h.GetRates)
	authGroup.POST("/exchange", h.Exchange)

	return r, nil
}
