package jump

import "testing"

func TestCanJump(t *testing.T) {
	var nums []int
	var hope, ret bool

	nums = []int{2, 3, 1, 1, 4}
	hope = true
	ret = canJump(nums)
	t.Logf("nums=%v hope=%v ret=%v", nums, hope, ret)

	nums = []int{3, 2, 1, 0, 4}
	hope = false
	ret = canJump(nums)
	t.Logf("nums=%v hope=%v ret=%v", nums, hope, ret)
}
