package dynamic_programming

import "testing"

func TestNumDistinct(te *testing.T) {
	var s, t string
	var h, r int

	s = "rabbbit"
	t = "rabbit"
	h = 3
	r = numDistinct(s, t)
	te.Logf("%t s=%s t=%s h=%d r=%d", h == r, s, t, h, r)

	s = "babgbag"
	t = "bag"
	h = 5
	r = numDistinct(s, t)
	te.Logf("%t s=%s t=%s h=%d r=%d", h == r, s, t, h, r)
}
