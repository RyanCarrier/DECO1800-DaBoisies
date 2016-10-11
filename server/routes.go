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
		Route{"/api/weight/{search}/", handleWeight},
		Route{"/api/weight/{search}/year/{year}", handleWeightYear},
		Route{"/api/people/{peopleid}", handleNames},
		Route{"/api/getlist/", handleList},
		Route{"/api/save/", handleSave},
		Route{"/api/load/", handleLoad},
		Route{"/api/", handleHelp},
	}
}
