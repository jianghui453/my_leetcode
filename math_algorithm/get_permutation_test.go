package math_algorithm

import "testing"

func TestGetPermutation(t *testing.T) {
	var n, k int
	var h, r string

	n = 2
	k = 2
	h = "21"
	r = getPermutation(n, k)
	t.Logf("%t n=%d k=%d h=%s r=%s", r == h, n, k, h, r)

	n = 3
	k = 7
	h = "123"
	r = getPermutation(n, k)
	t.Logf("%t n=%d k=%d h=%s r=%s", r == h, n, k, h, r)

	n = 3
	k = 3
	h = "213"
	r = getPermutation(n, k)
	t.Logf("%t n=%d k=%d h=%s r=%s", r == h, n, k, h, r)

	n = 3
	k = 5
	h = "312"
	r = getPermutation(n, k)
	t.Logf("%t n=%d k=%d h=%s r=%s", r == h, n, k, h, r)

	n = 4
	k = 9
	h = "2314"
	r = getPermutation(n, k)
	t.Logf("%t n=%d k=%d h=%s r=%s", r == h, n, k, h, r)
}
