package math_algorithm

import "testing"

func TestReverse(t *testing.T) {
	var x, h, r int

	x = 123
	h = 321
	r = reverse(x)
	t.Logf("x=%d h=%d r=%d", x, h, r)

	x = -123
	h = -321
	r = reverse(x)
	t.Logf("x=%d h=%d r=%d", x, h, r)

	x = 120
	h = 21
	r = reverse(x)
	t.Logf("x=%d h=%d r=%d", x, h, r)
}
