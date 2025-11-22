package routes

import (
	"Api-Aula1/controller"
	"net/http"
)

var booksRoutes = []Routes{
	{
		URI:    "/books",
		Method: http.MethodGet,
		Func:   controller.HandleSearch,
	},
}
