package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Ben-harder/estate/choreManager"
	"github.com/Ben-harder/estate/household"
)

func NewHouseholdServer(port int, choreManager choreManager.ChoreManagerInterface, household household.HouseholdInterface) HouseholdServerInterface {
	return &householdServer{port: strconv.Itoa(port), choreManager: choreManager, household: household}
}

type HouseholdServerInterface interface {
	ListenAndServe()
	mainPageHandler(w http.ResponseWriter, r *http.Request)
}

type householdServer struct {
	port         string
	choreManager choreManager.ChoreManagerInterface
	household    household.HouseholdInterface
}

func (svr *householdServer) ListenAndServe() {
	log.Println("Starting server on port", svr.port)

	checkJobs := time.NewTicker(1 * time.Minute)
	go func() {
		for {
			select {
			case <-checkJobs.C:
				svr.choreManager.UpdateChoresIfOld()
			}
		}
	}()

	http.HandleFunc("/", svr.mainPageHandler)

	err := http.ListenAndServe(":"+svr.port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (svr *householdServer) mainPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	// job, date, whoseTurn := svr.schedule.NextJob()
	fmt.Fprintln(w, "Hello and welcome to The Estate")
	fmt.Fprintf(w, "Residents of The Estate: %v\n", svr.household.String())
	// fmt.Fprintf(w, "Garbage: %v's turn on %v to take out %v\n", whoseTurn, date, job)
	fmt.Fprintf(w, "\nChores:\n"+strings.Join(svr.choreManager.Chores(), "\n"))
}
