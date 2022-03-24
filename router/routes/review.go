package routes

import (
	"net/http"
	"starbuy/controllers"
)

var Review = []Route{
	{
		URI:         "/{user}/reviews",
		Method:      http.MethodGet,
		RequireAuth: false,
		Action:      controllers.GetReviews,
	},
	{
		URI:         "/review/{id}",
		Method:      http.MethodGet,
		RequireAuth: false,
		Action:      controllers.GetReview,
	},
	{
		URI:         "/{item}/review",
		Method:      http.MethodPost,
		RequireAuth: true,
		Action:      controllers.PostReview,
	},
	{
		URI:         "/review",
		Method:      http.MethodPut,
		RequireAuth: true,
		Action:      controllers.PutReview,
	},
}
