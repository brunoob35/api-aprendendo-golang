package routes

import (
	"Api-Aula1/controller"
	"net/http"
)

var loginRoutes = []Routes{
	{
		URI:    "/login",
		Method: http.MethodPost,
		Func:   controller.Login,
		Auth:   false,
	},
}
