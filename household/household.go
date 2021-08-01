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
