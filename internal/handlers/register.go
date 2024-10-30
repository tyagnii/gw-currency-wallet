package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register adds new user account
//
//	@Summary      Register user
//	@Description  register new users
//	@Tags         accounts
//	@Accept       json
//	@Produce      json
//	@Param
//	@Success      201  {object}  http.StatusOK
//	@Failure      400  {object}  http.StatusBadRequest
//	@Router       /api/v1/register [post]
func Register(c *gin.Context) {
	c.Status(http.StatusOK)
}
