package leet_code

import (
	"testing"
)

func TestThreeSum(t *testing.T) {
	var nums []int
	var hope, ret [][]int
	// var ok bool

	nums = []int{
		-1, 0, 1, 2, -1, -4,
	}
	ret = threeSum(nums)
	hope = [][]int{
		{
			-1, 0, 1,
		},
		{
			-1, -1, 2,
		},
	}
	t.Logf("nums = %v; hope = %v; ret = %v \n", nums, hope, ret)

	nums = []int{
		0, 0, 0, 0, 0,
	}
	ret = threeSum(nums)
	hope = [][]int{
		{
			0, 0, 0,
		},
	}
	t.Logf("nums = %v; hope = %v; ret = %v \n", nums, hope, ret)

	nums = []int{
		-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0,
	}
	ret = threeSum(nums)
	hope = [][]int{
		{
			-5, 1, 4,
		},
		{
			-4, 0, 4,
		},
		{
			-4, 1, 3,
		},
		{
			-2, -2, 4,
		},
		{
			-2, 1, 1,
		},
		{
			0, 0, 0,
		},
	}
	t.Logf("nums = %v; hope = %v; ret = %v \n", nums, hope, ret)
}
