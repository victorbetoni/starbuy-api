package routes

import (
	"net/http"
	"starbuy/controllers"
)

var User = []Route{
	{
		URI:         "/login",
		Method:      http.MethodPost,
		RequireAuth: false,
		Action:      controllers.Auth,
	},
	{
		URI:         "/register",
		Method:      http.MethodPost,
		RequireAuth: false,
		Action:      controllers.Register,
	},
	{
		URI:         "/user/{username}",
		Method:      http.MethodGet,
		RequireAuth: false,
		Action:      controllers.GetUser,
	},
}
