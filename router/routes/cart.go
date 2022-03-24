package routes

import (
	"net/http"
	"starbuy/controllers"
)

var Cart = []Route{
	{
		URI:         "/cart",
		Method:      http.MethodGet,
		RequireAuth: true,
		Action:      controllers.QueryCart,
	},
	{
		URI:         "/cart",
		Method:      http.MethodPost,
		RequireAuth: true,
		Action:      controllers.PostCart,
	},
}
