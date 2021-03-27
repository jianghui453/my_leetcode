package array

import "testing"

func Test_majorityElement(t *testing.T) {
	var nums []int
	var hope, ret int

	nums = []int{2, 2, 1, 1, 1, 2, 2}
	hope = 2
	ret = majorityElement(nums)
	t.Log(ret == hope, nums, hope, ret)
}
