package dynamic_programming

import "testing"

func TestNumDecodings(t *testing.T) {
	var s string
	var h, r int

	s = "12"
	h = 2
	r = numDecodings(s)
	t.Logf("%t s=%s h=%d r=%d", h == r, s, h, r)

	s = "226"
	h = 3
	r = numDecodings(s)
	t.Logf("%t s=%s h=%d r=%d", h == r, s, h, r)

	s = "2222"
	h = 5
	r = numDecodings(s)
	t.Logf("%t s=%s h=%d r=%d", h == r, s, h, r)

	s = "227"
	h = 2
	r = numDecodings(s)
	t.Logf("%t s=%s h=%d r=%d", h == r, s, h, r)

	s = "0"
	h = 0
	r = numDecodings(s)
	t.Logf("%t s=%s h=%d r=%d", h == r, s, h, r)

	s = "10"
	h = 1
	r = numDecodings(s)
	t.Logf("%t s=%s h=%d r=%d", h == r, s, h, r)

	s = "26"
	h = 2
	r = numDecodings(s)
	t.Logf("%t s=%s h=%d r=%d", h == r, s, h, r)

	s = "101"
	h = 1
	r = numDecodings(s)
	t.Logf("%t s=%s h=%d r=%d", h == r, s, h, r)

	s = "01"
	h = 0
	r = numDecodings(s)
	t.Logf("%t s=%s h=%d r=%d", h == r, s, h, r)
}
