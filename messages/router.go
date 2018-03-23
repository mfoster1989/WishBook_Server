package messages

import (
	"net/http"
)

type Route Struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

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
 }
}

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