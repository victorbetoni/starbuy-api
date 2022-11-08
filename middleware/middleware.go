package middleware

import (
	"net/http"
	"starbuy/authorization"
	"starbuy/util"

	"github.com/gin-gonic/gin"
)

func Authorize(next util.HandlerFunc) util.HandlerFunc {
	return func(c *gin.Context) (int, error) {
		if err := authorization.ValidateToken(c); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "unauthorized"})
			return 0, nil
		}

		if _, err := authorization.ExtractUser(c); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "invalid token"})
			return 0, nil
		}
		return next(c)
	}
}

func AbortOnError(next util.HandlerFunc) util.HandlerFunc {
	return func(c *gin.Context) (int, error) {
		if status, err := next(c); err != nil {
			c.AbortWithStatusJSON(status, gin.H{"status": false, "message": err.Error()})
			return 0, nil
		}
		return next(c)
	}
}

func Convert(handler util.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(c)
	}
}
