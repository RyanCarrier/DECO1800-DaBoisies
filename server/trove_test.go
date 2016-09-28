package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestParseSearchAll(t *testing.T) {
	body, err := ioutil.ReadFile("testsearch.json")
	if err != nil {
		t.Error("Error reading from test file", err)
	}
	var got TopResponse
	err = json.Unmarshal(body, &got)
	if err != nil {
		t.Error("Error unmarshaling", err)
	}
	want := TopResponse{
		response: []Response{
			Response{
				query: "test",
				zone: []Zone{
					Zone{
						name: "test",
						records: Record{
							total: 0,
						},
					},
				},
			},
		},
	}

	gots, err := json.Marshal(got)
	if err != nil {
		t.Error("Error marshaling got", err)
	}
	wants, err := json.Marshal(want)
	if err != nil {
		t.Error("Error marshaling want", err)
	}
	fmt.Println(string(gots))
	fmt.Println(string(wants))
	if strings.Compare(string(gots), string(wants)) != 0 {
		t.Error("GOT\n", gots, "WANT\n", wants)
	}
	fmt.Println(string(gots))
}
