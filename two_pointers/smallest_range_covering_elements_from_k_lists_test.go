package two_pointers

import (
	"testing"
)

func TestSmallestRange(t *testing.T) {
	var nums [][]int
	var hope, ret []int

	nums = [][]int{
		{4, 10, 15, 24, 26},
		{0, 9, 12, 20},
		{5, 18, 22, 30},
	}
	hope = []int{20, 24}
	ret = smallestRange(nums)
	t.Log(nums, hope, ret)

	nums = [][]int{
		{4, 10, 15, 24, 26},
	}
	hope = []int{24, 26}
	ret = smallestRange(nums)
	t.Log(nums, hope, ret)

	nums = [][]int{
		{4, 10, 15, 24, 26},
		{0, 9, 12, 20},
	}
	hope = []int{9, 10}
	ret = smallestRange(nums)
	t.Log(nums, hope, ret)

	nums = [][]int{
		{1,4,7,10},
		{2,5,8,11},
		{3,6,9,12},
	}
	hope = []int{1, 3}
	ret = smallestRange(nums)
	t.Log(nums, hope, ret)

	nums = [][]int{
		{66,83,85},{-67,-67,-30,-12,-6,13,16,20,20,22,26,26,27},{2,2,5,7,9,10,11,12,13},{-17,24,36,47},{37,38,38,38,40},{-20,8,10,11,13,13,13,14},{-5,36,45,57,61,65,67,68,69},{-19,35,59,60},{91,96},{-33,-16,21,30,47,48,50,73,77,77,77,78,78,78,80},{8,9,9,9,9,9,9,9,12,12,13},{1,12},{9,18},{25,30,52,54,57},{14,25,31,34},{57,59,65},{-15,82,94,95,95,97,99},{-13,2,26,57,85,87,90,90,91,91,92},
	}
	hope = []int{12, 91}
	ret = smallestRange(nums)
	t.Log(nums, hope, ret)

	nums = [][]int{
		{1,2,3},{1,2,3},{1,2,3},
	}
	hope = []int{1, 1}
	ret = smallestRange(nums)
	t.Log(nums, hope, ret)
}