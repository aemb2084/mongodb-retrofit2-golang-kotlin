package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"mongo.api/controllers"
)

type Route struct {
	Nombre      string
	Metodo      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"Ping", "GET", "/ping", controllers.Ping},
	Route{"Dblist", "GET", "/dblist", controllers.Dblist},
	Route{"QueryByIdURL", "GET", "/queryById/{_id}", controllers.QueryByIdURL},
	Route{"QueryDocument", "GET", "/queryDocument", controllers.QueryDocument},
	Route{"QueryByIdJson", "POST", "/queryById", controllers.QueryByIdJson},
	Route{"CreateItemByStructure", "PUT", "/createItemByStructure", controllers.CreateItem},
	Route{"CreateItemByJson", "PUT", "/createItemByJson", controllers.CreateItemByJson},
	Route{"DeleteOneById", "DELETE", "/deleteOneById/{_id}", controllers.DeleteOneById},
}

func NewRoute() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.Metodo).
			Path(route.Pattern).
			Name(route.Nombre).
			Handler(route.HandlerFunc)
	}

	return router
}
