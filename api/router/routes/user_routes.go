package routes

import (
	"net/http"

	controller "tugasakhircoffe/TaCoffe/api/controllers"
)

var userRoutes = []Route{
	{
		URI:          "/api/v1/users",
		Method:       http.MethodGet,
		Handler:      controller.GetUsers,
		AuthRequired: true,
	},
	{
		URI:          "/api/v1/users",
		Method:       http.MethodPost,
		Handler:      controller.CreateUsers,
		AuthRequired: true,
	},
	{
		URI:          "/api/v1/users/{id}",
		Method:       http.MethodGet,
		Handler:      controller.GetUserDetails,
		AuthRequired: true,
	},
	{
		URI:          "/api/v1/users/{id}",
		Method:       http.MethodPut,
		Handler:      controller.UpdateUser,
		AuthRequired: true,
	},
	{
		URI:          "/api/v1/users/{id}",
		Method:       http.MethodDelete,
		Handler:      controller.DeleteUser,
		AuthRequired: true,
	},
}
