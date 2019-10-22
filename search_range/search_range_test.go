package search_range

import (
	"testing"
)

func TestSearchRange(t *testing.T) {
	var nums, hope, ret []int
	var target int

	nums = []int{5, 7, 7, 8, 8, 10}
	target = 8
	hope = []int{3, 4}
	ret = searchRange(nums, target)
	t.Logf("nums = %v; target = %d; hope = %v; ret = %v", nums, target, hope, ret)
	// return
	nums = []int{5, 7, 7, 8, 8, 10}
	target = 6
	hope = []int{-1, -1}
	ret = searchRange(nums, target)
	t.Logf("nums = %v; target = %d; hope = %v; ret = %v", nums, target, hope, ret)
	// return
	nums = []int{5, 7, 7, 8, 8, 10}
	target = 10
	hope = []int{5, 5}
	ret = searchRange(nums, target)
	t.Logf("nums = %v; target = %d; hope = %v; ret = %v", nums, target, hope, ret)
	// return
	nums = []int{5, 7, 7, 8, 8, 10}
	target = 5
	hope = []int{0, 0}
	ret = searchRange(nums, target)
	t.Logf("nums = %v; target = %d; hope = %v; ret = %v", nums, target, hope, ret)

	nums = []int{0, 0}
	target = 1
	hope = []int{-1, -1}
	ret = searchRange(nums, target)
	t.Logf("nums = %v; target = %d; hope = %v; ret = %v", nums, target, hope, ret)

	nums = []int{0, 0}
	target = 0
	hope = []int{0, 1}
	ret = searchRange(nums, target)
	t.Logf("nums = %v; target = %d; hope = %v; ret = %v", nums, target, hope, ret)

	nums = []int{0}
	target = 1
	hope = []int{-1, -1}
	ret = searchRange(nums, target)
	t.Logf("nums = %v; target = %d; hope = %v; ret = %v", nums, target, hope, ret)

	nums = []int{0}
	target = 0
	hope = []int{0, 0}
	ret = searchRange(nums, target)
	t.Logf("nums = %v; target = %d; hope = %v; ret = %v", nums, target, hope, ret)
}
