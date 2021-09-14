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

	garbageSchedule, err := schedule.NewGarbageSchedule("Garbage Schedule", "schedule/schedule.ics")
	if err != nil {
		log.Fatal(err)
	}
	choreManager := choreManager.NewChoreManager(theEstate)
	turnList := choreManager.DefaultTurnList()
	choreManager.AddSchedule(garbageSchedule, turnList, 0)
	err = choreManager.SetTurnForSchedule("Garbage Schedule", "Andrew Wright")
	if err != nil {
		log.Fatal(err)
	}

	houseCleanTurns, err := choreManager.CustomTurnList([]string{"Andrew Wright", "Ben Harder"}, []string{"Gus Koenigsfest", "David Gray"}, []string{"Grace Plaseski", "Dominick Laroche"})
	if err != nil {
		log.Fatal(err)
	}
	start := time.Date(2021, 9, 1, 0, 0, 0, 0, time.UTC)
	houseCleanSchedule := schedule.NewCustomSchedule("House Cleaning Schedule", start, 12, 14, []string{"Clean the house common areas"})
	choreManager.AddSchedule(houseCleanSchedule, houseCleanTurns, 0)
	svr := server.NewHouseholdServer(80, choreManager, theEstate)
	svr.ListenAndServe()

}
