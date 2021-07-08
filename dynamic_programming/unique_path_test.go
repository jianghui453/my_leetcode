package dynamic_programming

import "testing"

func TestUniquePath(t *testing.T) {
	var m, n, h, r int

	m = 3
	n = 2
	h = 3
	r = uniquePaths(m, n)
	t.Logf("m=%d n=%d h=%d r=%d", m, n, h, r)

	m = 7
	n = 3
	h = 28
	r = uniquePaths(m, n)
	t.Logf("m=%d n=%d h=%d r=%d", m, n, h, r)
}
