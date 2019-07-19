package permute_unique

import "testing"

func TestPermuteUnique(t *testing.T) {
	var nums []int
	var hope, rets [][]int

	nums = []int{1,1,2}
	hope = [][]int{
		{1,2,3},
	}
	rets = permuteUnique(nums)
	t.Logf("nums = %v, hope = %v, rets = %v.\n", nums, hope, rets)
}