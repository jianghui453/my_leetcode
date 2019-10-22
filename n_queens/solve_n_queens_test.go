package solve_n_queens

import "testing"

func TestSolveNQueens(t *testing.T) {
	var n int
	var rets [][]string

	n = 4
	rets = solveNQueens(n)
	t.Logf("n=%d rets=%v", n, rets)
}
