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
	Members() []MemberInterface
	String() string
	GetHouseholdMember(name string) (MemberInterface, error)
}

type household struct {
	members memberListInterface
}

// Members returns a copy of the household members
func (hHold *household) Members() []MemberInterface {
	membersCopy := make([]MemberInterface, hHold.members.length())
	copy(membersCopy, hHold.members.getMembers())
	return membersCopy
}

// GetHouseholdMember returns the household member who has the given name or an error if they don't exist
func (hHold *household) GetHouseholdMember(name string) (MemberInterface, error) {
	member, err := hHold.members.getMember(name)
	if err != nil {
		return nil, err
	} else {
		return member, nil
	}
}

// String converts the memberlist to a comma delimited list of names
func (hHold *household) String() string {
	return hHold.members.string()
}
