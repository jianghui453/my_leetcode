package permute

import "testing"

func TestPermute(t *testing.T) {
	var nums []int
	var hope, rets [][]int

	nums = []int{1,2,3}
	hope = [][]int{
		{1,2,3},
	}
	rets = permute(nums)
	t.Logf("nums = %v, hope = %v, rets = %v.\n", nums, hope, rets)
}