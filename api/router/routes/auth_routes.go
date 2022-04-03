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
	{
		URI:          "/api/v1/signup",
		Method:       http.MethodPost,
		Handler:      controller.Login,
		AuthRequired: false,
	},
	{
		URI:          "/api/v1/forgot_password",
		Method:       http.MethodPost,
		Handler:      controller.Login,
		AuthRequired: false,
	},
	{
		URI:          "/api/v1/auth_otp",
		Method:       http.MethodPost,
		Handler:      controller.Login,
		AuthRequired: false,
	},
}
