package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {

	UserRoutes(r)
	ProfileRoutes(r) // Add this code
	ProductRoutes(r) // Add this code
}
