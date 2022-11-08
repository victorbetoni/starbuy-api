package routes

import (
	"net/http"
	"starbuy/middleware"
	"starbuy/util"

	"github.com/gin-gonic/gin"
)

type AssignFunction func(*gin.Engine, gin.HandlerFunc, string)

// Route - Representação de todas as rotas da API
type Route struct {
	RequireAuth bool
	URI         string
	Action      util.HandlerFunc
	Assign      AssignFunction
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
				Assign(route.Assign, middleware.CORS(middleware.Authorize(middleware.AbortOnError(route.Action))), route.URI, router)
			} else {
				Assign(route.Assign, middleware.CORS(middleware.AbortOnError(route.Action)), route.URI, router)
			}
		}
	}

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Pong!")
	})
	return router
}

func Assign(assign AssignFunction, handler util.HandlerFunc, uri string, router *gin.Engine) {
	assign(router, func(context *gin.Context) {
		handler(context)
	}, uri)
}
