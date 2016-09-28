package server

import "net/http"

func sendHelp(w http.ResponseWriter) {
	w.Write([]byte(help()))
}

func help() string {
	return "ur bad\noptions;\n/api/people/:peopleid"
}
