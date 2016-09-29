package server

import (
	"strconv"
	"strings"
)

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
	Response Response `json:"response,omitempty"`
}

//Response is the main response from trove
type Response struct {
	Query string `json:"query,omitempty"`
	Zone  []Zone `json:"zone,omitempty"`
}

//Zone is the response split by the search in each zone, there will only be one
// of these if only searching one zone.
type Zone struct {
	Name    string `json:"name,omitempty"`
	Records Record `json:"records,omitempty"`
}

//Record is the records of that zone, only having the total amount in there.
type Record struct {
	//note trove uses the total as a string not int
	Total string `json:"total,omitempty"`
}

//Clean converts the TopResponse to a CleanResponse
func (tr TopResponse) Clean() CleanResponse {
	cr := CleanResponse{
		Query: tr.Response.Query,
		Zones: make([]CleanZone, len(tr.Response.Zone)),
	}
	for i, z := range tr.Response.Zone {
		cr.Zones[i].Name = z.Name
		cr.Zones[i].Total, _ = strconv.Atoi(z.Records.Total)
	}
	return cr
}

//CleanResponse is a clean version of the response from trove
type CleanResponse struct {
	Query string      `json:"query,omitempty"`
	Year  int         `json:"year,omitempty"`
	Zones []CleanZone `json:"zones,omitempty"`
}

//CleanZone is a clean version of the zone response from trove
type CleanZone struct {
	Name  string `json:"name,omitempty"`
	Total int    `json:"total,omitempty"`
}

//Return converts CleanResponse into a barebones returnable json struct CleanReturn
func (cr CleanResponse) Return() CleanReturn {
	ret := CleanReturn{
		Query: cr.Query,
		Year:  cr.Year,
		Total: 0,
	}
	for _, z := range cr.Zones {
		ret.Total += z.Total
	}
	return ret
}

//CleanReturn is the json struct used to return the minimum amount to the client
type CleanReturn struct {
	Query string `json:"query,omitempty"`
	Year  int    `json:"year,omitempty"`
	Total int    `json:"total,omitempty"`
}
