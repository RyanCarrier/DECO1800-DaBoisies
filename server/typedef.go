package server

import (
	"strconv"
	"strings"
)

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
	List  []List `json:"list,omitempty"`
}

//List is the list within the list for json parseing
type List struct {
	Title    string     `json:"title"`
	ListItem []ListItem `json:"listItem,omitempty"`
}

//ListItem is an item in the list when returned
type ListItem struct {
	People []People `json:"people,omitempty"`
}

//People is the json struct containing the people id
type People struct {
	ID string `json:"id,omitempty"`
}

//PeopleIDs returns a list of the people id's in the TopResponse
func (tr TopResponse) PeopleIDs() []int {
	final := []int{}
	for _, list := range tr.Response.Zone[0].Records.List {
		for _, listItem := range list.ListItem {
			for _, people := range listItem.People {
				id, _ := strconv.Atoi(people.ID)
				final = append(final, id)
			}
		}
	}
	return final
}

//Clean converts the TopResponse to a CleanResponse
func (tr TopResponse) Clean() CleanResponse {
	cr := CleanResponse{
		Query: strings.TrimSpace(strings.Split(tr.Response.Query, "date:")[0]),
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
	Query string      `json:"query"`
	Year  int         `json:"year,omitempty"`
	Zones []CleanZone `json:"zones"`
}

//CleanZone is a clean version of the zone response from trove
type CleanZone struct {
	Name  string `json:"name"`
	Total int    `json:"total"`
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
	Query string `json:"query"`
	Year  int    `json:"year,omitempty"`
	Total int    `json:"total"`
}

//CleanPeopleReturn is the json struct used to return the people from forbes list
type CleanPeopleReturn struct {
	People []string `json:"people,omitempty"`
}
