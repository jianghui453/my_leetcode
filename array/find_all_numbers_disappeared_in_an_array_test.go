package array

import (
	"testing"
)

func Test_findDisappearedNumbers(t *testing.T) {
	var (
		nums, hope, ret []int
	)

	nums = []int{4, 3, 2, 7, 8, 2, 3, 1}
	hope = []int{5, 6}
	ret = findDisappearedNumbers(nums)
	t.Log(nums)
	t.Log(hope)
	t.Log(ret)

	nums = []int{1, 1, 1, 1, 1, 1, 1}
	hope = []int{2, 3, 4, 5, 6, 7}
	ret = findDisappearedNumbers(nums)
	t.Log(nums)
	t.Log(hope)
	t.Log(ret)

	nums = []int{1}
	hope = []int{}
	ret = findDisappearedNumbers(nums)
	t.Log(nums)
	t.Log(hope)
	t.Log(ret)
}
