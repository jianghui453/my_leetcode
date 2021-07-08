package array

import (
	"testing"
)

func TestIntersection(t *testing.T) {
	var (
		nums1, nums2, hope, ret []int
	)

	nums1 = []int{1, 2, 2, 1}
	nums2 = []int{2, 2}
	hope = []int{2}
	ret = intersection(nums1, nums2)
	t.Log(nums1, nums2)
	t.Log(hope)
	t.Log(ret)

	nums1 = []int{4, 9, 5}
	nums2 = []int{9, 4, 9, 8, 4}
	hope = []int{4, 9}
	ret = intersection(nums1, nums2)
	t.Log(nums1, nums2)
	t.Log(hope)
	t.Log(ret)
}
