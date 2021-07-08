package array

import (
	"testing"
)

func TestSubarraySum(t *testing.T) {
	var (
		nums []int
		k    int
		hope int
		ret  int
	)

	nums = []int{1, 1, 1}
	k = 2
	hope = 2
	ret = subarraySum(nums, k)
	t.Log(ret == hope, nums, k, hope, ret)

	nums = []int{1, 1, 1}
	k = 1
	hope = 3
	ret = subarraySum(nums, k)
	t.Log(ret == hope, nums, k, hope, ret)

	nums = []int{1, 1, 1}
	k = 3
	hope = 1
	ret = subarraySum(nums, k)
	t.Log(ret == hope, nums, k, hope, ret)

	nums = []int{1, 2, 3}
	k = 3
	hope = 2
	ret = subarraySum(nums, k)
	t.Log(ret == hope, nums, k, hope, ret)

	nums = []int{-1, -1, 1, 2}
	k = 1
	hope = 2
	ret = subarraySum(nums, k)
	t.Log(ret == hope, nums, k, hope, ret)
}
