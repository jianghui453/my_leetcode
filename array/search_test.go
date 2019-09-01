package array

import "testing"

func TestSearch(t *testing.T) {
	var nums []int
	var hope, ret bool
	var target int

	nums = []int{2, 5, 6, 0, 0, 1, 2}
	hope = true
	target = 0
	ret = search(nums, target)
	t.Logf("nums=%v target=%d hope=%t ret=%t\n", nums, target, hope, ret)

	nums = []int{2, 5, 6, 0, 0, 1, 2}
	hope = false
	target = 3
	ret = search(nums, target)
	t.Logf("nums=%v target=%d hope=%t ret=%t\n", nums, target, hope, ret)

	nums = []int{1, 3}
	hope = false
	target = 2
	ret = search(nums, target)
	t.Logf("nums=%v target=%d hope=%t ret=%t\n", nums, target, hope, ret)

	nums = []int{1, 3, 1, 1, 1}
	hope = true
	target = 3
	ret = search(nums, target)
	t.Logf("nums=%v target=%d hope=%t ret=%t\n", nums, target, hope, ret)

	nums = []int{0}
	hope = false
	target = 1
	ret = search(nums, target)
	t.Logf("nums=%v target=%d hope=%t ret=%t\n", nums, target, hope, ret)

	nums = []int{1, 3}
	hope = true
	target = 3
	ret = search(nums, target)
	t.Logf("nums=%v target=%d hope=%t ret=%t\n", nums, target, hope, ret)

	nums = []int{1, 1, 3}
	hope = true
	target = 3
	ret = search(nums, target)
	t.Logf("nums=%v target=%d hope=%t ret=%t\n", nums, target, hope, ret)

	nums = []int{3, 1}
	hope = false
	target = 2
	ret = search(nums, target)
	t.Logf("nums=%v target=%d hope=%t ret=%t\n", nums, target, hope, ret)
}
