package string

import (
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	var s string
	var h, r int

	s = "hello world"
	h = 5
	r = lengthOfLastWord(s)
	t.Logf("%t s=%s h=%d r=%d", r == h, s, h, r)
}
