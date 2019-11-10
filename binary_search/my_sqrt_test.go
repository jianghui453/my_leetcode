package binary_search

import "testing"

func TestMySqrt(t *testing.T) {
	var x, hope, ret int

	x = 4
	hope = 2
	ret = mySqrt(x)
	t.Logf("%t x=%d hope=%d ret=%d", ret==hope, x, hope, ret)

	x = 5
	hope = 2
	ret = mySqrt(x)
	t.Logf("%t x=%d hope=%d ret=%d", ret==hope, x, hope, ret)

	x = 3
	hope = 1
	ret = mySqrt(x)
	t.Logf("%t x=%d hope=%d ret=%d", ret==hope, x, hope, ret)

	x = 8
	hope = 2
	ret = mySqrt(x)
	t.Logf("%t x=%d hope=%d ret=%d", ret==hope, x, hope, ret)

	x = 1
	hope = 1
	ret = mySqrt(x)
	t.Logf("%t x=%d hope=%d ret=%d", ret==hope, x, hope, ret)

	x = 101
	hope = 10
	ret = mySqrt(x)
	t.Logf("%t x=%d hope=%d ret=%d", ret==hope, x, hope, ret)
}
