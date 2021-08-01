package household

import (
	"sort"
	"strings"
)

func newMemberList(memberNames []string) memberListInterface {
	newMemberList := &memberList{members: make([]memberInterface, 0)}
	for _, m := range memberNames {
		firstName, lastName := cleanName(m)
		newMember := newMember(firstName, lastName)
		newMemberList.members = append(newMemberList.members, newMember)
	}
	newMemberList.sort()
	return newMemberList
}

type memberListInterface interface {
	sort()
	string() string
}

type memberList struct {
	members []memberInterface
}

func (mList *memberList) sort() {
	sort.Slice(mList.members, func(i, j int) bool {
		return mList.members[i].string() < mList.members[j].string()
	})
}

func (mList *memberList) string() string {
	str := ""
	for _, member := range mList.members {
		str = str + member.string() + " "
	}
	return str
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
