package routes

import (
	"starbuy/controllers"

	"github.com/gin-gonic/gin"
)

var Item = []Route{
	{
		URI:         "/item/:id",
		RequireAuth: false,
		Action:      controllers.GetItem,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.GET(uri, hf)
		},
	},
	{
		URI:         "item/search/:query",
		RequireAuth: false,
		Action:      controllers.QueryItems,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.GET(uri, hf)
		},
	},
	{
		URI:         "/item/category/:id",
		RequireAuth: false,
		Action:      controllers.GetCategory,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.GET(uri, hf)
		},
	},
	{
		URI:         "/items",
		RequireAuth: false,
		Action:      controllers.GetAllItems,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.GET(uri, hf)
		},
	},
	{
		URI:         "/item",
		RequireAuth: true,
		Action:      controllers.PostItem,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.POST(uri, hf)
		},
	},
	{
		URI:         "/item/:id",
		RequireAuth: true,
		Action:      controllers.DeleteItem,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.DELETE(uri, hf)
		},
	},
}
