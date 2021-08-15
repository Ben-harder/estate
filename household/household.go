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
	Members() []MemberInterface
	String() string
	MemberAfter(name string) (MemberInterface, error)
	IsHouseholdMember(name string) bool
	GetHouseholdMember(name string) (MemberInterface, error)
	First() MemberInterface
}

type household struct {
	members memberListInterface
}

func (hHold *household) PrintHouseholdMembers() {
	fmt.Println(hHold.members.string())
}

// Members returns a copy of the household members
func (hHold *household) Members() []MemberInterface {
	membersCopy := make([]MemberInterface, hHold.members.length())
	copy(membersCopy, hHold.members.getMembers())
	return membersCopy
}

// IsHouseholdMember returns true if the given name is one of the members of the household
func (hHold *household) IsHouseholdMember(name string) bool {
	_, err := hHold.members.getMember(name)
	if err != nil {
		return false
	} else {
		return true
	}
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
