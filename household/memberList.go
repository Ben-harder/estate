package household

import (
	"fmt"
	"sort"
)

func newMemberList(memberNames []string) memberListInterface {
	newMemberList := &memberList{members: make([]memberInterface, 0)}
	for _, name := range memberNames {
		newMember := newMember(name)
		newMemberList.members = append(newMemberList.members, newMember)
	}
	newMemberList.sort()
	return newMemberList
}

type memberListInterface interface {
	sort()
	string() string
	next(member memberInterface) (memberInterface, error)
	getMember(name string) (memberInterface, error)
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

func (mList *memberList) getMember(name string) (memberInterface, error) {
	mem := newMember(name)
	i, err := mList.indexOf(mem)
	if err != nil {
		return nil, err
	} else {
		return mList.members[i], nil
	}
}

func (mList *memberList) next(member memberInterface) (memberInterface, error) {
	currIndex, err := mList.indexOf(member)
	if err != nil {
		return nil, err
	}
	if currIndex == (mList.length() - 1) {
		return mList.members[0], nil
	} else {
		return mList.members[currIndex+1], nil
	}
}

func (mList *memberList) indexOf(member memberInterface) (int, error) {
	for i, mem := range mList.members {
		if mem.equals(member) {
			return i, nil
		}
	}
	return -1, fmt.Errorf("household member not found. Name: %s", member.string())
}

func (mList *memberList) length() int {
	return len(mList.members)
}
