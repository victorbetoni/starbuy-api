package routes

import (
	"starbuy/controllers"

	"github.com/gin-gonic/gin"
)

var Order = []Route{

	{
		URI:         "/order",
		RequireAuth: true,
		Action:      controllers.PostPurchase,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.POST(uri, hf)
		},
	},
	{
		URI:         "/order/:id",
		RequireAuth: true,
		Action:      controllers.PostPurchase,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.POST(uri, hf)
		},
	},
	{
		URI:         "/orders",
		RequireAuth: true,
		Action:      controllers.GetPurchases,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.GET(uri, hf)
		},
	},
	{
		URI:         "/order/:id",
		RequireAuth: true,
		Action:      controllers.GetPurchase,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.GET(uri, hf)
		},
	},
}
