package routes

import "net/http"

var loginRoutes = []Route{
	{
		URI:    "login",
		Method: http.MethodPost,
		Action: func(w http.ResponseWriter, r *http.Request) {

		},
	},
	{
		URI:    "register",
		Method: http.MethodPost,
		Action: func(w http.ResponseWriter, r *http.Request) {

		},
	},
}
