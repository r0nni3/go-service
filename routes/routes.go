package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/r0nn3/go-service/controllers"
)

type Route struct {
	Name        string
	Method      []string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method...).
			Path(route.Pattern).
			Name(route.Name).
			Handler(controllers.HandleRoute(route.HandlerFunc))
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		[]string{"GET"},
		"/",
		controllers.IndexHandler,
	},
	// Your Routes go here
	Route{
		"Presentations",
		[]string{"GET", "POST"},
		"/presentation",
		controllers.PresentationsHandler,
	},
	Route{
		"Presentation",
		[]string{"GET", "POST", "DELETE"},
		"/presentation/{id:[0-9]+}",
		controllers.PresentationHandler,
	},
}
