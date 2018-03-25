package messages

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mfoster1989/WishBook_Server/logger"
)

var controller = &Controller{Repository: Repository{}}

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		controller.Index,
	},
	Route{
		"AddMessage",
		"POST",
		"/",
		controller.AddMessage,
	},
	Route{
		"UpdateMessage",
		"PUT",
		"/",
		controller.UpdateMessage,
	},
	Route{
		"DeleteMessage",
		"DELETE",
		"/{id}",
		controller.DeleteMessage,
	},
}

// NewRouter creates new MongoDB router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}