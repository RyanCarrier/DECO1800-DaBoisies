package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
)

func getGoogleImageObject(query string) (string, error) {
	url := "https://www.googleapis.com/customsearch/v1?q=" + strings.Replace(strings.TrimSpace(query), " ", "%20", -1) + "&cx=" + googleCX + "&imgType=face&searchType=image&num=1&key=" + googleAPIKey
	log.Info("getting;", url)
	resp, err := http.Get(url)
	if err != nil {
		return errorImage, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errorImage, err
	}
	var sr SearchResponse
	err = json.Unmarshal(body, &sr)
	if err != nil {
		return errorImage, err
	}
	return sr.Clean()
}
