package math_algorithm

import (
	"testing"
)

func Test_findKthNumber(t *testing.T) {
	var n, k, hope, ret int

	n, k = 13, 2
	hope, ret = 10, findKthNumber(n, k)
	t.Logf("%t n=%d k=%d hope=%d ret=%d", ret == hope, n, k, hope, ret)

	n, k = 13, 9
	hope, ret = 5, findKthNumber(n, k)
	t.Logf("%t n=%d k=%d hope=%d ret=%d", ret == hope, n, k, hope, ret)

	n, k = 13, 1
	hope, ret = 1, findKthNumber(n, k)
	t.Logf("%t n=%d k=%d hope=%d ret=%d", ret == hope, n, k, hope, ret)

	n, k = 100, 10
	hope, ret = 17, findKthNumber(n, k)
	t.Logf("%t n=%d k=%d hope=%d ret=%d", ret == hope, n, k, hope, ret)
	
	n, k = 100, 90
	hope, ret = 9, findKthNumber(n, k)
	t.Logf("%t n=%d k=%d hope=%d ret=%d", ret == hope, n, k, hope, ret)

	n, k = 957747794, 424238336
	hope, ret = 481814499, findKthNumber(n, k)
	t.Logf("%t n=%d k=%d hope=%d ret=%d", ret == hope, n, k, hope, ret)

}