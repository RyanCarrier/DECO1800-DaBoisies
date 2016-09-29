package server

import (
	"io/ioutil"
	"net"
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
)

func logReq(r *http.Request, upath string) {
	if net.ParseIP(r.RemoteAddr) != nil {
		log.Info(upath, " from ", net.ParseIP(r.RemoteAddr))
	} else {
		log.Info(upath, " from ", r.Header.Get("X-Forwarded-For"))
	}
}

func handleLog(w http.ResponseWriter, r *http.Request) {
	if len(os.Args) < 2 {
		w.Write([]byte("No logging file specified"))
		return
	}
	filename := os.Args[1]
	if filename[len(filename)-4:] != ".log" {
		filename += ".log"
	}
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		w.Write([]byte("Couldn't read log at " + filename))
	} else {
		w.Write(body)
	}
}
