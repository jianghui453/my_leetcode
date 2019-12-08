package dynamic_programming

import (
	"testing"
)

func Test_rob(t *testing.T) {
	var nums []int
	var hope, ret int

	nums = []int{1, 2, 3, 1}
	hope = 4
	ret = rob(nums)
	t.Logf("%t nums=%v hope=%d ret=%d", ret == hope, nums, hope, ret)

	nums = []int{2, 7, 9, 3, 1}
	hope = 12
	ret = rob(nums)
	t.Logf("%t nums=%v hope=%d ret=%d", ret == hope, nums, hope, ret)
}