package binary_search

import "testing"

func TestSearch(t *testing.T) {
	var nums []int
	var h, r, target int

	nums = []int{5, 6, 0, 1, 2}
	target = 0
	h = 2
	r = search(nums, target)
	t.Logf("%t nums=%v target=%d h=%d r=%d\n", r==h, nums, target, h, r)

	nums = []int{5, 6, 0, 1, 2}
	target = 3
	h = -1
	r = search(nums, target)
	t.Logf("%t nums=%v target=%d h=%d r=%d\n", r==h, nums, target, h, r)

	nums = []int{1, 3}
	target = 2
	h = -1
	r = search(nums, target)
	t.Logf("%t nums=%v target=%d h=%d r=%d\n", r==h, nums, target, h, r)

	nums = []int{1, 3}
	target = 1
	h = 0
	r = search(nums, target)
	t.Logf("%t nums=%v target=%d h=%d r=%d\n", r==h, nums, target, h, r)

	nums = []int{0}
	target = 1
	h = -1
	r = search(nums, target)
	t.Logf("%t nums=%v target=%d h=%d r=%d\n", r==h, nums, target, h, r)

	nums = []int{1, 3}
	target = 3
	h = 1
	r = search(nums, target)
	t.Logf("%t nums=%v target=%d h=%d r=%d\n", r==h, nums, target, h, r)

	nums = []int{1, 2, 3}
	target = 3
	h = 2
	r = search(nums, target)
	t.Logf("%t nums=%v target=%d h=%d r=%d\n", r==h, nums, target, h, r)

	nums = []int{3, 1}
	target = 2
	h = -1
	r = search(nums, target)
	t.Logf("%t nums=%v target=%d h=%d r=%d\n", r==h, nums, target, h, r)

	nums = []int{4,5,6,7,0,1,2}
	target = 0
	h = 4
	r = search(nums, target)
	t.Logf("%t nums=%v target=%d h=%d r=%d\n", r==h, nums, target, h, r)

	nums = []int{278,280,281,286,287,290,2,3,4,8,9,14,15,16,21,24,25,31,32,34,36,37,42,45,51,52,54,55,60,63,66,68,69,71,76,81,83,84,85,86,87,94,97,99,106,107,110,113,114,115,118,120,121,125,134,136,137,138,142,143,147,150,152,159,160,161,165,166,174,176,178,186,187,189,190,191,195,196,198,204,212,216,217,220,221,222,225,227,229,232,237,239,242,245,251,263,264,274,275,276,277}
	target = 286
	h = 3
	r = search(nums, target)
	t.Logf("%t nums=%v target=%d h=%d r=%d\n", r==h, nums, target, h, r)
}
