package chore

import (
	"strings"

	"github.com/Ben-harder/estate/household/member"
)

func NewChore(scheduleName string, responsibilities string, date string, turnList [][]member.MemberInterface) ChoreInterface {
	choreTurnList := NewChoreTurnList(turnList)
	return &chore{scheduleName: scheduleName, responsibilities: responsibilities, date: date, choreTurnList: choreTurnList}
}

type ChoreInterface interface {
	String() string
	WhoseTurn() []member.MemberInterface
	Schedule() string
	Responsibilities() string
	Date() string
	SetTurn(index int) error
	SetTurnWithMember(member member.MemberInterface) error
	SetResponsibilities(string)
	SetDate(date string)
	AdvanceToNextTurn()
}

type chore struct {
	date             string
	responsibilities string
	choreTurnList    choreTurnListInterface
	scheduleName     string
}

// String returns the string representation of a chore
func (chr *chore) String() string {
	whoseTurns := strings.Join(mapHouseholdToNames(chr.choreTurnList.whoseTurn()), ", ")
	whoseTurns = strings.Trim(whoseTurns, ", ")
	return "Responsibilities: " + chr.responsibilities + " | Date: " + chr.date + " | Whose turn: " + whoseTurns
}

// WhoseTurn returns the person responsible for this chore
func (chr *chore) WhoseTurn() []member.MemberInterface {
	return chr.choreTurnList.whoseTurn()
}

// Schedule returns the name of the schedule this chore is from
func (chr *chore) Schedule() string {
	return chr.scheduleName
}

// Responsibilities returns the responsibilities for this chore
func (chr *chore) Responsibilities() string {
	return chr.responsibilities
}

// Date returns the date for this chore
func (chr *chore) Date() string {
	return chr.date
}

// SetTurn will set who is responsible for the chore using an index
func (chr *chore) SetTurn(index int) error {
	err := chr.choreTurnList.setTurn(index)
	if err != nil {
		return err
	}
	return nil
}

// SetTurnWithMember will search the turn list for the given member and set those person(s) to the current turn
func (chr *chore) SetTurnWithMember(member member.MemberInterface) error {
	err := chr.choreTurnList.setTurnWithMember(member)
	if err != nil {
		return err
	} else {
		return nil
	}
}

// SetResponsibilities will set the responsibilities for this chore
func (chr *chore) SetResponsibilities(responsibilities string) {
	chr.responsibilities = responsibilities
}

// SetDate sets the date that this chore should be done by
func (chr *chore) SetDate(date string) {
	chr.date = date
}

// AdvanceToNextTurn will advance the chore to the next people in the turn list
func (chr *chore) AdvanceToNextTurn() {
	chr.choreTurnList.advanceToNext()
}

// mapHouseholdToNames takes a list of member interfaces and converts it to a slice of their names
func mapHouseholdToNames(members []member.MemberInterface) []string {
	names := make([]string, 0)
	for _, mem := range members {
		names = append(names, mem.String())
	}
	return names
}
