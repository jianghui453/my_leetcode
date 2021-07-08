package array

import (
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	var (
		nums, hope, ret []int
	)

	nums = []int{4,3,2,7,8,2,3,1}
	hope = []int{2, 3}
	ret = findDuplicates(nums)
	t.Log(nums)
	t.Log(hope)
	t.Log(ret)
	t.Log()

	nums = []int{1}
	hope = []int{}
	ret = findDuplicates(nums)
	t.Log(nums)
	t.Log(hope)
	t.Log(ret)
	t.Log()

	nums = []int{1, 1}
	hope = []int{1}
	ret = findDuplicates(nums)
	t.Log(nums)
	t.Log(hope)
	t.Log(ret)
	t.Log()

	nums = []int{2, 2}
	hope = []int{2}
	ret = findDuplicates(nums)
	t.Log(nums)
	t.Log(hope)
	t.Log(ret)
	t.Log()

	nums = []int{4,3,2,7,8,2,3,4}
	hope = []int{2, 3,4}
	ret = findDuplicates(nums)
	t.Log(nums)
	t.Log(hope)
	t.Log(ret)
	t.Log()

	nums = []int{}
	hope = []int{}
	ret = findDuplicates(nums)
	t.Log(nums)
	t.Log(hope)
	t.Log(ret)
	t.Log()

	nums = []int{10,2,5,10,9,1,1,4,3,7}
	hope = []int{10, 1}
	ret = findDuplicates(nums)
	t.Log(nums)
	t.Log(hope)
	t.Log(ret)
	t.Log()
}