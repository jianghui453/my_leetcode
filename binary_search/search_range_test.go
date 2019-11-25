package binary_search

import (
	"testing"
)

func TestSearchRange(t *testing.T) {
	var nums, h, r []int
	var target int

	nums = []int{5, 7, 7, 8, 8, 10}
	target = 8
	h = []int{3, 4}
	r = searchRange(nums, target)
	t.Logf("nums=%v target=%d h=%v r=%v", nums, target, h, r)

	nums = []int{5, 7, 7, 8, 8, 10}
	target = 6
	h = []int{-1, -1}
	r = searchRange(nums, target)
	t.Logf("nums=%v  target=%d  h=%v  r=%v", nums, target, h, r)

	nums = []int{5, 7, 7, 8, 8, 10}
	target = 10
	h = []int{5, 5}
	r = searchRange(nums, target)
	t.Logf("nums=%v  target=%d  h=%v  r=%v", nums, target, h, r)

	nums = []int{5, 7, 7, 8, 8, 10}
	target = 5
	h = []int{0, 0}
	r = searchRange(nums, target)
	t.Logf("nums=%v  target=%d  h=%v  r=%v", nums, target, h, r)

	nums = []int{0, 0}
	target = 1
	h = []int{-1, -1}
	r = searchRange(nums, target)
	t.Logf("nums=%v  target=%d  h=%v  r=%v", nums, target, h, r)

	nums = []int{0, 0}
	target = 0
	h = []int{0, 1}
	r = searchRange(nums, target)
	t.Logf("nums=%v  target=%d  h=%v  r=%v", nums, target, h, r)

	nums = []int{0}
	target = 1
	h = []int{-1, -1}
	r = searchRange(nums, target)
	t.Logf("nums=%v  target=%d  h=%v  r=%v", nums, target, h, r)

	nums = []int{0}
	target = 0
	h = []int{0, 0}
	r = searchRange(nums, target)
	t.Logf("nums=%v  target=%d  h=%v  r=%v", nums, target, h, r)
}
