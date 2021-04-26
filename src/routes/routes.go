package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route type
type Route struct {
	URI           string
	Method        string
	Fn            func(http.ResponseWriter, *http.Request)
	Authenticated bool
}

// Subscribe routes into router
func Subscribe(r *mux.Router) *mux.Router {
	routes := UsersRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Fn).Methods(route.Method)
	}

	return r
}
