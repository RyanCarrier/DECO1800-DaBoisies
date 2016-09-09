package main

import (
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
)

type fileHandler struct {
	root http.FileSystem
}
type apiHandler struct{}

func init() {
	if !isTTY() {
		log.SetFormatter(&StdFormatter{})
	}
}
func main() {

	log.Info("Listening on http://localhost:8080/")
	log.Info(`To get people use the api;
http://localhost:8080/api/people/:peopleid
Try using some of the id's from forbes list;
946924
1456345
1163655
1479624
1503477
		`)
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

func isTTY() bool {
	return terminal.IsTerminal(int(os.Stdout.Fd()))
}

//StdFormatter for non tty
type StdFormatter struct {
	// TimestampFormat sets the format used for marshaling timestamps.
	TimestampFormat string
}

//Format the entry to stdstring for non tty
func (f *StdFormatter) Format(entry *log.Entry) ([]byte, error) {
	data := make(log.Fields, len(entry.Data)+3)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			// Otherwise errors are ignored by `encoding/json`
			// https://github.com/Sirupsen/logrus/issues/137
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = log.DefaultTimestampFormat
	}

	return []byte(entry.Time.Format(timestampFormat) + " " + entry.Message), nil
}
