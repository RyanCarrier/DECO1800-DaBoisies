package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestParseSearchAll(t *testing.T) {
	got := getTopResponse(t).Clean()
	want := CleanResponse{
		Query: "bob ross",
		Zones: []CleanZone{
			CleanZone{
				Name:  "people",
				Total: 21,
			},
			CleanZone{
				Name:  "book",
				Total: 2770,
			},
			CleanZone{
				Name:  "picture",
				Total: 2172,
			},
			CleanZone{
				Name:  "music",
				Total: 3126,
			},
			CleanZone{
				Name:  "list",
				Total: 119,
			},
			CleanZone{
				Name:  "map",
				Total: 104,
			},
			CleanZone{
				Name:  "collection",
				Total: 567,
			},
			CleanZone{
				Name:  "article",
				Total: 47880,
			},
			CleanZone{
				Name:  "newspaper",
				Total: 146843,
			},
		},
	}
	marshalAndTest(t, got, want)

}

func TestReturnResult(t *testing.T) {
	got := getTopResponse(t).Clean().Return()
	want := CleanReturn{
		Query: "bob ross",
		//Year:blank
		Total: 146843 + 47880 + 567 + 104 + 119 + 3126 + 2172 + 2770 + 21,
	}
	marshalAndTest(t, got, want)
}

func TestListStructs(t *testing.T) {
	body, _ := json.Marshal(getTopResponseList(t))
	fmt.Println(string(body))
}
func TestReturnResultWithYear(t *testing.T) {
	gotr := getTopResponse(t).Clean()
	gotr.Year = 1234
	got := gotr.Return()
	want := CleanReturn{
		Query: "bob ross",
		Year:  1234,
		Total: 146843 + 47880 + 567 + 104 + 119 + 3126 + 2172 + 2770 + 21,
	}
	marshalAndTest(t, got, want)
}

func getTopResponse(t *testing.T) TopResponse {
	body, err := ioutil.ReadFile("testsearch.json")
	if err != nil {
		t.Error("Error reading from test file", err)
	}
	var gotr TopResponse
	err = json.Unmarshal(body, &gotr)
	if err != nil {
		t.Error("Error unmarshaling", err)
	}
	return gotr
}

func getTopResponseList(t *testing.T) TopResponse {
	body, err := ioutil.ReadFile("testlist.json")
	if err != nil {
		t.Error("Error reading from test file", err)
	}
	var gotr TopResponse
	err = json.Unmarshal(body, &gotr)
	if err != nil {
		t.Error("Error unmarshaling", err)
	}
	return gotr
}

func marshalAndTest(t *testing.T, got, want interface{}) {
	gotb, err := json.Marshal(got)
	if err != nil {
		t.Error("Error marshaling got", err)
	}
	wantb, err := json.Marshal(want)
	if err != nil {
		t.Error("Error marshaling want", err)
	}
	gots := string(gotb)
	wants := string(wantb)
	if strings.Compare(gots, wants) != 0 {
		t.Error("GOT\n", gots, "\nWANT\n", wants)
	}
}
