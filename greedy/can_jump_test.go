package greedy

import "testing"

func TestCanJump(t *testing.T) {
	var nums []int
	var h, r bool

	nums = []int{2, 3, 1, 1, 4}
	h = true
	r = canJump(nums)
	t.Logf("%t nums=%v h=%v r=%v", r==h, nums, h, r)

	nums = []int{3, 2, 1, 0, 4}
	h = false
	r = canJump(nums)
	t.Logf("%t nums=%v h=%v r=%v", r==h, nums, h, r)

	nums = []int{3, 2, 2, 0, 0, 0}
	h = false
	r = canJump(nums)
	t.Logf("%t nums=%v h=%v r=%v", r==h, nums, h, r)
}
