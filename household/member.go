package household

import "strings"

func newMember(name string) memberInterface {
	firstName, lastName := cleanName(name)
	return &member{firstName: firstName, lastName: lastName}
}

type memberInterface interface {
	string() string
	equals(member memberInterface) bool
}

type member struct {
	firstName string
	lastName  string
}

func (mem *member) string() string {
	return mem.firstName + " " + mem.lastName
}

func (mem *member) equals(otherMember memberInterface) bool {
	if mem.string() == otherMember.string() {
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
