package household

func newMember(firstName string, lastName string) memberInterface {
	return &member{firstName: firstName, lastName: lastName}
}

type memberInterface interface {
	string() string
}

type member struct {
	firstName string
	lastName  string
}

func (mem *member) string() string {
	return mem.firstName + " " + mem.lastName
}
