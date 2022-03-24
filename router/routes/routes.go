package routes

import (
	"net/http"
	"starbuy/controllers"
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

var routes = []Route{
	{
		URI:         "/login",
		Method:      http.MethodPost,
		RequireAuth: false,
		Action:      controllers.Auth,
	},
	{
		URI:         "/register",
		Method:      http.MethodPost,
		RequireAuth: false,
		Action:      controllers.Register,
	},
	{
		URI:         "/user/{username}",
		Method:      http.MethodGet,
		RequireAuth: false,
		Action:      controllers.QueryUser,
	},
	{
		URI:         "/item/{id}",
		Method:      http.MethodGet,
		RequireAuth: false,
		Action:      controllers.GetItem,
	},
	{
		URI:         "/item/category/{id}",
		Method:      http.MethodGet,
		RequireAuth: false,
		Action:      controllers.GetCategory,
	},
	{
		URI:         "/items",
		Method:      http.MethodGet,
		RequireAuth: false,
		Action:      controllers.GetAllItems,
	},
	{
		URI:         "/item",
		Method:      http.MethodPost,
		RequireAuth: true,
		Action:      controllers.PostItem,
	},
	{
		URI:         "/cart",
		Method:      http.MethodGet,
		RequireAuth: true,
		Action:      controllers.QueryCart,
	},
	{
		URI:         "/cart",
		Method:      http.MethodPost,
		RequireAuth: true,
		Action:      controllers.PostCart,
	},
	{
		URI:         "/{item}/review",
		Method:      http.MethodPost,
		RequireAuth: true,
		Action:      controllers.PostReview,
	},
	{
		URI:         "/purchase",
		Method:      http.MethodPost,
		RequireAuth: true,
		Action:      controllers.PostPurchase,
	},
	{
		URI:         "/purchase/{id}",
		Method:      http.MethodPost,
		RequireAuth: true,
		Action:      controllers.PostPurchase,
	},
	{
		URI:         "/purchases",
		Method:      http.MethodGet,
		RequireAuth: true,
		Action:      controllers.GetPurchases,
	},
	{
		URI:         "/purchase/{id}",
		Method:      http.MethodGet,
		RequireAuth: true,
		Action:      controllers.GetPurchase,
	},
	{
		URI:         "/{user}/reviews",
		Method:      http.MethodGet,
		RequireAuth: false,
		Action:      controllers.GetReviews,
	},
	{
		URI:         "/review/{id}",
		Method:      http.MethodGet,
		RequireAuth: false,
		Action:      controllers.GetReview,
	},
}

func Configure(router *mux.Router) *mux.Router {
	for _, route := range routes {

		if route.RequireAuth {
			router.HandleFunc(route.URI, middleware.Authorize(route.Action)).Methods(route.Method)
		} else {
			router.HandleFunc(route.URI, route.Action).Methods(route.Method)
		}
	}

	return router
}
