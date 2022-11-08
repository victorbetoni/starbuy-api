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

func CORS(next util.HandlerFunc) util.HandlerFunc {
	return func(c *gin.Context) (int, error) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		return next(c)
	}
}

func AbortOnError(next util.HandlerFunc) util.HandlerFunc {
	return func(c *gin.Context) (int, error) {
		if status, err := next(c); err != nil {
			c.AbortWithStatusJSON(status, gin.H{"status": false, "message": err.Error()})
			return 0, nil
		}
		return 0, nil
	}
}

func Convert(handler util.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(c)
	}
}
