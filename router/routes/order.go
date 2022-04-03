package routes

import (
	"net/http"
	"starbuy/controllers"
)

var Order = []Route{

	{
		URI:         "/order",
		Method:      http.MethodPost,
		RequireAuth: true,
		Action:      controllers.PostPurchase,
	},
	{
		URI:         "/order/{id}",
		Method:      http.MethodPost,
		RequireAuth: true,
		Action:      controllers.PostPurchase,
	},
	{
		URI:         "/orders",
		Method:      http.MethodGet,
		RequireAuth: true,
		Action:      controllers.GetPurchases,
	},
	{
		URI:         "/order/{id}",
		Method:      http.MethodGet,
		RequireAuth: true,
		Action:      controllers.GetPurchase,
	},
}
