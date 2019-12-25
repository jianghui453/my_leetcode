package array

import "testing"

func Test_singleNumberMine(t *testing.T) {
	var nums []int
	var h, r int

	nums = []int{2, 2, 3, 2}
	h = 3
	r = singleNumber(nums)
	t.Logf("%t nums=%v h=%d r=%d", r == h, nums, h, r)

	nums = []int{0, 1, 0, 1, 0, 1, 99}
	h = 99
	r = singleNumber(nums)
	t.Logf("%t nums=%v h=%d r=%d", r == h, nums, h, r)

	nums = []int{-2, -2, 1, 1, -3, 1, -3, -3, -4, -2}
	h = -4
	r = singleNumber(nums)
	t.Logf("%t nums=%v h=%d r=%d", r == h, nums, h, r)
}
