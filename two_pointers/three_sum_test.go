package two_pointers

import (
	"testing"
)

func TestThreeSum(t *testing.T) {
	var nums []int
	var h, r [][]int

	nums = []int{
		-1, 0, 1, 2, -1, -4,
	}
	r = threeSum(nums)
	h = [][]int{
		{
			-1, 0, 1,
		},
		{
			-1, -1, 2,
		},
	}
	t.Logf("nums=%v h=%v r=%v\n", nums, h, r)

	nums = []int{
		0, 0, 0, 0, 0,
	}
	r = threeSum(nums)
	h = [][]int{
		{
			0, 0, 0,
		},
	}
	t.Logf("nums=%v h=%v r=%v\n", nums, h, r)

	nums = []int{
		-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0,
	}
	r = threeSum(nums)
	h = [][]int{
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
	t.Logf("nums=%v h=%v r=%v\n", nums, h, r)

	nums = []int{-2, 0, 1, 1, 2}
	r = threeSum(nums)
	h = [][]int{
		{-2, 0, 2}, {-2, 1, 1},
	}
	t.Logf("nums=%v h=%v r=%v\n", nums, h, r)
}
