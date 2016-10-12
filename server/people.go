package server

import (
	"sync"

	log "github.com/Sirupsen/logrus"
)

//Peoples is a struct to remember the names already found
type Peoples struct {
	sync.Mutex
	Data map[string]FinalPerson `json:"data"`
}

var peoples = newPeoples()

func newPeoples() *Peoples {
	return &Peoples{
		Data: make(map[string]FinalPerson),
	}
}

func (p *Peoples) get(i string) (FinalPerson, bool) {
	p.Lock()
	defer p.Unlock()
	got, ok := p.Data[i]
	return got, ok
}

func (p *Peoples) put(s string, fp FinalPerson) {
	p.Lock()
	defer p.Unlock()
	p.Data[s] = fp
}

func getPeopleI(IDs []int) []FinalPerson {
	final := make([]FinalPerson, len(IDs))
	names := getNamesI(IDs)
	for i, n := range names {
		if person, ok := peoples.get(n); ok {
			final[i] = person
		} else {
			io, err := getGoogleImageObject(n)
			fp := FinalPerson{Name: n, ImageObject: io}
			if err != nil {
				log.Warn("getting image object error;", err)
			} else {
				peoples.put(n, fp)
			}
			final[i] = fp
		}
	}
	return final
}
