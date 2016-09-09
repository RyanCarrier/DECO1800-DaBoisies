package main

import (
	"io"
	"io/ioutil"
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
		logReq(r, r.URL.Path)
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
	URL := "http://www.nla.gov.au/apps/srw/opensearch/peopleaustralia?q=" + strconv.Itoa(ID)
	resp, err := http.Get(URL)
	if err != nil {
		//We don't give a fuck about http status codes
		return "ERROR GETTING NAME"
	}
	return decodeXML(resp.Body)
}
func decodeXML(r io.Reader) string {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		//yo fuck using errors
		log.Error(err)
		return "ERROR READING XML"
	}
	//HAHA DON"T EVEN PARSE IT
	split := strings.Split(string(body), "title")
	if len(split) < 7 {
		log.Error(split)
		return "ERROR SPLITTING XML"
	}
	split = strings.Split(split[5][1:len(split[5])-2], ", ")
	return split[1] + " " + split[0]
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
	logReq(r, upath)
	http.ServeFile(w, r, "html"+upath)
}

func logReq(r *http.Request, upath string) {
	if net.ParseIP(r.RemoteAddr) == nil {
		log.Info(upath, " from ", net.ParseIP(r.RemoteAddr))
	} else {
		log.Info(upath, " from ", r.Header.Get("X-Forwarded-For"))
	}
}
