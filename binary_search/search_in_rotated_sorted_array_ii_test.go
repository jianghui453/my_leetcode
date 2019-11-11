package binary_search

import "testing"

func TestSearch(t *testing.T) {
	var nums []int
	var h, r bool
	var target int

	nums = []int{2, 5, 6, 0, 0, 1, 2}
	h = true
	target = 0
	r = search(nums, target)
	t.Logf("%t nums=%v target=%d h=%t r=%t\n", r==h, nums, target, h, r)

	nums = []int{2, 5, 6, 0, 0, 1, 2}
	h = false
	target = 3
	r = search(nums, target)
	t.Logf("%t nums=%v target=%d h=%t r=%t\n", r==h, nums, target, h, r)

	nums = []int{1, 3}
	h = false
	target = 2
	r = search(nums, target)
	t.Logf("%t nums=%v target=%d h=%t r=%t\n", r==h, nums, target, h, r)

	nums = []int{1, 3, 1, 1, 1}
	h = true
	target = 3
	r = search(nums, target)
	t.Logf("%t nums=%v target=%d h=%t r=%t\n", r==h, nums, target, h, r)

	nums = []int{0}
	h = false
	target = 1
	r = search(nums, target)
	t.Logf("%t nums=%v target=%d h=%t r=%t\n", r==h, nums, target, h, r)

	nums = []int{1, 3}
	h = true
	target = 3
	r = search(nums, target)
	t.Logf("%t nums=%v target=%d h=%t r=%t\n", r==h, nums, target, h, r)

	nums = []int{1, 1, 3}
	h = true
	target = 3
	r = search(nums, target)
	t.Logf("%t nums=%v target=%d h=%t r=%t\n", r==h, nums, target, h, r)

	nums = []int{3, 1}
	h = false
	target = 2
	r = search(nums, target)
	t.Logf("%t nums=%v target=%d h=%t r=%t\n", r==h, nums, target, h, r)

	nums = []int{1}
	h = true
	target = 1
	r = search(nums, target)
	t.Logf("%t nums=%v target=%d h=%t r=%t\n", r==h, nums, target, h, r)
}
