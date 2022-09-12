package middleware

import (
	"net/http"
	"starbuy/authorization"
	"starbuy/util"

	"github.com/gin-gonic/gin"
)

func Authorize(next util.HandlerFuncError) util.HandlerFuncError {
	return func(c *gin.Context) error {
		if err := authorization.ValidateToken(c); err != nil {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "unauthorized"})
		}

		if _, err := authorization.ExtractUser(c); err != nil {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "invalid token"})
		}
		return next(c)
	}
}

func AbortOnError(handler util.HandlerFuncError) util.HandlerFuncError {
	return func(c *gin.Context) error {
		if err := handler(c); err != nil {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		}
		return nil
	}
}

func Convert(handler util.HandlerFuncError) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(c)
	}
}
