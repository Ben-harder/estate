package household

import (
	"fmt"
	"sort"
	"strings"
)

func newMemberList(memberNames []string) memberListInterface {
	newMemberList := &memberList{members: make([]MemberInterface, 0)}
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
	memberNames() []string
	first() MemberInterface
	next(member MemberInterface) (MemberInterface, error)
	getMember(name string) (MemberInterface, error)
	getMembers() []MemberInterface
	indexOf(member MemberInterface) (int, error)
	length() int
}

type memberList struct {
	members []MemberInterface
}

func (mList *memberList) sort() {
	sort.Slice(mList.members, func(i, j int) bool {
		return mList.members[i].String() < mList.members[j].String()
	})
}

// string converts the memberlist to a comma delimited list of names
func (mList *memberList) string() string {
	str := ""
	for _, member := range mList.members {
		str = str + member.String() + ", "
	}
	str = strings.Trim(str, ", ")
	return str
}

// memberNames returns a slice of the member names
func (mList *memberList) memberNames() []string {
	mListNames := make([]string, 0)
	for _, member := range mList.members {
		mListNames = append(mListNames, member.String())
	}
	return mListNames
}

// members returns the members in the list as a slice
func (mList *memberList) getMembers() []MemberInterface {
	return mList.members
}

func (mList *memberList) first() MemberInterface {
	return mList.members[0]
}

func (mList *memberList) getMember(name string) (MemberInterface, error) {
	for _, member := range mList.members {
		if member.String() == name {
			return member, nil
		}
	}
	return nil, fmt.Errorf("household member, %v, not found", name)
}

func (mList *memberList) next(member MemberInterface) (MemberInterface, error) {
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

func (mList *memberList) indexOf(member MemberInterface) (int, error) {
	for i, mem := range mList.members {
		if mem.Equals(member) {
			return i, nil
		}
	}
	return -1, fmt.Errorf("household member not found. Name: %s", member.String())
}

func (mList *memberList) length() int {
	return len(mList.members)
}
