package household

type member struct {
	firstName string
	lastName  string
}

func (mem *member) string() string {
	return mem.firstName + " " + mem.lastName
}
