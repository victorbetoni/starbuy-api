package routes

import (
	"net/http"
	"starbuy/controllers"
)

var Purchase = []Route{

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
}
