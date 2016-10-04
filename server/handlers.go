package server

import (
	"net/http"
	"strings"
)

type fileHandler struct {
	root http.FileSystem
}
type apiHandler struct{}

const htmlDir = "html"

//SetupHandlers returns a handler for all the routes
func SetupHandlers() *http.ServeMux {
	h := http.NewServeMux()
	//get and set all routes
	for _, r := range GetRoutes() {
		h.HandleFunc(r.Path, r.Handler)
	}
	//Set the default
	h.Handle("/", http.Handler(&fileHandler{http.Dir(htmlDir)}))
	return h
}

func (f *fileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}
	logReq(r, upath)
	http.ServeFile(w, r, "html"+upath)
}

func handleHelp(w http.ResponseWriter, r *http.Request) {
	final := ""
	for _, route := range GetRoutes() {
		final += "<a href=\"" + route.Path + "\">" + route.Path + "</a><br>"
	}
	w.Write([]byte(final))
}
