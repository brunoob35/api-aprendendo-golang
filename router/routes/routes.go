package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Routes struct {
	URI    string
	Method string
	Func   func(http.ResponseWriter, *http.Request)
}

func Register(r *mux.Router) {
	var routes []Routes
	routes = append(routes, booksRoutes...)
	routes = append(routes, usersRoutes...)

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}
}
