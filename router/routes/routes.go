package routes

import (
	"net/http"
	"starbuy/middleware"

	"github.com/gorilla/mux"
)

// Route - Representação de todas as rotas da API
type Route struct {
	URI         string
	Method      string
	RequireAuth bool
	Action      func(http.ResponseWriter, *http.Request)
}

func Configure(router *mux.Router) *mux.Router {
	var routes [][]Route
	routes = append(routes, Item)
	routes = append(routes, User)
	routes = append(routes, Cart)
	routes = append(routes, Review)
	routes = append(routes, Order)
	routes = append(routes, Address)

	for _, x := range routes {
		for _, route := range x {
			if route.RequireAuth {
				router.HandleFunc(route.URI, middleware.Authorize(route.Action)).Methods(route.Method)
			} else {
				router.HandleFunc(route.URI, route.Action).Methods(route.Method)
			}
		}
	}

	return router
}
