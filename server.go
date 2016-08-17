package main

import (
	"log"
	"net/http"
	"strings"
	"net"
)

type fileHandler struct {
	root http.FileSystem
}

func main() {
	fh := http.Handler(&fileHandler{http.Dir("html")})
	http.Handle("/", fh)

	log.Println("Listening on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func (f *fileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}
	if net.ParseIP(r.RemoteAddr) == nil {
		log.Println(upath," from ", net.ParseIP(r.RemoteAddr))
	}else{
		log.Println(upath," from ", r.Header.Get("X-Forwarded-For"))
	}
	http.ServeFile(w, r, "html"+upath)
}
