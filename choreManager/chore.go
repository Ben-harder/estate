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
	String() string
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
	return "Date: " + chr.date + "| Whose turn: " + whoseTurns
}

func (chr *chore) SetTurn(members []household.MemberInterface) {
	chr.whoseTurn = members
}

// mapHouseholdToNames takes a list of member interfaces and converts it to a slice of their names
func mapHouseholdToNames(members []household.MemberInterface) []string {
	names := make([]string, len(members))
	for _, mem := range members {
		names = append(names, mem.String())
	}
	return names
}
