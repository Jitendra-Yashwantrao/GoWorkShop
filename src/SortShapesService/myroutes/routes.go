package myroutes

import (
	"SortShapesService/handlers"
	"SortShapesService/mylogger"
	"net/http"

	"github.com/gorilla/mux"
)

//Route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes
type Routes []Route

//RoutesList
var RoutesList = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},

	Route{
		"SortShapes",
		"POST",
		"/SortShapes",
		handlers.SortShapes,
	},
}

//NewRouter
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range RoutesList {

		var handler http.Handler

		handler = route.HandlerFunc
		handler = mylogger.Logger(handler, route.Name)

		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
