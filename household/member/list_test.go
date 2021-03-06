package member

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// newMemberList sorts the incoming list of names
func testSortOnCreation(t *testing.T, ml ListInterface) {
	assert.Equal(t, "Abe Peters", ml.first().String())
}

func testGetMember(t *testing.T, ml ListInterface) {
	_, err := ml.GetMember("Doesn't Exist")
	assert.Error(t, err)
	m, err := ml.GetMember("Maggie Karl")
	assert.Nil(t, err)
	assert.Equal(t, "Maggie Karl", m.String())
}

func TestMemberList(t *testing.T) {
	// <setup code>
	ml := NewList([]string{"John Doe", "Maggie Karl", "Abe Peters"})
	t.Run("Test member list creation", func(t *testing.T) {
		testSortOnCreation(t, ml)
	})
	t.Run("Test get member", func(t *testing.T) {
		testSortOnCreation(t, ml)
	})
	// <tear-down code>
}
