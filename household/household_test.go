package household

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// testMemberAfter tests the loops nature of member after
func testMemberAfter(t *testing.T, h HouseholdInterface) {
	m, err := h.MemberAfter("John Doe")
	assert.Nil(t, err)
	assert.Equal(t, "Maggie Karl", m.String())
	m, err = h.MemberAfter(m.String())
	assert.Nil(t, err)
	assert.Equal(t, "Abe Peters", m.String())
	m, err = h.MemberAfter(m.String())
	assert.Nil(t, err)
	assert.Equal(t, "John Doe", m.String())
	_, err = h.MemberAfter("Doesn't Exist")
	assert.Error(t, err)
}

func TestHousehold(t *testing.T) {
	// <setup code>
	h, err := NewHousehold([]string{"John Doe", "Maggie Karl", "Abe Peters"})
	assert.Nil(t, err)

	t.Run("Test member list creation", func(t *testing.T) {
		testMemberAfter(t, h)
	})

	// <tear-down code>
}
