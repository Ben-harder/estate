package household

import (
	"fmt"
	"log"
)

func NewHousehold(memberNames []string) (Household, error) {
	if len(memberNames) < 1 {
		return nil, fmt.Errorf("please supply at least one household member")
	}

	log.Println("Created new household with members: ", memberNames)
	return &household{members: newMemberList(memberNames)}, nil
}

type Household interface {
	PrintHouseholdMembers()
}

type household struct {
	members memberListInterface
}

func (hHold *household) PrintHouseholdMembers() {
	fmt.Println(hHold.members.string())
}
