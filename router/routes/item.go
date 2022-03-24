package routes

import (
	"net/http"
	"starbuy/controllers"
)

var Item = []Route{
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
}
