package backtracking

import "testing"

func TestCombine(t *testing.T) {
	var n, k int
	var hope, ret [][]int

	n = 4
	k = 2
	hope = [][]int{
		{2, 4},
		{3, 4},
		{2, 3},
		{1, 2},
		{1, 3},
		{1, 4},
	}
	ret = combine(n, k)
	t.Logf("n=%d k=%d \nhope=%v \n ret=%v", n, k, hope, ret)
}
