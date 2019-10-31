package array

import (
	"testing"
)

func TestNextPermutation(t *testing.T) {
	var nums, h []int

	nums = []int{1, 2, 3}
	h = []int{1, 3, 2}
	nextPermutation(nums)
	t.Logf("\nnums=%v \nh=%v", nums, h)

	nums = []int{3, 2, 1}
	h = []int{1, 2, 3}
	nextPermutation(nums)
	t.Logf("\nnums=%v \nh=%v", nums, h)

	nums = []int{1, 3, 2}
	h = []int{2, 1, 3}
	nextPermutation(nums)
	t.Logf("\nnums=%v \nh=%v", nums, h)

	nums = []int{1, 1, 5}
	h = []int{1, 5, 1}
	nextPermutation(nums)
	t.Logf("\nnums=%v \nh=%v", nums, h)

	nums = []int{1, 5, 1}
	h = []int{5, 1, 1}
	nextPermutation(nums)
	t.Logf("\nnums=%v \nh=%v", nums, h)

	nums = []int{5, 1, 1}
	h = []int{1, 1, 5}
	nextPermutation(nums)
	t.Logf("\nnums=%v \nh=%v", nums, h)

	nums = []int{4, 3, 2, 1}
	h = []int{1, 2, 3, 4}
	nextPermutation(nums)
	t.Logf("\nnums=%v \nh=%v", nums, h)

	nums = []int{1, 2, 4, 5, 3}
	h = []int{1, 2, 5, 3, 4}
	nextPermutation(nums)
	t.Logf("\nnums=%v \nh=%v", nums, h)

	nums = []int{1}
	h = []int{1}
	nextPermutation(nums)
	t.Logf("\nnums=%v \nh=%v", nums, h)

	nums = []int{}
	h = []int{}
	nextPermutation(nums)
	t.Logf("\nnums=%v \nh=%v", nums, h)
}
