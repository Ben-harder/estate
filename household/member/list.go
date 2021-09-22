package member

import (
	"fmt"
	"sort"
	"strings"
)

func NewList(memberNames []string) ListInterface {
	newList := &list{members: make([]MemberInterface, 0)}
	for _, name := range memberNames {
		newMember := NewMember(name)
		newList.members = append(newList.members, newMember)
	}
	newList.sort()
	return newList
}

type ListInterface interface {
	sort()
	String() string
	MemberNames() []string
	first() MemberInterface
	next(member MemberInterface) (MemberInterface, error)
	GetMember(name string) (MemberInterface, error)
	GetMembers() []MemberInterface
	indexOf(member MemberInterface) (int, error)
	Length() int
}

type list struct {
	members []MemberInterface
}

func (mList *list) sort() {
	sort.Slice(mList.members, func(i, j int) bool {
		return mList.members[i].String() < mList.members[j].String()
	})
}

// string converts the list to a comma delimited list of names
func (mList *list) String() string {
	str := ""
	for _, member := range mList.members {
		str = str + member.String() + ", "
	}
	str = strings.Trim(str, ", ")
	return str
}

// memberNames returns a slice of the member names
func (mList *list) MemberNames() []string {
	mListNames := make([]string, 0)
	for _, member := range mList.members {
		mListNames = append(mListNames, member.String())
	}
	return mListNames
}

// members returns the members in the list as a slice
func (mList *list) GetMembers() []MemberInterface {
	return mList.members
}

func (mList *list) first() MemberInterface {
	return mList.members[0]
}

func (mList *list) GetMember(name string) (MemberInterface, error) {
	for _, member := range mList.members {
		if member.String() == name {
			return member, nil
		}
	}
	return nil, fmt.Errorf("household member, %v, not found", name)
}

func (mList *list) next(member MemberInterface) (MemberInterface, error) {
	currIndex, err := mList.indexOf(member)
	if err != nil {
		return nil, err
	}
	if currIndex == (mList.Length() - 1) {
		return mList.members[0], nil
	} else {
		return mList.members[currIndex+1], nil
	}
}

func (mList *list) indexOf(member MemberInterface) (int, error) {
	for i, mem := range mList.members {
		if mem.Equals(member) {
			return i, nil
		}
	}
	return -1, fmt.Errorf("household member not found. Name: %s", member.String())
}

func (mList *list) Length() int {
	return len(mList.members)
}
