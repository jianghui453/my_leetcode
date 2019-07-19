package jump

import "testing"

func TestJump(t *testing.T) {
	var nums []int
	var ret, hope int

	nums = []int{1, 2, 3}
	ret = jump(nums)
	hope = 2
	t.Logf("nums = %v, ret = %d, hope = %d.\n", nums, ret, hope)

	nums = []int{2,3,1,1,4}
	ret = jump(nums)
	hope = 2
	t.Logf("nums = %v, ret = %d, hope = %d.\n", nums, ret, hope)
}