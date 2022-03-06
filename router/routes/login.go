package routes

import (
	"authentication-service/controllers"
	"net/http"
)

var loginRoutes = []Route{
	{
		URI:    "/login",
		Method: http.MethodPost,
		Action: controllers.Auth,
	},
	{
		URI:    "/register",
		Method: http.MethodPost,
		Action: controllers.Register,
	},
}
