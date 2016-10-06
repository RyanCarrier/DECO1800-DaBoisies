package server

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

const htmlDir = "html"

//SetupHandlers returns a handler for all the routes
func SetupHandlers() *mux.Router {
	r := mux.NewRouter()
	//get and set all routes
	for _, route := range GetRoutes() {
		r.HandleFunc(route.Path, route.Handler)
	}
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(htmlDir)))
	//Set the default
	return r
}

func handleHelp(w http.ResponseWriter, r *http.Request) {
	final := ""
	for _, route := range GetRoutes() {
		final += "<a href=\"" + strings.Replace(strings.Replace(route.Path, "{", "",
			-1), "}", "Example", -1) + "\">" + route.Path + "</a><br>"
	}
	w.Write([]byte(final))
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}
	logReq(r, upath)
	http.ServeFile(w, r, htmlDir+upath)
}
