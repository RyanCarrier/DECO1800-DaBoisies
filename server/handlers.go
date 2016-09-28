package server

import (
	"net/http"
	"os"
	"strings"
)

type fileHandler struct {
	root http.FileSystem
}
type apiHandler struct{}

//SetupHandlers returns a handler for all the routes
func SetupHandlers() *http.ServeMux {
	h := http.NewServeMux()
	fh := http.Handler(&fileHandler{http.Dir("html")})
	if len(os.Args) > 1 {
		h.HandleFunc("/log", handleLog)
	}
	//handle files
	h.HandleFunc("/api/people/", handleNames)

	h.Handle("/", fh)
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
