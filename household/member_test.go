package household

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	n1 := newMember("Bob T. Johnson")
	assert.Equal(t, n1.String(), "Bob Johnson")
	n2 := newMember("Donald")
	assert.Equal(t, "Donald", n2.String())
}

func TestEquals(t *testing.T) {
	n1 := newMember("John Doe")
	n2 := newMember("John A. Doe")
	assert.True(t, n1.equals(n2))
}
