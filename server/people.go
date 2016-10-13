package server

import (
	"sort"
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
func (p *Peoples) clear() {
	p.Lock()
	defer p.Unlock()
	p.Data = make(map[string]FinalPerson)
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
			fp := FinalPerson{Name: n, Image: io}
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

func forceUpdateAll() {
	peoples.clear()
	updateAll()
}
func updateAll() {
	all := names.GetAll()
	wg := sync.WaitGroup{}
	wg.Add(len(all))
	for _, name := range all {
		go syncUpdate(name, &wg)
	}
	wg.Wait()
	log.Info("Update finished.")
}

func syncUpdate(name string, wg *sync.WaitGroup) {
	update(name)
	wg.Done()
}
func update(name string) {
	log.Info("Updating " + name + "...")
	totalOK := true
	updated := false
	got, _ := peoples.get(name)
	if got.Name == "" {
		updated = true
		got.Name = name
	}
	if got.Blurb == "" {
		cr, err := getWeightYear(0, name, 0)
		if err != nil {
			totalOK = false
		} else {
			updated = true
			got.Blurb = getBlurb(cr.Zones)
		}
	}
	if got.Image == "" || got.Image == errorImage {
		io, err := getGoogleImageObject(name)
		updated = true
		if err != nil {
			log.Warn("getting image object error;", err)
			totalOK = false
		}
		got.Image = io
	}
	if totalOK && updated == true {
		peoples.put(name, got)
	}
}

func getBlurb(zs []CleanZone) string {
	cz := CleanZones(zs)
	sort.Sort(cz)
	z, y := len(cz)-1, len(cz)-2
	if float32(cz[z].Total)/float32(cz[y].Total) < float32(0.8) {
		return "Appears most in " + filterName(cz[z].Name) + "s."
	}
	return "Appears most in " + filterName(cz[z].Name) + " and " + filterName(cz[y].Name) + "."
}

func filterName(s string) string {
	if s == "music" {
		return s
	}
	return s + "s"
}
