package routes

import (
	"net/http"
	"tugasakhircoffe/TaCoffe/api/middlewares"

	"github.com/gorilla/mux"
)

//Route struct
type Route struct {
	URI          string
	Method       string
	Handler      func(w http.ResponseWriter, r *http.Request)
	AuthRequired bool
}

//Load the routes
func Load() []Route {
	routes := userRoutes
	routes = append(routes, authRoutes...)
	return routes
}

// SetupRoutesWithMiddlewares config routes with middlewares
func SetupRoutesWithMiddlewares(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		if route.AuthRequired {
			r.HandleFunc(route.URI,
				middlewares.SetMiddlewareLogger(
					middlewares.SetMiddlewareJSON(
						middlewares.SetMiddlewareAuthentication(route.Handler),
					),
				),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI,
				middlewares.SetMiddlewareLogger(
					middlewares.SetMiddlewareJSON(route.Handler),
				),
			).Methods(route.Method)
		}

	}
	return r
}
