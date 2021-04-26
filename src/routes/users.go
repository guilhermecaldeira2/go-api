package routes

import (
	"backend/src/controllers"
	"net/http"
)

var UsersRoutes = []Route{
	{
		URI:           "/users",
		Method:        http.MethodGet,
		Authenticated: false,
		Fn:            controllers.Show,
	},
	{
		URI:           "/users",
		Method:        http.MethodPost,
		Authenticated: false,
		Fn:            controllers.Create,
	},
	{
		URI:           "/users/{id}",
		Method:        http.MethodGet,
		Authenticated: false,
		Fn:            controllers.Find,
	},
	{
		URI:           "/users/{id}",
		Method:        http.MethodPut,
		Authenticated: false,
		Fn:            controllers.Update,
	},
	{
		URI:           "/users/{id}",
		Method:        http.MethodDelete,
		Authenticated: false,
		Fn:            controllers.Delete,
	},
}
