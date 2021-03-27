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
}
