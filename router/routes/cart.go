package routes

import (
	"starbuy/controllers"

	"github.com/gin-gonic/gin"
)

var Cart = []Route{
	{
		URI:         "/cart",
		RequireAuth: true,
		Action:      controllers.QueryCart,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.GET(uri, hf)
		},
	},
	{
		URI:         "/cart",
		RequireAuth: true,
		Action:      controllers.PostCart,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.POST(uri, hf)
		},
	},
	{
		URI:         "/cart",
		RequireAuth: true,
		Action:      controllers.PostCart,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.DELETE(uri, hf)
		},
	},
}
