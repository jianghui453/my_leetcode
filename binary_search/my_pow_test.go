package binary_search

import "testing"

func TestMyPow(t *testing.T) {
	var x, h, r float64
	var n int

	x = 2.0
	n = -2
	h = 0.25
	r = myPow(x, n)
	t.Logf("x=%g n=%d h=%g r=%g", x, n, h, r)

	x = 2.0
	n = 10
	h = 1024
	r = myPow(x, n)
	t.Logf("x=%g n=%d h=%g r=%g", x, n, h, r)

	x = 2.1
	n = 3
	h = 9.261
	r = myPow(x, n)
	t.Logf("x=%g n=%d h=%g r=%g", x, n, h, r)
}
