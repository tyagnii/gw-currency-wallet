package handlers

import "github.com/gin-gonic/gin"

func InitRoutes() {
	r := gin.Default()
	r.POST("/api/v1/register", Register)
	r.POST("/api/v1/login", Login)
	r.GET("/api/v1/balance", GetBalance)
	r.POST("/api/v1/wallet/deposit", Deposit)
	r.POST("/api/v1/wallet/withdraw", Withdraw)
	r.GET("/api/v1/exchange/rates", GetRates)
	r.POST("/api/v1/exchange", Exchange)

}
