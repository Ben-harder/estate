package choreManager

import (
	"fmt"

	"github.com/Ben-harder/estate/household/member"
)

func NewChoreTurnList(turnList [][]member.MemberInterface) choreTurnListInterface {
	return &choreTurnList{turnList: turnList, index: 0}
}

// choreTurnListInterface manages the state of a chores turn list. Ie, who's turn it is to complete a job and who is next up.
type choreTurnListInterface interface {
	advanceToNext()
	whoseTurn() []member.MemberInterface
	setTurn(index int) error
	setTurnWithMember(member member.MemberInterface) error
	currIndex() int
}

type choreTurnList struct {
	turnList [][]member.MemberInterface
	index    int
}

// advanceToNext will increment one to the index for the turn list
func (chrTurnList *choreTurnList) advanceToNext() {
	chrTurnList.index = chrTurnList.index + 1%len(chrTurnList.turnList)
}

func (chrTurnList *choreTurnList) whoseTurn() []member.MemberInterface {
	return chrTurnList.turnList[chrTurnList.index]
}

func (chrTurnList *choreTurnList) setTurn(index int) error {
	if index > len(chrTurnList.turnList)-1 || index < 0 {
		return fmt.Errorf("turn index out of bounds")
	} else {
		chrTurnList.index = index
		return nil
	}
}

func (chrTurnList *choreTurnList) setTurnWithMember(member member.MemberInterface) error {
	for i, sublist := range chrTurnList.turnList {
		for _, sublistMember := range sublist {
			if sublistMember.Equals(member) {
				chrTurnList.setTurn(i)
				return nil
			}
		}
	}
	return fmt.Errorf("could not find member, %v, in turn lists", member.String())
}

func (chrTurnList *choreTurnList) currIndex() int {
	return chrTurnList.index
}
