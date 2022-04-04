package middleware

import (
	"net/http"
	"starbuy/authorization"

	"github.com/gin-gonic/gin"
)

func Authorize(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := authorization.ValidateToken(c); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		next(c)
	}
}
