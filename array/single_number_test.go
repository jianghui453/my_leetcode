package array

import "testing"

func Test_singleNumberMine(t *testing.T) {
	var nums []int
	var h, r int

	// singleNumber
	nums = []int{1, 2, 3, 3, 2}
	h = 1
	r = singleNumber(nums)
	t.Logf("%t nums=%v h=%d r=%d", r == h, nums, h, r)
}
