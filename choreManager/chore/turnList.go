package chore

import (
	"fmt"

	"github.com/Ben-harder/estate/household/member"
)

func NewChoreTurnList(turnTeams [][]member.MemberInterface) turnListInterface {
	return &turnList{turnTeams: turnTeams, index: 0}
}

// turnListInterface manages the state of a chores turn list. Ie, who's turn it is to complete a job and who is next up.
type turnListInterface interface {
	advanceToNext()
	whoseTurn() []member.MemberInterface
	setTurn(index int) error
	setTurnWithMember(member member.MemberInterface) error
	currIndex() int
}

type turnList struct {
	turnTeams [][]member.MemberInterface
	index     int
}

// advanceToNext will increment one to the index for the turn list
func (chrTurnList *turnList) advanceToNext() {
	chrTurnList.index = (chrTurnList.index + 1) % len(chrTurnList.turnTeams)
}

func (chrTurnList *turnList) whoseTurn() []member.MemberInterface {
	return chrTurnList.turnTeams[chrTurnList.index]
}

func (chrTurnList *turnList) setTurn(index int) error {
	if index > len(chrTurnList.turnTeams)-1 || index < 0 {
		return fmt.Errorf("turn index out of bounds")
	} else {
		chrTurnList.index = index
		return nil
	}
}

func (chrTurnList *turnList) setTurnWithMember(member member.MemberInterface) error {
	for i, sublist := range chrTurnList.turnTeams {
		for _, sublistMember := range sublist {
			if sublistMember.Equals(member) {
				chrTurnList.setTurn(i)
				return nil
			}
		}
	}
	return fmt.Errorf("could not find member, %v, in turn lists", member.String())
}

func (chrTurnList *turnList) currIndex() int {
	return chrTurnList.index
}
