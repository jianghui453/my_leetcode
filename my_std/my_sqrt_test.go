package my_std

import "testing"

func TestMySqrt(t *testing.T) {
	var x, hope, ret int

	x = 4
	hope = 2
	ret = mySqrt(x)
	t.Logf("x=%d hope=%d ret=%d", x, hope, ret)

	x = 8
	hope = 2
	ret = mySqrt(x)
	t.Logf("x=%d hope=%d ret=%d", x, hope, ret)
}
