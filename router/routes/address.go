package routes

import (
	"net/http"
	"starbuy/controllers"
)

var Address = []Route{
	{
		URI:         "/{user}/address",
		Method:      http.MethodGet,
		RequireAuth: true,
		Action:      controllers.GetAddresses,
	},
	{
		URI:         "/{user}/address",
		Method:      http.MethodPost,
		RequireAuth: true,
		Action:      controllers.PostCart,
	},
	{
		URI:         "/{user}/address/{id}",
		Method:      http.MethodGet,
		RequireAuth: true,
		Action:      controllers.PostCart,
	},
}
