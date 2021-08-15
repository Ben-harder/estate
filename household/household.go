package household

import (
	"fmt"
	"log"
)

func NewHousehold(memberNames []string) (HouseholdInterface, error) {
	if len(memberNames) < 1 {
		return nil, fmt.Errorf("please supply at least one household member")
	}

	log.Println("Created new household with members: ", memberNames)
	return &household{members: newMemberList(memberNames)}, nil
}

type HouseholdInterface interface {
	PrintHouseholdMembers()
	String() string
	MemberAfter(name string) (MemberInterface, error)
	First() MemberInterface
}

type household struct {
	members memberListInterface
}

func (hHold *household) PrintHouseholdMembers() {
	fmt.Println(hHold.members.string())
}

func (hHold *household) String() string {
	return hHold.members.string()
}

// Finds the next member of the household in alphabetical order
func (hHold *household) MemberAfter(name string) (MemberInterface, error) {
	mem, err := hHold.members.getMember(name)
	if err != nil {
		return nil, err
	}
	nextMember, err := hHold.members.next(mem)
	if err != nil {
		return nil, err
	}
	return nextMember, nil
}

func (hHold *household) First() MemberInterface {
	return hHold.members.first()
}
