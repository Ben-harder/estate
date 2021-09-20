package household

import (
	"fmt"
	"log"

	"github.com/Ben-harder/estate/household/member"
)

func NewHousehold(memberNames []string) (HouseholdInterface, error) {
	if len(memberNames) < 1 {
		return nil, fmt.Errorf("please supply at least one household member")
	}

	log.Println("Created new household with members: ", memberNames)
	return &household{members: member.NewList(memberNames)}, nil
}

type HouseholdInterface interface {
	Members() []member.MemberInterface
	String() string
	GetHouseholdMember(name string) (member.MemberInterface, error)
}

type household struct {
	members member.ListInterface
}

// Members returns a copy of the household members
func (hHold *household) Members() []member.MemberInterface {
	membersCopy := make([]member.MemberInterface, hHold.members.Length())
	copy(membersCopy, hHold.members.GetMembers())
	return membersCopy
}

// GetHouseholdMember returns the household member who has the given name or an error if they don't exist
func (hHold *household) GetHouseholdMember(name string) (member.MemberInterface, error) {
	member, err := hHold.members.GetMember(name)
	if err != nil {
		return nil, err
	} else {
		return member, nil
	}
}

// String converts the memberlist to a comma delimited list of names
func (hHold *household) String() string {
	return hHold.members.String()
}
