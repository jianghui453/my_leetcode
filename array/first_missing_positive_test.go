package array

import (
	"testing"
)

func TestFirstMissingPositive(t *testing.T) {
	var nums []int
	var h, r int

	nums = []int{1, 2, 0}
	h = 3
	r = firstMissingPositive(nums)
	t.Logf("nums=%v h=%d r=%d\n", nums, h, r)

	nums = []int{3, 4, -1, 1}
	h = 2
	r = firstMissingPositive(nums)
	t.Logf("nums=%v h=%d r=%d\n", nums, h, r)

	nums = []int{7, 8, 9, 11, 12}
	h = 1
	r = firstMissingPositive(nums)
	t.Logf("nums=%v h=%d r=%d\n", nums, h, r)

	nums = []int{}
	h = 1
	r = firstMissingPositive(nums)
	t.Logf("nums=%v h=%d r=%d\n", nums, h, r)
}
