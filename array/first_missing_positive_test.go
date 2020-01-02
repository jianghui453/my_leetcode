package array

import (
	"testing"
)

func TestFirstMissingPositive(t *testing.T) {
	var nums []int
	var h, r int

	nums = []int{1, 2, 0}
	h = 3
	t.Log(nums)
	r = firstMissingPositive(nums)
	t.Logf("%t nums=%v h=%d r=%d\n", r == h, nums, h, r)

	nums = []int{3, 4, -1, 1}
	h = 2
	t.Log(nums)
	r = firstMissingPositive(nums)
	t.Logf("%t nums=%v h=%d r=%d\n", r == h, nums, h, r)

	nums = []int{7, 8, 9, 11, 12}
	h = 1
	t.Log(nums)
	r = firstMissingPositive(nums)
	t.Logf("%t nums=%v h=%d r=%d\n", r == h, nums, h, r)

	nums = []int{2}
	h = 1
	t.Log(nums)
	r = firstMissingPositive(nums)
	t.Logf("%t nums=%v h=%d r=%d\n", r == h, nums, h, r)

	nums = []int{1}
	h = 2
	t.Log(nums)
	r = firstMissingPositive(nums)
	t.Logf("%t nums=%v h=%d r=%d\n", r == h, nums, h, r)

	nums = []int{1, 2}
	h = 3
	t.Log(nums)
	r = firstMissingPositive(nums)
	t.Logf("%t nums=%v h=%d r=%d\n", r == h, nums, h, r)

	nums = []int{}
	h = 1
	t.Log(nums)
	r = firstMissingPositive(nums)
	t.Logf("%t nums=%v h=%d r=%d\n", r == h, nums, h, r)
}
