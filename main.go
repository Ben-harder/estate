package main

import (
	"fmt"
	"log"

	"github.com/Ben-harder/estate/household"
)

func main() {
	fmt.Println("Welcome to the Estate")
	names := []string{"Ben Harder", "Andrew Wright", "David Gray"}
	theEstate, err := household.NewHousehold(names)
	if err != nil {
		log.Fatal("failed to create household. err: ", err)
	}
	theEstate.PrintHouseholdMembers()
	// svr := server.NewServer(8080)
	// svr.ListenAndServe()
}
