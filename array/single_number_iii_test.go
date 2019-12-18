package array

import (
	"testing"
)

func Test_singleNumber(t *testing.T) {
	var nums, hope, ret []int

	nums = []int{1,2,1,3,2,5}
	hope = []int{3, 5}
	ret = singleNumber(nums)
	t.Log(nums)
	t.Log("hope =", hope)
	t.Log(" ret =", ret)

	nums = []int{1,2,1,3,2,5, 0, 0}
	hope = []int{3, 5}
	ret = singleNumber(nums)
	t.Log(nums)
	t.Log("hope =", hope)
	t.Log(" ret =", ret)

	nums = []int{3,5}
	hope = []int{3, 5}
	ret = singleNumber(nums)
	t.Log(nums)
	t.Log("hope =", hope)
	t.Log(" ret =", ret)

	nums = []int{1, 1,3,5}
	hope = []int{3, 5}
	ret = singleNumber(nums)
	t.Log(nums)
	t.Log("hope =", hope)
	t.Log(" ret =", ret)
}