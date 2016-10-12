package server

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestParseBeyonceSearch(t *testing.T) {
	want := ImageObject{ContentURL: "http://data.whicdn.com/images/35549397/large.jpg",
		Thumbnail: "http://data.whicdn.com/images/35549397/thumbnail.jpg",
	}
	got := getSearchResponse(t).Clean()
	if got.ContentURL != want.ContentURL || got.Thumbnail != got.Thumbnail {
		t.Error("GOT Thumb:\n" + got.Thumbnail +
			"\nWANT Thumb:\n" + want.Thumbnail +
			"\nGOT ContentURL:\n" + got.ContentURL +
			"\nWANT ContentURL:\n" + want.ContentURL)
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
