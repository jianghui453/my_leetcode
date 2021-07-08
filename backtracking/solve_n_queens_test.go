package backtracking

import "testing"

func TestSolveNQueens(t *testing.T) {
	var n int
	var h, r [][]string

	n = 4
	h = [][]string{
		{".Q..", "...Q", "Q...", "..Q."},
		{"..Q.", "Q...", "...Q", ".Q.."},
	}
	r = solveNQueens(n)
	t.Logf("n=%d \nh=%v \nr=%v", n, h, r)

	n = 1
	h = [][]string{
		{"Q"},
	}
	r = solveNQueens(n)
	t.Logf("n=%d \nh=%v \nr=%v", n, h, r)
}
