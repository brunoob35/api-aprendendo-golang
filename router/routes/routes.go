package routes

import (
	"Api-Aula1/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Routes struct {
	URI    string
	Method string
	Func   func(http.ResponseWriter, *http.Request)
	Auth   bool
}

func Register(r *mux.Router) {
	var routes []Routes
	routes = append(routes, booksRoutes...)
	routes = append(routes, usersRoutes...)
	routes = append(routes, loginRoutes...)

	for _, route := range routes {
		if route.Auth {
			r.HandleFunc(route.URI,
				middlewares.Logger(
					middlewares.Autenticate(
						route.Func))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Func)).Methods(route.Method)
		}

	}
}
