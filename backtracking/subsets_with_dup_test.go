package backtracking

import "testing"

func TestSubsetsWithDup(t *testing.T) {
	var nums []int
	var hope, ret [][]int

	nums = []int{1, 2, 2}
	hope = [][]int{
		{2},
		{1},
		{1, 2, 2},
		{2, 2},
		{1, 2},
		{},
	}
	ret = subsetsWithDup(nums)
	t.Logf("nums=%v \nhope=%v \n ret=%v", nums, hope, ret)
}
