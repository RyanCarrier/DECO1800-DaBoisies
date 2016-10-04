package server

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
)

var zones = []string{"map", "collection", "list", "people", "book", "article", "music", "picture", "newspaper"}

const maxAttemptsList = 3

//Trove api key rcarrier's
var apikeys = []string{
	"j0porbqbr4efdh2c", //rcarrier
	"ulsmhsa32qhk0fhv", //robin
	"a79q82q1nosa67ck", //sam
	"8lkjcg45qi640t9s", //big dogs
	"grcr2nt2i61ourfj", //georgie
}

var keyIndex = 0

func getKey() int {
	keyIndex = (keyIndex + 1) % len(apikeys)
	return keyIndex
}

func troveURLBuilder(zone, search string) string {
	search = strings.Replace(search, " ", "%20", -1)
	return "http://api.trove.nla.gov.au/result?key=" + apikeys[getKey()] + "&encoding=json&zone=" +
		zone + "&q=" + search //+ "&callback=?"
}

func handleList(w http.ResponseWriter, r *http.Request) {
	logReq(r, r.URL.Path)
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
	cpr := CleanPeopleReturn{People: getNamesI(tr.PeopleIDs())}
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
