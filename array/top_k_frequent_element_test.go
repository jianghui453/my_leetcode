package array

import (
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	var (
		nums, hope, ret []int
		k               int
	)

	nums = []int{1, 1, 1, 2, 2, 3}
	k = 2
	hope = []int{1, 2}
	ret = topKFrequent(nums, k)
	t.Log(nums, k, hope, ret)

	nums = []int{1, 1, 1, 2, 2, 3}
	k = 3
	hope = []int{1, 2, 3}
	ret = topKFrequent(nums, k)
	t.Log(nums, k, hope, ret)

	nums = []int{1, 1, 1, 2, 2, 3}
	k = 4
	hope = []int{1, 2, 3}
	ret = topKFrequent(nums, k)
	t.Log(nums, k, hope, ret)

	nums = []int{4, 1, -1, 2, -1, 2, 3}
	k = 2
	hope = []int{-1, 2}
	ret = topKFrequent(nums, k)
	t.Log(nums, k, hope, ret)
}
