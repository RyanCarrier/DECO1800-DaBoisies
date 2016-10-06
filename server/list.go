package server

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"

	log "github.com/Sirupsen/logrus"
)

type ids struct {
	sync.Mutex
	data []int
}

var iDs = newids()

func newids() *ids {
	return &ids{
		data: []int{},
	}
}

func (i *ids) get() []int {
	i.Lock()
	defer i.Unlock()
	data := i.data
	return data
}

func (i *ids) put(data []int) {
	i.Lock()
	defer i.Unlock()
	i.data = data
}

func handleList(w http.ResponseWriter, r *http.Request) {
	logReq(r, r.URL.Path)
	if len(iDs.get()) == 0 {
		tr, err := getListStruct(0)
		if err != nil {
			body, err := ioutil.ReadFile("testlist.json")
			if err != nil {
				log.Error("can't read from file", err)
				w.Write([]byte("Couldn't get list soz about it nerd"))
				return
			}
			err = json.Unmarshal(body, &tr)
			if err != nil {
				log.Error("can't unmarshal from file", err)
				w.Write([]byte("Couldn't get list soz about it nerd2.0"))
				return
			}
		}

		//regardless tr is full now yum yum
		iDs.put(tr.PeopleIDs())
	}
	cpr := CleanPeopleReturn{People: getNamesI(iDs.get())}
	body, err := json.Marshal(cpr)
	if err != nil {
		log.Error("UGHHHHHHH", err)
		w.Write([]byte("Couldn't get list soz about it nerd"))
		return
	}
	w.Write(body)
}

func getListStruct(attempt int) (TopResponse, error) {
	if attempt > 0 {
		log.Warn("attempting to get list again, attempt: ", attempt)
	}
	var tr TopResponse
	if attempt > maxAttemptsList {
		return tr, errors.New("Max list get attempts reached")
	}
	body, err := getList(attempt)
	if err != nil {
		log.Error(err)
		return getListStruct(attempt + 1)
	}
	err = json.Unmarshal(body, &tr)
	//fmt.Println(string(body))
	if err != nil {
		log.Error(err)
		return getListStruct(attempt + 1)
	}
	return tr, nil
}

func getList(attempt int) ([]byte, error) {
	if attempt > maxAttemptsList {
		return []byte{}, errors.New("Max list get attempts reached")
	}
	listURL := troveURLBuilder("list", "top") + "&include=listItems&n=1"
	response, err := http.Get(listURL)
	if err != nil {
		log.Error("error getting trove list, trying again", err)
		return getList(attempt + 1)
	}
	if response.StatusCode/100 != 2 {
		err = errors.New("Response code non 200 " + strconv.Itoa(response.StatusCode))
		log.Error(err)
		return getList(attempt + 1)
	}
	log.Info("successful? get from trove list")
	return ioutil.ReadAll(response.Body)
}
