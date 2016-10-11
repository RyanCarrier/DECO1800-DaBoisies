package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

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
	go server.DelayLoad(3)
}

func main() {
	initHelp()
	mux := server.SetupHandlers()
	log.Error(http.ListenAndServe(":"+strconv.Itoa(port), mux))
}

func initHelp() {
	log.Info("Listening on http://localhost:" + strconv.Itoa(port))
}

func cliHelp() {
	fmt.Println("./server [logfile|--help]")
}
