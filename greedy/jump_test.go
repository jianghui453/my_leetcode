package greedy

import "testing"

func TestJump(t *testing.T) {
	var nums []int
	var r, h int

	nums = []int{1, 2, 3}
	r = jump(nums)
	h = 2
	t.Logf("nums=%v r=%d h=%d\n", nums, r, h)

	nums = []int{2, 3, 1, 1, 4}
	r = jump(nums)
	h = 2
	t.Logf("nums=%v r=%d h=%d\n", nums, r, h)

	nums = []int{1, 1}
	r = jump(nums)
	h = 1
	t.Logf("nums=%v r=%d h=%d\n", nums, r, h)

	nums = []int{1, 1, 1}
	r = jump(nums)
	h = 2
	t.Logf("nums=%v r=%d h=%d\n", nums, r, h)

	nums = []int{1, 1, 1, 1}
	r = jump(nums)
	h = 3
	t.Logf("nums=%v r=%d h=%d\n", nums, r, h)
}
