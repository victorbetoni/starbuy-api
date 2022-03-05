package router

import "github.com/gorilla/mux"

// Build vai retornar um router com as rotas configuradas
func Build() *mux.Router {
	return mux.NewRouter()
}
