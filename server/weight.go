package server

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

type weighting struct {
	sync.Mutex
	data map[string][]CleanResponse
}

var weightings = newWeighting()

func newWeighting() *weighting {
	return &weighting{
		data: make(map[string][]CleanResponse),
	}
}

func (w *weighting) convert(i int) (int, error) {
	//1=1980-2020
	if i == 0 {
		return 0, nil
	}
	if i < 1980 {
		return 0, errors.New("year " + strconv.Itoa(i) + " too low")
	}
	if i > 2020 {
		return 0, errors.New("year " + strconv.Itoa(i) + " too high")
	}
	return i - 1979, nil
}

func (w *weighting) get(s string, i int) (CleanResponse, error) {
	w.Lock()
	defer w.Unlock()
	i, err := w.convert(i)
	if err != nil {
		log.Warn("year wrong", err)
		return CleanResponse{}, err
	}
	got, ok := w.data[s]
	if !ok {
		return CleanResponse{}, errors.New(s + " not found")
	}
	if len(got[i].Zones) == 0 {
		err = errors.New("not set yet")
	}
	return got[i], err
}

func (w *weighting) put(s string, i int, cr CleanResponse) error {
	w.Lock()
	defer w.Unlock()
	i, err := w.convert(i)
	if err != nil {
		log.Warn("year wrong", err)
		return err
	}
	if a, ok := w.data[s]; ok {
		a[i] = cr
		w.data[s] = a
	} else {
		a := make([]CleanResponse, 41)
		a[i] = cr
		w.data[s] = a
	}
	return nil
}

func handleWeight(w http.ResponseWriter, r *http.Request) {
	handleWeightYear(w, r)
}

func handleWeightYear(w http.ResponseWriter, r *http.Request) {
	logReq(r, r.URL.Path)
	attemptWeightYear(w, r, 0)
}

func getWeightYear(year int, query string, attempt int) (CleanResponse, error) {
	var cr CleanResponse
	if attempt > attemptMax {
		return cr, errors.New("Too many attempts")
	}

	if cr, err := weightings.get(query, year); err != nil {
		if strings.Contains(err.Error(), "year") {
			return cr, err
		}
		url := troveSearchURLBuilder(query, year)
		log.Info("Accessing url; ", url)
		response, err := http.Get(url)
		if err != nil {
			log.Warn("Error getting weighting", err)
			return getWeightYear(year, query, attempt+1)
		}
		var gotr TopResponse
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Warn("Error reading response body", err)
			return getWeightYear(year, query, attempt+1)
		}
		err = json.Unmarshal(body, &gotr)
		if err != nil {
			log.Warn("Error unmarshalling response body", err)
			return getWeightYear(year, query, attempt+1)
		}
		cr = gotr.Clean()
		if year != 0 {
			cr.Year = year
		}
		//log.Info("Putting", query, year, cr)
		weightings.put(query, year, cr)
		return cr, nil
	}
	return cr, nil
}

func attemptWeightYear(w http.ResponseWriter, r *http.Request, attempt int) {
	if attempt > attemptMax {
		w.Write([]byte("Failed to get weight... check server logs..."))
		return
	}
	vars := mux.Vars(r)
	search, ok := vars["search"]
	if !ok {
		sendHelp(w)
		return
	}
	year, ok := vars["year"]
	var err error
	yeari := 0
	if ok && year != "" {
		yeari, err = strconv.Atoi(year)
		if err != nil {
			log.Warn("Error converting year", err)
			attemptWeightYear(w, r, attempt+1)
			return
		}
	}
	cr, err := getWeightYear(yeari, search, attempt)
	if err != nil {
		w.Write([]byte("Failed to get weight... check server logs..."))
		return
	}
	body, err := json.Marshal(cr)
	if err != nil {
		log.Warn("Error unmarshalling for responding", err)
		attemptWeightYear(w, r, attempt+1)
		return
	}
	w.Write(body)
}
