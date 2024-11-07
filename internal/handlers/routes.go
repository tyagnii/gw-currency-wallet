package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tyagnii/gw-currency-wallet/internal/middleware"
	"go.uber.org/zap"
)

// NewRouter create an instance of gin Router
func NewRouter(sLogger *zap.SugaredLogger) (*gin.Engine, error) {
	// Root context
	ctx := context.Background()

	// Init handlers
	h, err := NewHandler(ctx, sLogger)
	if err != nil {
		sLogger.Errorf("Could not initialize handler: %v", err)
		return nil, err
	}

	// TODO: Replace default router with custom one
	r := gin.Default()
	r.POST("/api/v1/register", h.Register)
	r.POST("/api/v1/login", h.Login)

	// Create an authrization group of endpoints
	// for and auth middleware
	authGroup := r.Group("/api/v1").Use(middleware.Auth())
	authGroup.GET("/balance", h.GetBalance)
	authGroup.POST("/wallet/deposit", h.Deposit)
	authGroup.POST("/wallet/withdraw", h.Withdraw)
	authGroup.GET("/exchange/rates", h.GetRates)
	authGroup.POST("/exchange", h.Exchange)

	sLogger.Debugf("Routes initialized")
	return r, nil
}
