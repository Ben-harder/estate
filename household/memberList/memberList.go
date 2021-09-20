package memberList

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Ben-harder/estate/household/member"
)

func NewMemberList(memberNames []string) MemberListInterface {
	newMemberList := &memberList{members: make([]member.MemberInterface, 0)}
	for _, name := range memberNames {
		newMember := member.NewMember(name)
		newMemberList.members = append(newMemberList.members, newMember)
	}
	newMemberList.sort()
	return newMemberList
}

type MemberListInterface interface {
	sort()
	String() string
	MemberNames() []string
	first() member.MemberInterface
	next(member member.MemberInterface) (member.MemberInterface, error)
	GetMember(name string) (member.MemberInterface, error)
	GetMembers() []member.MemberInterface
	indexOf(member member.MemberInterface) (int, error)
	Length() int
}

type memberList struct {
	members []member.MemberInterface
}

func (mList *memberList) sort() {
	sort.Slice(mList.members, func(i, j int) bool {
		return mList.members[i].String() < mList.members[j].String()
	})
}

// string converts the memberlist to a comma delimited list of names
func (mList *memberList) String() string {
	str := ""
	for _, member := range mList.members {
		str = str + member.String() + ", "
	}
	str = strings.Trim(str, ", ")
	return str
}

// memberNames returns a slice of the member names
func (mList *memberList) MemberNames() []string {
	mListNames := make([]string, 0)
	for _, member := range mList.members {
		mListNames = append(mListNames, member.String())
	}
	return mListNames
}

// members returns the members in the list as a slice
func (mList *memberList) GetMembers() []member.MemberInterface {
	return mList.members
}

func (mList *memberList) first() member.MemberInterface {
	return mList.members[0]
}

func (mList *memberList) GetMember(name string) (member.MemberInterface, error) {
	for _, member := range mList.members {
		if member.String() == name {
			return member, nil
		}
	}
	return nil, fmt.Errorf("household member, %v, not found", name)
}

func (mList *memberList) next(member member.MemberInterface) (member.MemberInterface, error) {
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

func (mList *memberList) indexOf(member member.MemberInterface) (int, error) {
	for i, mem := range mList.members {
		if mem.Equals(member) {
			return i, nil
		}
	}
	return -1, fmt.Errorf("household member not found. Name: %s", member.String())
}

func (mList *memberList) Length() int {
	return len(mList.members)
}
