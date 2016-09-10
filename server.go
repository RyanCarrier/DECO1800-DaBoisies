package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
)

const attemptMax = 2
const attemptMax503 = 10

const port = 80

type fileHandler struct {
	root http.FileSystem
}
type apiHandler struct{}

func init() {
	if len(os.Args) > 1 {
		if os.Args[1][0] == '-' {
			cliHelp()
			os.Exit(0)
		}
		filename := os.Args[1]
		if filename[len(filename)-4:] != ".log" {
			filename += ".log"
		}
		f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
		if err != nil {
			log.Error("Couldn't open file")
		} else {
			log.SetOutput(f)
			fmt.Println("Listening on :" + strconv.Itoa(port))
			fmt.Println("Logging to", filename)
		}
	}
}

func main() {
	initHelp()
	log.Error(http.ListenAndServe(":"+strconv.Itoa(port), setupHandlers()))
}

func setupHandlers() *http.ServeMux {
	h := http.NewServeMux()
	fh := http.Handler(&fileHandler{http.Dir("html")})
	if len(os.Args) > 1 {
		h.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
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
		})
	}
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
			return
		}
		if strings.Contains(split[2], ",") {
			w.Write([]byte(getNames(split[2])))
		} else if i, err := strconv.Atoi(split[2]); err == nil {
			w.Write([]byte(getName(i, 0, 0)))
		} else {
			sendHelp(w)
		}
		log.Info("Response sent")
	})

	h.Handle("/", fh)
	return h
}

func getNames(IDs string) string {
	split := strings.Split(IDs, ",")
	final := make([]string, len(split))
	wg := &sync.WaitGroup{}
	log.Info("wait group adding ", len(split))
	log.Info(split)
	wg.Add(len(split))
	for i, id := range split {
		go getMultiName(wg, &final[i], id)
	}
	wg.Wait()
	finalCollective := ""
	for _, jd := range final {
		finalCollective += jd + ","
	}
	return finalCollective[:len(finalCollective)-1]
}

func getMultiName(wg *sync.WaitGroup, s *string, id string) {
	eyedee, err := strconv.Atoi(id)
	if err == nil {
		*s = getName(eyedee, 0, 0)
	} else {
		*s = "Don't be retarded"
	}
	wg.Done()
}

func getName(ID int, attempt int, attempt503 int) string {
	if attempt503 >= attemptMax503 {
		log.Info("too many 503 responses", ID)
		return "Too many 503 responses"
	}
	URL := "http://www.nla.gov.au/apps/srw/opensearch/peopleaustralia?q=" + strconv.Itoa(ID)
	resp, err := http.Get(URL)
	if err != nil {
		//We don't give a fuck about http status codes
		return tryAgain(ID, attempt, err)
	}
	if resp.StatusCode == 503 {
		log.Error("503 try again;", ID)
		time.Sleep(time.Millisecond * 50)
		return getName(ID, attempt, attempt503+1)
	}
	return decodeXML(ID, attempt, resp.Body)
}
func decodeXML(ID int, attempt int, r io.Reader) string {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		//yo fuck using errors
		log.Error(err)
		return tryAgain(ID, attempt, errors.New("ERROR READING XML"))
	}
	//HAHA DON"T EVEN PARSE IT
	split := strings.Split(string(body), "title")
	if len(split) < 7 {
		log.Error(split)
		return tryAgain(ID, attempt, errors.New("ERROR SPLITTING XML BY title"))
	}
	if len(split[5]) < 4 {
		log.Error(split[5])
		return tryAgain(ID, attempt, errors.New("ERROR SPLITTING XML 5th field too small"))
	}
	s := split[5][1 : len(split[5])-2]
	if strings.Contains(s, "(") && strings.Contains(s, ")") {
		s = strings.TrimSpace(strings.Split(s, "(")[0] + strings.Split(s, ")")[1])
	}

	split = strings.Split(s, ", ")

	if len(split) > 1 {
		return strings.TrimSpace(split[1] + " " + split[0])
	}
	return s
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
	if net.ParseIP(r.RemoteAddr) != nil {
		log.Info(upath, " from ", net.ParseIP(r.RemoteAddr))
	} else {
		log.Info(upath, " from ", r.Header.Get("X-Forwarded-For"))
	}
}

func tryAgain(ID int, attempt int, err error) string {
	attempt++
	if attempt < attemptMax {
		return getName(ID, attempt, 0)
	}
	return err.Error()
}

func initHelp() {
	log.Info("Listening on http://localhost:" + strconv.Itoa(port))
	log.Info(`To get people use the api;
http://localhost:` + strconv.Itoa(port) + `/api/people/:peopleid
Try using some of the id's from forbes list;
946924
1456345
1163655
1479624
1503477
		`)
}

func cliHelp() {
	fmt.Println("./server [logfile|--help]")
}
