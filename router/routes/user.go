package routes

import (
	"starbuy/controllers"

	"github.com/gin-gonic/gin"
)

var User = []Route{
	{
		RequireAuth: false,
		URI:         "/login",
		Action:      controllers.Auth,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.POST(uri, hf)
		},
	},
	{
		RequireAuth: false,
		URI:         "/register",
		Action:      controllers.Register,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.POST(uri, hf)
		},
	},
	{
		RequireAuth: false,
		URI:         "/user/:user",
		Action:      controllers.GetUser,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.GET(uri, hf)
		},
	},
	{
		URI:         "/received_orders",
		RequireAuth: true,
		Action:      controllers.GetReceivedOrders,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.GET(uri, hf)
		},
	},
	{
		URI:         "/profile_picture",
		RequireAuth: true,
		Action:      controllers.PostUserProfilePicture,
		Assign: func(engine *gin.Engine, handlerFunc gin.HandlerFunc, s string) {
			engine.POST(s, handlerFunc)
		},
	},
}
