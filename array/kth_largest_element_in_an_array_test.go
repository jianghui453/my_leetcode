package array

import (
	"testing"
)

func Test_findKthLargest(t *testing.T) {
	var nums []int
	var k, hope, ret int

	nums = []int{3, 2, 1, 5, 6, 4}
	hope = 5
	k = 2
	ret = findKthLargest(nums, k)
	t.Logf("%t nums=%v k=%d hope=%d ret=%d", ret == hope, nums, k, hope, ret)

	nums = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	hope = 4
	k = 4
	ret = findKthLargest(nums, k)
	t.Logf("%t nums=%v k=%d hope=%d ret=%d", ret == hope, nums, k, hope, ret)

	nums = []int{1}
	hope = 1
	k = 1
	ret = findKthLargest(nums, k)
	t.Logf("%t nums=%v k=%d hope=%d ret=%d", ret == hope, nums, k, hope, ret)
}
