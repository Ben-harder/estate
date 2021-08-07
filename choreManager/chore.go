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

func (chr *chore) String() string {
	return "Date: " + chr.date + "\nWhose turn: " + strings.Join(mapToName(chr.whoseTurn), ", ")
}

func (chr *chore) SetTurn(members []household.MemberInterface) {
	chr.whoseTurn = members
}

func mapToName(members []household.MemberInterface) []string {
	names := make([]string, len(members))
	for _, mem := range members {
		names = append(names, mem.String())
	}
	return names
}
