package routes

import (
	"Api-Aula1/controller"
	"net/http"
)

var usersRoutes = []Routes{
	{
		URI:    "/users",
		Method: http.MethodPost,
		Func:   controller.CreateUser,
	},
	{
		URI:    "/users",
		Method: http.MethodGet,
		Func:   controller.FetchUser,
	},
	{
		URI:    "/users{userID}",
		Method: http.MethodGet,
		Func:   controller.UpdateUser,
	},
	{
		URI:    "/users{userID}",
		Method: http.MethodDelete,
		Func:   controller.DeleteUser,
	},
}
