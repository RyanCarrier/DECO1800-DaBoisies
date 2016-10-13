package server

import "errors"

const googleAPIKey = "AIzaSyDLwaqloOdBw5W9purSCN6gSoByqEuizyI"
const googleCX = "013041446805835708369%3Adgn8rbiqu_8"

const errorImage = "http://www.freeiconspng.com/uploads/error-icon-28.png"

//GET https://www.googleapis.com/customsearch/v1?q=Beyonce%2520Knowles&cx=013041446805835708369%3Adgn8rbiqu_8&key={YOUR_API_KEY}

//SearchResponse is the top level response returned from trove
type SearchResponse struct {
	SearchItem []SearchItem `json:"items"`
}

//SearchItem is the item returned fromt he google search
type SearchItem struct {
	//Pagemap Pagemap `json:"pagemap"`
	Link string `json:"link"`
}

/*
//Pagemap is a mapping of the items page
type Pagemap struct {
	ImageObject []ImageObject `json:"imageobject"`
}

//ImageObject is the object for the images returned
type ImageObject struct {
	Thumbnail  string `json:"thumbnail"`
	ContentURL string `json:"contenturl"`
}*/

//Clean cleans the search ready for returning
func (sr SearchResponse) Clean() (string, error) {
	if len(sr.SearchItem) == 0 {
		return errorImage, errors.New("searchresponse item len is 0")
	}
	if sr.SearchItem[0].Link == "" {
		return sr.SearchItem[0].Link, errors.New("Thumbnail empty")
	}
	return sr.SearchItem[0].Link, nil
}

//FinalPerson is a cleaner version of the search json for returning
type FinalPerson struct {
	Image string `json:"image,omitempty"`
	Name  string `json:"query"`
	Blurb string `json:"blurb,omitempty"`
}
