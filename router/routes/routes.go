package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route - Representação de todas as rotas da API
type Route struct {
	URI    string
	Method string
	Action func(http.ResponseWriter, *http.Request)
}

func Configure(router *mux.Router) *mux.Router {
	for _, route := range loginRoutes {
		router.HandleFunc(route.URI, route.Action).Methods(route.Method)
	}

	return router
}
