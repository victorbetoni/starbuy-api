package routes

import (
	"authentication-service/controllers"
	"net/http"
)

var loginRoutes = []Route{
	{
		URI:    "/login",
		Method: http.MethodPost,
		Action: controllers.Login,
	},
	{
		URI:    "/register",
		Method: http.MethodPost,
		Action: controllers.Register,
	},
}
