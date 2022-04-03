package middleware

import (
	"net/http"
	"starbuy/authorization"
	"starbuy/responses"

	"github.com/gin-gonic/gin"
)

func Authorize(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := authorization.ValidateToken(c); err != nil {
			responses.Error(w, http.StatusUnauthorized, err)
			return
		}
		next(c)
	}
}
