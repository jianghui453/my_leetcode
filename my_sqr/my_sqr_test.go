package my_sqr

import "testing"

func TestMySqrt(t *testing.T) {
	var n, hope, ret float64

	n = 2.0
	hope = 1.4142135623
	ret = mySqr(n)
	t.Logf("n=%g hope=%g ret=%g", n, hope, ret)
}