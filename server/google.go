package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
)

func getGoogleImageObject(query string) (ImageObject, error) {
	url := "https://www.googleapis.com/customsearch/v1?q=" + strings.Replace(strings.TrimSpace(query), " ", "%20", -1) + "&cx=" + googleCX + "&key=" + googleAPIKey
	log.Info("getting;", url)
	resp, err := http.Get(url)
	if err != nil {
		return ImageObject{Thumbnail: errorImage, ContentURL: errorImage}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ImageObject{Thumbnail: errorImage, ContentURL: errorImage}, err
	}
	var sr SearchResponse
	err = json.Unmarshal(body, &sr)
	if err != nil {
		return ImageObject{Thumbnail: errorImage, ContentURL: errorImage}, err
	}
	return sr.Clean()
}
