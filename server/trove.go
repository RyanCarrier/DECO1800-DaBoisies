package server

import "strings"

var zones = []string{"map", "collection", "list", "people", "book", "article", "music", "picture", "newspaper"}

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

//TopResponse is the top level response returned from trove
type TopResponse struct {
	response []Response
}

//Response is the main response from trove
type Response struct {
	query string
	zone  []Zone
}

//Zone is the response split by the search in each zone, there will only be one
// of these if only searching one zone.
type Zone struct {
	name    string
	records Record
}

//Record is the records of that zone, only having the total amount in there.
type Record struct {
	total int
}
