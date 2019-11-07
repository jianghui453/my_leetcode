package dynamic_programming

import "testing"

func TestMaxSubArray(t *testing.T) {
	var nums []int
	var h, r int

	nums = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	h = 6
	r = maxSubArray(nums)
	t.Logf("%t nums=%v h=%d r=%d", r==h, nums, h, r)

	nums = []int{-1, 0, -2}
	h = 0
	r = maxSubArray(nums)
	t.Logf("%t nums=%v h=%d r=%d", r==h, nums, h, r)

	nums = []int{-1, 1, 2, 1}
	h = 4
	r = maxSubArray(nums)
	t.Logf("%t nums=%v h=%d r=%d", r==h, nums, h, r)

	nums = []int{-1}
	h = -1
	r = maxSubArray(nums)
	t.Logf("%t nums=%v h=%d r=%d", r==h, nums, h, r)
}
