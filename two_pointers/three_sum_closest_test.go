package two_pointers

import (
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	var nums []int
	var target, h, r int

	// nums = []int{-1, 2, 1, -4}
	// target = 1
	// h = 2
	// r = threeSumClosest(nums, target)
	// t.Logf("%t nums=%v target=%d h=%d r=%d", r==h, nums, target, h, r)

	nums = []int{-3, -2, -5, 3, -4}
	target = -1
	h = -2
	r = threeSumClosest(nums, target)
	t.Logf("%t nums=%v target=%d h=%d r=%d", r == h, nums, target, h, r)
}
