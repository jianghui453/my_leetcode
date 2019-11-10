package two_pointers

import "testing"

func TestSortColors(t *testing.T) {
	var nums, hope []int

	nums = []int{2, 0, 2, 1, 1, 0}
	hope = []int{0, 0, 1, 1, 2, 2}
	sortColors(nums)
	t.Logf("\nnums=%v \nhope=%v\n", nums, hope)

	nums = []int{0, 2, 1, 0, 2, 1, 0, 0, 2, 1, 0, 2}
	hope = []int{0, 0, 0, 0, 0, 1, 1, 1, 2, 2, 2, 2}
	sortColors(nums)
	t.Logf("\nnums=%v \nhope=%v\n", nums, hope)

	nums = []int{}
	hope = []int{}
	sortColors(nums)
	t.Logf("\nnums=%v \nhope=%v\n", nums, hope)

	nums = []int{1, 0}
	hope = []int{0, 1}
	sortColors(nums)
	t.Logf("\nnums=%v \nhope=%v\n", nums, hope)

	nums = []int{1, 1, 1}
	hope = []int{1, 1, 1}
	sortColors(nums)
	t.Logf("\nnums=%v \nhope=%v\n", nums, hope)

	nums = []int{0}
	hope = []int{0}
	sortColors(nums)
	t.Logf("\nnums=%v \nhope=%v\n", nums, hope)

	nums = []int{2, 1, 0}
	hope = []int{0, 1, 2}
	sortColors(nums)
	t.Logf("\nnums=%v \nhope=%v\n", nums, hope)
}
