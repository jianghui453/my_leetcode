package max_sub_array

import "testing"

func TestMaxSubArray(t *testing.T) {
	var nums []int
	var hope, ret int

	nums = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	hope = 6
	ret = maxSubArray(nums)
	t.Logf("nums=%v hope=%d ret=%d", nums, hope, ret)

	nums = []int{-1, 0, -2}
	hope = 0
	ret = maxSubArray(nums)
	t.Logf("nums=%v hope=%d ret=%d", nums, hope, ret)

	nums = []int{-1, 1, 2, 1}
	hope = 4
	ret = maxSubArray(nums)
	t.Logf("nums=%v hope=%d ret=%d", nums, hope, ret)
}
