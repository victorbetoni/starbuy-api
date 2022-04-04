package routes

import (
	"starbuy/middleware"

	"github.com/gin-gonic/gin"
)

// Route - Representação de todas as rotas da API
type Route struct {
	RequireAuth bool
	URI         string
	Action      gin.HandlerFunc
	Assign      func(*gin.Engine, gin.HandlerFunc, string)
}

func Configure(router *gin.Engine) *gin.Engine {
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
				route.Assign(router, middleware.Authorize(route.Action), route.URI)
			} else {
				route.Assign(router, route.Action, route.URI)
			}
		}
	}
	return router
}
