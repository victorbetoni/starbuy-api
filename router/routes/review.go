package routes

import (
	"starbuy/controllers"

	"github.com/gin-gonic/gin"
)

var Review = []Route{
	{
		URI:         "/user/reviews/received",
		RequireAuth: true,
		Action:      controllers.GetUserReceivedReviews,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.GET(uri, hf)
		},
	},
	{
		URI:         "/item/reviews/:item",
		RequireAuth: false,
		Action:      controllers.GetItemReviews,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.GET(uri, hf)
		},
	},
	{
		URI:         "/user/reviews",
		RequireAuth: true,
		Action:      controllers.GetUserReviews,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.GET(uri, hf)
		},
	},
	{
		URI:         "/review",
		RequireAuth: false,
		Action:      controllers.GetReview,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.GET(uri, hf)
		},
	},
	{
		URI:         "/review",
		RequireAuth: true,
		Action:      controllers.PostReview,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.POST(uri, hf)
		},
	},
	{
		URI:         "/review",
		RequireAuth: true,
		Action:      controllers.PutReview,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.PUT(uri, hf)
		},
	},
	{
		URI:         "/review/:item",
		RequireAuth: true,
		Action:      controllers.DeleteReview,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.DELETE(uri, hf)
		},
	},
}
