package array

import "testing"

func TestMerge1(t *testing.T) {
	var nums1, nums2, hope []int
	var m, n int

	nums1 = []int{1, 2, 3, 0, 0, 0}
	nums2 = []int{2, 5, 6}
	m = 3
	n = 3
	hope = []int{1, 2, 2, 3, 5, 6}
	merge1(nums1, m, nums2, n)
	t.Logf("\n hope=%v \nnums1=%v", hope, nums1)
}
