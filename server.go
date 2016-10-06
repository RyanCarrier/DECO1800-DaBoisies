package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/RyanCarrier/DECO1800-DaBoisies/server"
	log "github.com/Sirupsen/logrus"
)

const port = 80

func init() {
	if len(os.Args) > 1 {
		if os.Args[1][0] == '-' {
			cliHelp()
			os.Exit(0)
		}
		filename := os.Args[1]
		if filename[len(filename)-4:] != ".log" {
			filename += ".log"
		}
		f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
		if err != nil {
			log.Error("Couldn't open file")
		} else {
			log.SetOutput(f)
			fmt.Println("Listening on :" + strconv.Itoa(port))
			fmt.Println("Logging to", filename)
		}
	}
	go load()
}

func main() {
	initHelp()
	mux := server.SetupHandlers()
	log.Error(http.ListenAndServe(":"+strconv.Itoa(port), mux))
}

func load() {
	time.Sleep(time.Second * 10)
	http.Get("http://localhost/api/getlist/")
}

func initHelp() {
	log.Info("Listening on http://localhost:" + strconv.Itoa(port))
	log.Info(`To get people use the api;
http://localhost:` + strconv.Itoa(port) + `/api/people/:peopleid
Try using some of the id's from forbes list;
946924
1456345
1163655
1479624
1503477
		`)
}

func cliHelp() {
	fmt.Println("./server [logfile|--help]")
}
