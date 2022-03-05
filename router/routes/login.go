package routes

import (
	"authentication-service/controllers"
	"net/http"
)

var loginRoutes = []Route{
	{
		URI:    "login",
		Method: http.MethodPost,
		Action: controllers.Login,
	},
	{
		URI:    "register",
		Method: http.MethodPost,
		Action: func(w http.ResponseWriter, r *http.Request) {

		},
	},
}
