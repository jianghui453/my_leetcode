package string

import (
	"testing"
)

func TestConvert(t *testing.T)  {
	var numRows int
	var s, r, h string

	s = "LEETCODEISHIRING"
	numRows = 3
	h = "LCIRETOESIIGEDHN"
	r = convert(s, numRows)
	t.Logf("%t s=%s numRows=%d h=%s r=%s", h==r, s, numRows, h, r)

	s = "LEETCODEISHIRING"
	numRows = 4
	h = "LDREOEIIECIHNTSG"
	r = convert(s, numRows)
	t.Logf("%t s=%s numRows=%d h=%s r=%s", h==r, s, numRows, h, r)
}