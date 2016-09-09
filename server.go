package main

import (
	"net"
	"net/http"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
)

type fileHandler struct {
	root http.FileSystem
}
type apiHandler struct{}

func main() {

	log.Info("Listening on http://localhost:8080/")
	http.ListenAndServe(":8080", setupHandlers())
}

func setupHandlers() *http.ServeMux {
	h := http.NewServeMux()
	fh := http.Handler(&fileHandler{http.Dir("html")})

	//handle files
	h.HandleFunc("/api/people/", func(w http.ResponseWriter, r *http.Request) {
		URL := strings.TrimSpace(r.URL.EscapedPath())[1:]
		if URL[len(URL)-1] == '/' {
			URL = URL[:len(URL)-2]
		}
		split := strings.Split(URL, "/")
		if len(split) != 3 {
			sendHelp(w)
		} else if i, err := strconv.Atoi(split[2]); err == nil {
			w.Write([]byte(getName(i)))
		} else {
			sendHelp(w)
		}
	})

	h.Handle("/", fh)
	return h
}

func getName(ID int) string {
	return strconv.Itoa(ID)
}

func sendHelp(w http.ResponseWriter) {
	w.Write([]byte(help()))
}
func help() string {
	return "ur bad\noptions;\n/api/people/:peopleid"
}

func (f *fileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}
	if net.ParseIP(r.RemoteAddr) == nil {
		log.Info(upath, " from ", net.ParseIP(r.RemoteAddr))
	} else {
		log.Info(upath, " from ", r.Header.Get("X-Forwarded-For"))
	}
	http.ServeFile(w, r, "html"+upath)
}
