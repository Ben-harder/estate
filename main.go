package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Ben-harder/estate/choreManager"
	"github.com/Ben-harder/estate/household"
	"github.com/Ben-harder/estate/schedule"
	"github.com/Ben-harder/estate/server"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println("Welcome to the Estate")
	names := []string{"Ben Harder", "Andrew Wright", "David Gray", "Dominick Laroche", "Gus Koenigsfest", "Grace Plaseski"}
	theEstate, err := household.NewHousehold(names)
	if err != nil {
		log.Fatal("failed to create household. err: ", err)
	}
	jobSchedule, err := schedule.NewGarbageSchedule("garbage schedule", theEstate, "schedule/schedule.ics")
	if err != nil {
		log.Fatal(err)
	}
	choreManager := choreManager.NewChoreManager(theEstate)
	turnList := choreManager.DefaultTurnList()
	choreManager.AddSchedule(jobSchedule, turnList, 0)
	choreManager.SetTurnForSchedule("garbage schedule", "Gus Koenigsfest")
	svr := server.NewHouseholdServer(80, choreManager)
	svr.ListenAndServe()
}
