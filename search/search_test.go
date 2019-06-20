package search

import (
	"testing"
)

func TestSearch(t *testing.T) {
	var nums []int
	var target, hope, ret int

	nums = []int{4, 5, 6, 7, 0, 1, 2}
	target = 0
	hope = 4
	ret = search(nums, target)
	t.Logf("nums = %v; target = %d; hope = %d; ret = %d\n", nums, target, hope, ret)
	
	target = 6
	hope = 2
	ret = search(nums, target)
	t.Logf("nums = %v; target = %d; hope = %d; ret = %d\n", nums, target, hope, ret)
	
	target = 4
	hope = 0
	ret = search(nums, target)
	t.Logf("nums = %v; target = %d; hope = %d; ret = %d\n", nums, target, hope, ret)
	
	target = 2
	hope = 6
	ret = search(nums, target)
	t.Logf("nums = %v; target = %d; hope = %d; ret = %d\n", nums, target, hope, ret)

	target = 3
	hope = -1
	ret = search(nums, target)
	t.Logf("nums = %v; target = %d; hope = %d; ret = %d\n", nums, target, hope, ret)

	nums = []int{4, 5, 6, 7, 0}
	target = 4
	hope = 0
	ret = search(nums, target)
	t.Logf("nums = %v; target = %d; hope = %d; ret = %d\n", nums, target, hope, ret)
	
	target = 0
	hope = 4
	ret = search(nums, target)
	t.Logf("nums = %v; target = %d; hope = %d; ret = %d\n", nums, target, hope, ret)
	
	target = 6
	hope = 2
	ret = search(nums, target)
	t.Logf("nums = %v; target = %d; hope = %d; ret = %d\n", nums, target, hope, ret)

	nums = []int{0}
	target = 1
	hope = -1
	ret = search(nums, target)
	t.Logf("nums = %v; target = %d; hope = %d; ret = %d\n", nums, target, hope, ret)
	
	target = 0
	hope = 0
	ret = search(nums, target)
	t.Logf("nums = %v; target = %d; hope = %d; ret = %d\n", nums, target, hope, ret)

	nums = []int{1, 3}
	target = 2
	hope = -1
	ret = search(nums, target)
	t.Logf("nums = %v; target = %d; hope = %d; ret = %d\n", nums, target, hope, ret)
	
	target = 3
	hope = 1
	ret = search(nums, target)
	t.Logf("nums = %v; target = %d; hope = %d; ret = %d\n", nums, target, hope, ret)
}