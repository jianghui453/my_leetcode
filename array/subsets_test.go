package array

import "testing"

func TestSubsets1(t *testing.T) {
	var nums []int
	var hope, ret [][]int

	nums = []int{1, 2, 3}
	hope = [][]int{
		{3},
		{1},
		{2},
		{1, 2, 3},
		{1, 3},
		{2, 3},
		{1, 2},
		{},
	}
	ret = subsets1(nums)
	t.Logf("\nnums=%v \nhope=%v \n ret=%v\n", nums, hope, ret)

	nums = []int{}
	hope = [][]int{{}}
	ret = subsets1(nums)
	t.Logf("\nnums=%v \nhope=%v \n ret=%v\n", nums, hope, ret)

	nums = []int{1}
	hope = [][]int{{}, {1}}
	ret = subsets1(nums)
	t.Logf("\nnums=%v \nhope=%v \n ret=%v\n", nums, hope, ret)
}
