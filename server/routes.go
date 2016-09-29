package server

import "net/http"

//Routes just wraps the Route in a slice for easier readibility
type Routes []Route

//Route specifies a single route to define
type Route struct {
	Path    string
	Handler func(http.ResponseWriter, *http.Request)
}

//GetRoutes gets all the set routes
func GetRoutes() Routes {
	return Routes{
		Route{"/log", handleLog},
		Route{"/api/people/", handleNames},
		Route{"/api/getlist/", handleList},
	}
}
