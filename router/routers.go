package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jacobintern/MyselfCryptoRecord/controllers"
)

type Route struct {
	Method     string
	Path       string
	Handler    http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

var routes []Route

func init() {
	// users
	register(http.MethodGet, "/api/users/{id}", controllers.GetUser, nil)
	register(http.MethodPost, "/api/users/{id}", controllers.CreateUser, nil)
	register(http.MethodPut, "/api/users/{id}", controllers.UpdateUser, nil)
	register(http.MethodDelete, "/api/users/{id}", controllers.DeleteUser, nil)

	// login
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	for _, route := range routes {
		r.HandleFunc(route.Path, route.Handler).Methods(route.Method)
		//r.Methods(route.Method).Path(route.Path).Handler(route.Handler)

		if route.Middleware != nil {
			r.Use(route.Middleware)
		}
	}
	return r
}

func register(method, path string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, Route{method, path, handler, middleware})
}
