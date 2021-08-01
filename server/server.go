package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Ben-harder/estate/household"
)

func NewHouseholdServer(port int, household household.HouseholdInterface) HouseholdServerInterface {
	svr := &householdServer{}
	svr.port = strconv.Itoa(port)
	svr.household = household
	return svr
}

type HouseholdServerInterface interface {
	ListenAndServe()
}

type householdServer struct {
	port      string
	household household.HouseholdInterface
}

func (svr *householdServer) ListenAndServe() {
	log.Println("Starting server on port", svr.port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello and welcome to the Estate")
		fmt.Fprintln(w, "Current tenants: "+svr.household.String())
	})

	err := http.ListenAndServe(":"+svr.port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
