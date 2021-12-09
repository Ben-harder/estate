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
	names := []string{"Ben Harder", "Andrew Wright", "David Gray", "Dominick Laroche", "Georgia Stel", "Natalia Johnston"}
	theEstate, err := household.NewHousehold(names)
	if err != nil {
		log.Fatal("failed to create household. err: ", err)
	}

	garbageSchedule, err := schedule.NewGarbageSchedule("Garbage", "schedule/schedule.ics")
	if err != nil {
		log.Fatal(err)
	}
	choreManager := choreManager.NewChoreManager(theEstate)
	turnList := choreManager.DefaultTurnList()
	choreManager.AddSchedule(garbageSchedule, turnList, 0)
	start := time.Date(2021, 9, 1, 0, 0, 0, 0, time.UTC)

	mainFloorSchedule := schedule.NewCustomSchedule("Main Floor", start, 12, 7, []string{"Vacuum or sweep the main floor excluding the entry, kitchen, and dining room"})
	choreManager.AddSchedule(mainFloorSchedule, choreManager.DefaultTurnList(), 0)
	entrySchedule := schedule.NewCustomSchedule("Entry", start, 12, 7, []string{"Sweep the entry and put away the shoes"})
	choreManager.AddSchedule(entrySchedule, choreManager.DefaultTurnList(), 1)
	kitchenSchedule := schedule.NewCustomSchedule("Kitchen", start, 12, 7, []string{"Wipe down counters, stove, and table; clean the sink; vacuum / sweep the floor"})
	choreManager.AddSchedule(kitchenSchedule, choreManager.DefaultTurnList(), 2)
	ragsSchedule := schedule.NewCustomSchedule("Rags", start, 12, 7, []string{"Wash and fold the kitchen rags"})
	choreManager.AddSchedule(ragsSchedule, choreManager.DefaultTurnList(), 3)
	upstairsSchedule := schedule.NewCustomSchedule("Upstairs", start, 12, 7, []string{"Vacuum or sweep, clean the table"})
	choreManager.AddSchedule(upstairsSchedule, choreManager.DefaultTurnList(), 4)
	downstairsSchedule := schedule.NewCustomSchedule("Downstairs", start, 12, 14, []string{"Vacuum or sweep"})
	choreManager.AddSchedule(downstairsSchedule, choreManager.DefaultTurnList(), 5)
	choreManager.UpdateChoresIfOld()
	svr := server.NewHouseholdServer(80, choreManager, theEstate)
	svr.ListenAndServe()

}
