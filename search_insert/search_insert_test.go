package search_insert

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	var nums []int
	var target, hope, ret int

	nums = []int{1}
	target = 1
	hope = 0
	ret = searchInsert(nums, target)
	t.Logf("nums = %v; target = %d; hope = %d; ret = %d;\n", nums, target, hope, ret)
	// return

	nums = []int{1}
	target = 0
	hope = 0
	ret = searchInsert(nums, target)
	t.Logf("nums = %v; target = %d; hope = %d; ret = %d;\n", nums, target, hope, ret)

	nums = []int{1}
	target = 0
	hope = 0
	ret = searchInsert(nums, target)
	t.Logf("nums = %v; target = %d; hope = %d; ret = %d;\n", nums, target, hope, ret)

	nums = []int{1}
	target = 2
	hope = 1
	ret = searchInsert(nums, target)
	t.Logf("nums = %v; target = %d; hope = %d; ret = %d;\n", nums, target, hope, ret)
	// return
	nums = []int{1, 2}
	target = 2
	hope = 1
	ret = searchInsert(nums, target)
	t.Logf("nums = %v; target = %d; hope = %d; ret = %d;\n", nums, target, hope, ret)

	nums = []int{1, 2}
	target = 1
	hope = 0
	ret = searchInsert(nums, target)
	t.Logf("nums = %v; target = %d; hope = %d; ret = %d;\n", nums, target, hope, ret)

	nums = []int{1, 2}
	target = 3
	hope = 2
	ret = searchInsert(nums, target)
	t.Logf("nums = %v; target = %d; hope = %d; ret = %d;\n", nums, target, hope, ret)

	nums = []int{1, 2}
	target = 0
	hope = 0
	ret = searchInsert(nums, target)
	t.Logf("nums = %v; target = %d; hope = %d; ret = %d;\n", nums, target, hope, ret)

	nums = []int{1, 3, 5, 6}
	target = 2
	hope = 1
	ret = searchInsert(nums, target)
	t.Logf("nums = %v; target = %d; hope = %d; ret = %d;\n", nums, target, hope, ret)

	nums = []int{1, 3}
	target = 2
	hope = 1
	ret = searchInsert(nums, target)
	t.Logf("nums = %v; target = %d; hope = %d; ret = %d;\n", nums, target, hope, ret)
}
