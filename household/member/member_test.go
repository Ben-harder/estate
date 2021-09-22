package member

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	n1 := NewMember("Bob T. Johnson")
	assert.Equal(t, n1.String(), "Bob Johnson")
	n2 := NewMember("Donald")
	assert.Equal(t, "Donald", n2.String())
}

func TestEquals(t *testing.T) {
	n1 := NewMember("John Doe")
	n2 := NewMember("John A. Doe")
	assert.True(t, n1.Equals(n2))
}
