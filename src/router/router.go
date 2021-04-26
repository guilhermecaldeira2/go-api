package router

import (
	"backend/src/routes"

	"github.com/gorilla/mux"
)

func Load() *mux.Router {
	r := mux.NewRouter()

	return routes.Subscribe(r)
}
