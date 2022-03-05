package routes

import "net/http"

// Route - Representação de todas as rotas da API
type Route struct {
	URI    string
	Method string
	Action func(http.ResponseWriter, *http.Request)
}
