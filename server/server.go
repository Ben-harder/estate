package server

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Ben-harder/estate/choreManager"
	"github.com/Ben-harder/estate/household"
	"github.com/Ben-harder/estate/html"
)

type Peena struct {
	Length int
}

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
	
	html.MainPage(svr.household.MemberNames(), svr.choreManager.Chores(), w)
}
