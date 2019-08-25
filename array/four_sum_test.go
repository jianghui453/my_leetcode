package array

import (
	"testing"
)

func TestFourSum(t *testing.T) {
	var nums []int
	var target int
	var hope, ret [][]int

	nums = []int{1, 0, -1, 0, -2, 2}
	target = 0
	hope = [][]int{
		[]int{-1, 0, 0, 1},
		[]int{-2, -1, 1, 2},
		[]int{-2, 0, 0, 2},
	}
	ret = fourSum(nums, target)
	t.Logf("nums = %v; target = %d; hope = %v; ret = %v", nums, target, hope, ret)
}
