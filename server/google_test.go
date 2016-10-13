package server

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestParseBeyonceSearch(t *testing.T) {
	want := "https://fuzfeed.com/wp-content/uploads/2014/09/Beyonc%C3%A9-newest-photos.jpg"
	got, err := getSearchResponse(t).Clean()

	if err != nil || got != want {
		t.Error("GOT Thumb:\n" + got +
			"\nWANT Thumb:\n" + want)
	}
}

func getSearchResponse(t *testing.T) SearchResponse {
	body, err := ioutil.ReadFile("beyonceSearch.json")
	if err != nil {
		t.Error("Error reading from test file", err)
	}
	var gotr SearchResponse
	err = json.Unmarshal(body, &gotr)
	if err != nil {
		t.Error("Error unmarshaling", err)
	}
	return gotr
}
