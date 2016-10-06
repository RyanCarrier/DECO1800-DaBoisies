package server

import (
	"strconv"
	"strings"
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

func troveSearchURLBuilder(search string, year int) string {
	if year == 0 {
		return troveURLBuilder("all", search)
	}
	yearS := strconv.Itoa(year)
	return troveURLBuilder("all", search) + strings.Replace(" date:["+
		yearS+" TO "+yearS+"]", " ", "%20", -1)
}
