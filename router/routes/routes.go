package routes

import (
	"authentication-service/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

// Route - Representação de todas as rotas da API
type Route struct {
	URI    string
	Method string
	Action func(http.ResponseWriter, *http.Request)
}

var routes = []Route{
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
	{
		URI:    "/user/{username}",
		Method: http.MethodGet,
		Action: controllers.QueryUser,
	},
	{
		URI:    "/item/{id}",
		Method: http.MethodGet,
		Action: controllers.QueryItem,
	},
}

func Configure(router *mux.Router) *mux.Router {
	for _, route := range routes {
		router.HandleFunc(route.URI, route.Action).Methods(route.Method)
	}

	return router
}
