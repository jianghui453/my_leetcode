package array

import (
	"testing"
)

func Test_rotate(t *testing.T) {
	var (
		nums, hope []int
		k          int
	)

	nums = []int{1, 2, 3, 4, 5, 6, 7}
	k = 3
	hope = []int{5, 6, 7, 1, 2, 3, 4}
	t.Log("before nums =", nums)
	rotate(nums, k)
	t.Log(nums)
	t.Log(hope)

	nums = []int{1, 2, 3, 4, 5, 6, 7}
	k = 1
	hope = []int{7, 1, 2, 3, 4, 5, 6}
	t.Log("before nums =", nums)
	rotate(nums, k)
	t.Log(nums)
	t.Log(hope)

	nums = []int{1, 2, 3, 4, 5, 6, 7}
	k = 7
	hope = []int{1, 2, 3, 4, 5, 6, 7}
	t.Log("before nums =", nums)
	rotate(nums, k)
	t.Log(nums)
	t.Log(hope)

	nums = []int{1, 2, 3, 4, 5, 6, 7}
	k = 6
	hope = []int{2, 3, 4, 5, 6, 7, 1}
	t.Log("before nums =", nums)
	rotate(nums, k)
	t.Log(nums)
	t.Log(hope)
}
