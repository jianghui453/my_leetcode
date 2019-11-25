package two_pointers

import (
	"testing"
)

func TestFourSum(t *testing.T) {
	var nums []int
	var target int
	var h, r [][]int

	nums = []int{1, 0, -1, 0, -2, 2}
	target = 0
	h = [][]int{{-1, 0, 0, 1}, {-2, -1, 1, 2}, {-2, 0, 0, 2}}
	r = fourSum(nums, target)
	t.Logf("nums=%v target=%d h=%v r=%v", nums, target, h, r)

	nums = []int{5, 5, 3, 5, 1, -5, 1, -2}
	target = 4
	h = [][]int{{-5, 1, 3, 5}}
	r = fourSum(nums, target)
	t.Logf("nums=%v target=%d h=%v r=%v", nums, target, h, r)

	nums = []int{0, 0, 0, 0}
	target = 0
	h = [][]int{{0, 0, 0, 0}}
	r = fourSum(nums, target)
	t.Logf("nums=%v target=%d h=%v r=%v", nums, target, h, r)

	nums = []int{0, 0, 4, -2, -3, -2, -2, -3}
	target = -1
	h = [][]int{{-3, -2, 0, 4}}
	r = fourSum(nums, target)
	t.Logf("nums=%v target=%d h=%v r=%v", nums, target, h, r)

	nums = []int{1, -2, -5, -4, -3, 3, 3, 5}
	target = -11
	h = [][]int{{-5, -4, -3, 1}}
	r = fourSum(nums, target)
	t.Logf("nums=%v target=%d h=%v r=%v", nums, target, h, r)
}
