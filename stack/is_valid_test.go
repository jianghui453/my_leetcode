package stack

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	var s string
	var h, r bool

	s = "()[]{}"
	h = true
	r = isValid(s)
	t.Logf("%t s=%s h=%t r=%t", r==h, s, h, r)

	s = "{[]}"
	h = true
	r = isValid(s)
	t.Logf("%t s=%s h=%t r=%t", r==h, s, h, r)

	s = "([)]"
	h = false
	r = isValid(s)
	t.Logf("%t s=%s h=%t r=%t", r==h, s, h, r)
}
