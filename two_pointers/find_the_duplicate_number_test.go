package two_pointers

import (
	"testing"
)

func Test_findDuplicate(t *testing.T) {
	var nums []int
	var hope, ret int

	nums = []int{1, 1}
	hope = 1
	ret = findDuplicate(nums)
	t.Log(ret == hope, nums, hope, ret)

	nums = []int{3,1,2,2}
	hope = 2
	ret = findDuplicate(nums)
	t.Log(ret == hope, nums, hope, ret)

	nums = []int{1,3,4,2,2}
	hope = 2
	ret = findDuplicate(nums)
	t.Log(ret == hope, nums, hope, ret)

	nums = []int{3,1,3,4,2}
	hope = 3
	ret = findDuplicate(nums)
	t.Log(ret == hope, nums, hope, ret)

	nums = []int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1}
	hope = 9
	ret = findDuplicate(nums)
	t.Log(ret == hope, nums, hope, ret)
}
