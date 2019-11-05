package backtracking

import "testing"

func TestPermuteUnique(t *testing.T) {
	var nums []int
	var h, r [][]int

	nums = []int{1, 1, 2}
	h = [][]int{
	  	{1,1,2},
		{1,2,1},
		{2,1,1},
	}
	r = permuteUnique(nums)
	t.Logf("nums=%v h=%v r= %v\n", nums, h, r)
}
