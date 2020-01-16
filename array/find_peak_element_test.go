package array

import (
	"testing"
)

func TestFindPeakElement(t *testing.T) {
	var (
		nums      []int
		hope, ret int
	)

	nums = []int{1, 2, 3, 1}
	hope = 2
	ret = findPeakElement(nums)
	t.Log(ret == hope, nums, hope, ret)

	nums = []int{1, 2, 1, 3, 5, 6, 4}
	hope = 1
	ret = findPeakElement(nums)
	t.Log(ret == hope, nums, hope, ret)

	nums = []int{1, 2, 3, 4, 5, 6}
	hope = 5
	ret = findPeakElement(nums)
	t.Log(ret == hope, nums, hope, ret)

	nums = []int{3, 2, 1}
	hope = 0
	ret = findPeakElement(nums)
	t.Log(ret == hope, nums, hope, ret)
}
