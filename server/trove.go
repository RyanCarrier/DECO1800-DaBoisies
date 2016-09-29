package server

import (
	"errors"
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
)

var zones = []string{"map", "collection", "list", "people", "book", "article", "music", "picture", "newspaper"}

const maxAttemptsList = 2

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
	handleListAttempt(w, r, 0)
}

func handleListAttempt(w http.ResponseWriter, r *http.Request, attempt int) {

}

func getList(attempt int) ([]byte, error) {
	if attempt > maxAttemptsList {
		return []byte{}, errors.New("Max list get attempts reached")
	}
	listURL := troveURLBuilder("list", "top") + "&include=listItems"
	body, err := http.Get(listURL)
	if err != nil {
		log.Error("error getting trove list, trying again", err)
		return getList(attempt + 1)
	}
	return body, nil
}
