package backtracking

import "testing"

func TestPermute(t *testing.T) {
	var nums []int
	var h, r [][]int

	nums = []int{1, 2, 3}
	h = [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}
	r = permute(nums)
	t.Logf("nums=%v \nh=%v \nr=%v\n", nums, h, r)
}
