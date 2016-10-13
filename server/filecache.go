package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
)

//Worst file ever
const iDsFile = "idsFile.json"
const namesFile = "namesFile.json"
const peopleFile = "peopleFile.json"

//Load loads the cached stuff from file
func Load() {
	log.Info("Loading...")
	loadIDS()
	loadNames()
	loadPeople()
	time.Sleep(time.Second * 3)
	http.Get("http://localhost/api/getlist/")
	log.Info("Loaded.")
}

//DelayLoad delays the loading
func DelayLoad(i int) {
	time.Sleep(time.Second * time.Duration(i))
	Load()
}

//Save saves the cached stuff to file
func Save() {
	saveIDS()
	saveNames()
	savePeople()
}

func handleSave(w http.ResponseWriter, r *http.Request) {
	Save()
	w.Write([]byte("kk"))
}

func handleLoad(w http.ResponseWriter, r *http.Request) {
	Load()
	w.Write([]byte("kk"))
}

func saveIDS() {
	iDs.Lock()
	defer iDs.Unlock()
	body, err := json.Marshal(iDs)
	if err != nil {
		log.Warn("Error while marshalling iDsFile", err)
		return
	}
	if err := ioutil.WriteFile(iDsFile, body, 0666); err != nil {
		log.Warn("error writing to idsfile", err)
	}
}
func loadIDS() {
	iDs.Lock()
	defer iDs.Unlock()
	var got ids
	body, err := ioutil.ReadFile(iDsFile)
	if err != nil {
		log.Warn("Error while reading iDsFile", err)
		return
	}
	err = json.Unmarshal(body, &got)
	if err != nil {
		log.Warn("Error while unmarshaling iDsFile", err)
	}
	iDs.Data = got.Data
}

func saveNames() {
	names.Lock()
	defer names.Unlock()
	body, err := json.Marshal(names)
	if err != nil {
		log.Warn("Error while marshalling namesFile", err)
		return
	}
	if err := ioutil.WriteFile(namesFile, body, 0666); err != nil {
		log.Warn("error writing to namesFile", err)
	}
}
func loadNames() {
	names.Lock()
	defer names.Unlock()
	body, err := ioutil.ReadFile(namesFile)
	if err != nil {
		log.Warn("Error while reading namesFile", err)
		return
	}
	var got Names
	err = json.Unmarshal(body, &got)
	if err != nil {
		log.Warn("Error while unmarshaling namesFile", err)
	}
	names.Data = got.Data
}

func savePeople() {
	peoples.Lock()
	defer peoples.Unlock()
	body, err := json.Marshal(peoples)
	if err != nil {
		log.Warn("Error while marshalling peopleFile", err)
		return
	}
	if err := ioutil.WriteFile(peopleFile, body, 0666); err != nil {
		log.Warn("error writing to peopleFile", err)
	}
}
func loadPeople() {
	peoples.Lock()
	defer peoples.Unlock()
	body, err := ioutil.ReadFile(peopleFile)
	if err != nil {
		log.Warn("Error while reading peopleFile", err)
		return
	}
	var got Peoples
	err = json.Unmarshal(body, &got)
	if err != nil {
		log.Warn("Error while unmarshaling peopleFile", err)
	}
	peoples.Data = got.Data
}
