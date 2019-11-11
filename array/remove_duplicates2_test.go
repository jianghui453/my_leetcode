package array

import "testing"

func TestRemoveDuplicates2(t *testing.T) {
	var nums []int
	var hope, ret int

	nums = []int{1, 1, 1, 2, 2, 3}
	hope = 5
	ret = removeDuplicates(nums)
	t.Logf("nums=%v hope=%d ret=%d", nums, hope, ret)

	nums = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	hope = 7
	ret = removeDuplicates(nums)
	t.Logf("nums=%v hope=%d ret=%d", nums, hope, ret)

	nums = []int{1}
	hope = 1
	ret = removeDuplicates(nums)
	t.Logf("nums=%v hope=%d ret=%d", nums, hope, ret)

	nums = []int{}
	hope = 0
	ret = removeDuplicates(nums)
	t.Logf("nums=%v hope=%d ret=%d", nums, hope, ret)

	nums = []int{1, 2}
	hope = 2
	ret = removeDuplicates(nums)
	t.Logf("nums=%v hope=%d ret=%d", nums, hope, ret)
}
