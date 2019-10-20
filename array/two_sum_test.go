package array

import (
	"testing"
)

func Test_twoSum(t *testing.T) {
	var nums, h, r []int
	var target int

	nums = []int{2, 7, 11, 15}
	target = 9
	h = []int{0, 1}
	r = twoSum(nums, target)
	t.Logf("nums=%v h=%d r=%d", nums, h, r)
}
