package routes

import (
	"net/http"
	controller "tugasakhircoffe/TaCoffe/api/controllers"
)

var authRoutes = []Route{
	{
		URI:          "/api/v1/login",
		Method:       http.MethodPost,
		Handler:      controller.Login,
		AuthRequired: false,
	},
}
