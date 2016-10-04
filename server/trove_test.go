package server

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
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
	tr := getTopResponseList(t)
	expectedPeople := []int{946924, 1456345, 1163655, 1479624, 1503477, 1443296,
		1102410, 1454049, 1442529, 1046959, 1680377, 814325, 1000809, 1457148,
		1036861, 1445254, 1070798, 1052490, 1454383, 1453947, 1459965, 833430,
		867994, 783697, 914919, 1194709, 1446268, 1015739, 1639257, 943439,
		1107079, 1456818, 1446252, 1460045, 1455470, 948434, 611715, 891834,
		1018071, 1021436, 1457697, 688060, 1443355, 1457785, 1234441, 1459386,
		1458079, 1454092, 1068266, 1175064}
	gotPeople := tr.PeopleIDs()
	if len(gotPeople) != len(expectedPeople) {
		t.Error("got people not same size as expected people\nGOT:" + strconv.Itoa(len(gotPeople)) + "\nWANT:" + strconv.Itoa(len(expectedPeople)))
	}
	if !matches(gotPeople, expectedPeople) {
		t.Error("gotPeople != expecedPeople")
	}
}
func matches(a, b []int) bool {
	for _, c := range a {
		if !contains(b, c) {
			return false
		}
	}
	return true
}
func contains(a []int, b int) bool {
	for _, c := range a {
		if b == c {
			return true
		}
	}
	return false
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
