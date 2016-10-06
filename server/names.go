package server

import (
	"errors"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

const attemptMax = 2
const attemptMax503 = 30

//Names is a struct to remember the names already found
type Names struct {
	sync.Mutex
	data map[int]string
}

var names = newNames()

func newNames() *Names {
	return &Names{
		data: make(map[int]string),
	}
}

//Get synchronously gets from the map
func (n *Names) Get(i int) (string, bool) {
	n.Lock()
	defer n.Unlock()
	got, ok := n.data[i]
	return got, ok
}

//Put synchronously puts into the map
func (n *Names) Put(i int, s string) {
	n.Lock()
	defer n.Unlock()
	n.data[i] = s
}

func handleNames(w http.ResponseWriter, r *http.Request) {
	logReq(r, r.URL.Path)
	vars := mux.Vars(r)
	search, ok := vars["peopleid"]
	if !ok {
		sendHelp(w)
		return
	}
	if strings.Contains(search, ",") {
		w.Write([]byte(getNames(search)))
	} else if i, err := strconv.Atoi(search); err == nil {
		w.Write([]byte(getName(i, 0, 0)))
	} else {
		sendHelp(w)
	}
	log.Info("Response sent")
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

func getNamesI(IDs []int) []string {
	final := make([]string, len(IDs))
	wg := &sync.WaitGroup{}
	log.Info("wait group adding ", len(IDs))
	wg.Add(len(IDs))
	for i, id := range IDs {
		go getMultiNameI(wg, &final[i], id)
	}
	wg.Wait()
	return final
}

func getMultiNameI(wg *sync.WaitGroup, s *string, id int) {
	*s = getName(id, 0, 0)
	wg.Done()
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
	if name, ok := names.Get(ID); ok {
		return name
	}
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
		time.Sleep(time.Millisecond * 10 * time.Duration(rand.Intn(25)))
		return getName(ID, attempt, attempt503+1)
	}
	name := decodeXML(ID, attempt, resp.Body)
	names.Put(ID, name)
	return name
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

func tryAgain(ID int, attempt int, err error) string {
	attempt++
	if attempt < attemptMax {
		return getName(ID, attempt, 0)
	}
	return err.Error()
}
