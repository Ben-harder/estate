package household

import (
	"strings"
)

// newMember creates a new household member. Note that only Firstname or Firstname Lastname formats are accepted. Any middle names or abreviations will be ignored.
func newMember(name string) MemberInterface {
	firstName, lastName := cleanName(name)
	return &member{firstName: firstName, lastName: lastName}
}

type MemberInterface interface {
	String() string
	Equals(member MemberInterface) bool
}

type member struct {
	firstName string
	lastName  string
}

func (mem *member) String() string {
	return strings.TrimSpace(mem.firstName + " " + mem.lastName)
}

func (mem *member) Equals(otherMember MemberInterface) bool {
	if mem.String() == otherMember.String() {
		return true
	} else {
		return false
	}
}

func cleanName(name string) (string, string) {
	trimmed := strings.TrimSpace(name)
	titled := strings.Title(trimmed)
	split := make([]string, 0)
	split = strings.Split(titled, " ")
	if len(split) > 1 {
		return split[0], split[len(split)-1]
	}
	return split[0], ""
}
