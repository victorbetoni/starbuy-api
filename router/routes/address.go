package routes

import (
	"starbuy/controllers"

	"github.com/gin-gonic/gin"
)

var Address = []Route{
	{
		URI:         "/:user/address",
		RequireAuth: true,
		Action:      controllers.GetAddresses,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.GET(uri, hf)
		},
	},
	{
		URI:         "/:user/address",
		RequireAuth: true,
		Action:      controllers.PostAddress,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.POST(uri, hf)
		},
	},
	{
		URI:         "/:user/address/:id",
		RequireAuth: true,
		Action:      controllers.GetAddress,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.GET(uri, hf)
		},
	},
}
