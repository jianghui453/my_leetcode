package backtracking

import "testing"

func TestTotalNQueens(t *testing.T) {
	var n, h, r int

	n = 4
	h = 2
	r = totalNQueens(n)
	t.Logf("%t n=%d \nh=%d \nr=%d", r == h, n, h, r)

	n = 1
	h = 1
	r = totalNQueens(n)
	t.Logf("%t n=%d \nh=%d \nr=%d", r == h, n, h, r)
}
