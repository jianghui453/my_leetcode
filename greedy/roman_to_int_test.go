package greedy

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	var s string
	var h, r int

	s = "III"
	h = 3
	r = romanToInt(s)
	t.Logf("%t s=%s h=%d r=%d", r == h, s, h, r)

	s = "IV"
	h = 4
	r = romanToInt(s)
	t.Logf("%t s=%s h=%d r=%d", r == h, s, h, r)

	s = "IX"
	h = 9
	r = romanToInt(s)
	t.Logf("%t s=%s h=%d r=%d", r == h, s, h, r)

	s = "LVIII"
	h = 58
	r = romanToInt(s)
	t.Logf("%t s=%s h=%d r=%d", r == h, s, h, r)

	s = "MCMXCIV"
	h = 1994
	r = romanToInt(s)
	t.Logf("%t s=%s h=%d r=%d", r == h, s, h, r)
}
