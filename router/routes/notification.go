package routes

import (
	"starbuy/controllers"

	"github.com/gin-gonic/gin"
)

var Notification = []Route{
	{
		URI:         "/notification",
		RequireAuth: true,
		Action:      controllers.PostNotification,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.POST(uri, hf)
		},
	},
	{
		URI:         "/notification/:id",
		RequireAuth: true,
		Action:      controllers.GetNotification,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.GET(uri, hf)
		},
	},
	{
		URI:         "/notification/",
		RequireAuth: true,
		Action:      controllers.GetNotifications,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.GET(uri, hf)
		},
	},
}
