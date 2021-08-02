package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Ben-harder/estate/household"
	"github.com/Ben-harder/estate/schedule"
	"github.com/Ben-harder/estate/server"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println("Welcome to the Estate")
	names := []string{"Ben Harder", "Andrew Wright", "David Gray", "Dominick Laroche", "Natalia Johnston", "Georgia Stel"}
	theEstate, err := household.NewHousehold(names)
	jobSchedule, err := schedule.NewGarbageSchedule("schedule/schedule.ics")
	job, date := jobSchedule.NextJob()
	fmt.Println(job + " " + date)
	if err != nil {
		log.Fatal("failed to create household. err: ", err)
	}
	svr := server.NewHouseholdServer(8080, theEstate)
	svr.ListenAndServe()
}
