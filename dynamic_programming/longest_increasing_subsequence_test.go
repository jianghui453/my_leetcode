package dynamic_programming

import (
	"testing"
)

func Test_lengthOfLIS(t *testing.T) {
	var nums []int
	var hope, ret int

	nums = []int{10,9,2,5,3,7,101,18}
	hope = 4
	ret = lengthOfLIS(nums)
	t.Logf("%t nums=%v hope=%d ret=%d", ret == hope, nums, hope, ret)

	nums = []int{10,9,2,4,5,3,7,101,18}
	hope = 5
	ret = lengthOfLIS(nums)
	t.Logf("%t nums=%v hope=%d ret=%d", ret == hope, nums, hope, ret)

	
	nums = []int{0}
	hope = 1
	ret = lengthOfLIS(nums)
	t.Logf("%t nums=%v hope=%d ret=%d", ret == hope, nums, hope, ret)

	nums = []int{2, 2}
	hope = 1
	ret = lengthOfLIS(nums)
	t.Logf("%t nums=%v hope=%d ret=%d", ret == hope, nums, hope, ret)
}