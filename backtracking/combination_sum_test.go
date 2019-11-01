package backtracking

import "testing"

func TestCombinationSum(t *testing.T) {
	var (
		candidates []int
		target     int
		h          [][]int
		r          [][]int
	)

	candidates = []int{2, 3, 6, 7}
	target = 7
	h = [][]int{
		{7},
		{2, 2, 3},
	}
	r = combinationSum(candidates, target)
	t.Logf("candidates=%v \ntarget=%d \nh=%v \nr=%v", candidates, target, h, r)

	candidates = []int{7, 3, 2}
	target = 18
	h = [][]int{
		{2, 2, 2, 2, 2, 2, 2, 2, 2},
		{2, 2, 2, 2, 2, 2, 3, 3},
		{2, 2, 2, 2, 7, 3},
		{2, 2, 7, 7},
		{2, 2, 2, 3, 3, 3, 3},
		{7, 2, 3, 3, 3},
		{3, 3, 3, 3, 3, 3},
	}
	r = combinationSum(candidates, target)
	t.Logf("candidates=%v \ntarget=%d \nh=%v \nr=%v", candidates, target, h, r)
}
