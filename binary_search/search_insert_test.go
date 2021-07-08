package binary_search

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	var nums []int
	var target, h, r int

	nums = []int{1}
	target = 1
	h = 0
	r = searchInsert(nums, target)
	t.Logf("nums=%v target=%d h=%d r=%d", nums, target, h, r)

	nums = []int{1}
	target = 0
	h = 0
	r = searchInsert(nums, target)
	t.Logf("nums=%v target=%d h=%d r=%d", nums, target, h, r)

	nums = []int{1}
	target = 0
	h = 0
	r = searchInsert(nums, target)
	t.Logf("nums=%v target=%d h=%d r=%d", nums, target, h, r)

	nums = []int{1}
	target = 2
	h = 1
	r = searchInsert(nums, target)
	t.Logf("nums=%v target=%d h=%d r=%d", nums, target, h, r)
	nums = []int{1, 2}
	target = 2
	h = 1
	r = searchInsert(nums, target)
	t.Logf("nums=%v target=%d h=%d r=%d", nums, target, h, r)

	nums = []int{1, 2}
	target = 1
	h = 0
	r = searchInsert(nums, target)
	t.Logf("nums=%v target=%d h=%d r=%d", nums, target, h, r)

	nums = []int{1, 2}
	target = 3
	h = 2
	r = searchInsert(nums, target)
	t.Logf("nums=%v target=%d h=%d r=%d", nums, target, h, r)

	nums = []int{1, 2}
	target = 0
	h = 0
	r = searchInsert(nums, target)
	t.Logf("nums=%v target=%d h=%d r=%d", nums, target, h, r)

	nums = []int{1, 3, 5, 6}
	target = 2
	h = 1
	r = searchInsert(nums, target)
	t.Logf("nums=%v target=%d h=%d r=%d", nums, target, h, r)

	nums = []int{1, 3}
	target = 2
	h = 1
	r = searchInsert(nums, target)
	t.Logf("nums=%v target=%d h=%d r=%d", nums, target, h, r)
}
