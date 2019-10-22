package array

import "testing"

func Test_singleNumberMine(t *testing.T) {
	var nums []int
	var h, r int

	// singleNumber
	//nums = []int{1, 2, 3, 3, 2}
	//h = 1
	//r = singleNumber(nums)
	//t.Logf("nums=%v h=%d r=%d", nums, h, r)

	// singleNumberII
	nums = []int{2, 2, 3, 2}
	h = 3
	r = singleNumber(nums)
	t.Logf("nums=%v h=%d r=%d", nums, h, r)

	nums = []int{0, 1, 0, 1, 0, 1, 99}
	h = 99
	r = singleNumber(nums)
	t.Logf("nums=%v h=%d r=%d", nums, h, r)

	nums = []int{-2, -2, 1, 1, -3, 1, -3, -3, -4, -2}
	h = -4
	r = singleNumber(nums)
	t.Logf("nums=%v h=%d r=%d", nums, h, r)
}
