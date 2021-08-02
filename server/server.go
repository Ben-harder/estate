package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Ben-harder/estate/household"
	"github.com/Ben-harder/estate/schedule"
)

func NewHouseholdServer(port int, household household.HouseholdInterface) HouseholdServerInterface {
	svr := &householdServer{}
	svr.port = strconv.Itoa(port)
	svr.household = household
	jobSchedule, err := schedule.NewGarbageSchedule(household, "schedule/schedule.ics")
	if err != nil {
		log.Fatal(err)
	}
	jobSchedule.SetTurn("Dominick Laroche")
	svr.schedule = jobSchedule
	return svr
}

type HouseholdServerInterface interface {
	ListenAndServe()
}

type householdServer struct {
	port      string
	household household.HouseholdInterface
	schedule  schedule.ScheduleInterface
}

func (svr *householdServer) ListenAndServe() {
	log.Println("Starting server on port", svr.port)

	job, date, whoseTurn := svr.schedule.NextJob()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello and welcome to the Estate")
		fmt.Fprintf(w, "Residents of The Estate: %v\n", svr.household.String())
		fmt.Fprintf(w, "Garbage: %v's turn on %v to take out %v\n", whoseTurn, date, job)
	})

	err := http.ListenAndServe(":"+svr.port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
