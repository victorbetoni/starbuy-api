package routes

import (
	"starbuy/controllers"

	"github.com/gin-gonic/gin"
)

var Review = []Route{
	{
		URI:         "/:user/reviews",
		RequireAuth: false,
		Action:      controllers.GetReviews,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.GET(uri, hf)
		},
	},
	{
		URI:         "/review/:id",
		RequireAuth: false,
		Action:      controllers.GetReview,
		Assign: func(e *gin.Engine, hf gin.HandlerFunc, uri string) {
			e.GET(uri, hf)
		},
	},
	{
		URI:         "/:item/review",
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
}
