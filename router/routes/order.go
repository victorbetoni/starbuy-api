package routes

import (
	"starbuy/controllers"

	"github.com/gin-gonic/gin"
)

var Order = []Route{

	{
		URI:         "/order",
		RequireAuth: true,
		Action:      controllers.CreateOrder,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.POST(uri, hf)
		},
	},
	{
		URI:         "/order/:id",
		RequireAuth: true,
		Action:      controllers.GetOrder,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.GET(uri, hf)
		},
	},
	{
		URI:         "/orders",
		RequireAuth: true,
		Action:      controllers.GetOrders,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.GET(uri, hf)
		},
	},
}
