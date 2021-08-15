package choreManager

import (
	"strings"

	"github.com/Ben-harder/estate/household"
)

func NewChore(scheduleName string, responsibilities string, date string) ChoreInterface {
	return &chore{scheduleName: scheduleName, responsibilities: responsibilities, date: date}
}

type ChoreInterface interface {
	SetTurn(members []household.MemberInterface)
	Schedule() string
	String() string
	WhoseTurn() []household.MemberInterface
}

type chore struct {
	date             string
	responsibilities string
	whoseTurn        []household.MemberInterface
	scheduleName     string
}

// String returns the string representation of a chore
func (chr *chore) String() string {
	whoseTurns := strings.Join(mapHouseholdToNames(chr.whoseTurn), ", ")
	whoseTurns = strings.Trim(whoseTurns, ", ")
	return "Responsibilities: " + chr.responsibilities + " | Date: " + chr.date + " | Whose turn: " + whoseTurns
}

// WhoseTurn returns the person responsible for this chore
func (chr *chore) WhoseTurn() []household.MemberInterface {
	return chr.whoseTurn
}

// Schedule returns the name of the schedule this chore is from
func (chr *chore) Schedule() string {
	return chr.scheduleName
}

func (chr *chore) SetTurn(members []household.MemberInterface) {
	chr.whoseTurn = members
}

// mapHouseholdToNames takes a list of member interfaces and converts it to a slice of their names
func mapHouseholdToNames(members []household.MemberInterface) []string {
	names := make([]string, 0)
	for _, mem := range members {
		names = append(names, mem.String())
	}
	return names
}
