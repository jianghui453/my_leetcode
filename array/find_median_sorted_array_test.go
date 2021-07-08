package array

import (
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T) {
	var nums1, nums2 []int
	var h, r float64

	nums1 = []int{1}
	nums2 = []int{1}
	r = findMedianSortedArrays(nums1, nums2)
	h = 1.0
	t.Logf("\nnums1=%v \nnums2=%v \nr=%g \nh=%g\n", nums1, nums2, r, h)

	nums1 = []int{1, 3}
	nums2 = []int{2, 4, 5, 6}
	r = findMedianSortedArrays(nums1, nums2)
	h = 3.5
	t.Logf("\nnums1=%v \nnums2=%v \nr=%g \nh=%g\n", nums1, nums2, r, h)
	//return

	nums1 = []int{1, 2}
	nums2 = []int{3, 4, 5}
	r = findMedianSortedArrays(nums1, nums2)
	h = 3.0
	t.Logf("\nnums1=%v \nnums2=%v \nr=%g \nh=%g\n", nums1, nums2, r, h)
	//return

	nums1 = []int{1, 2}
	nums2 = []int{3, 4, 5, 6}
	r = findMedianSortedArrays(nums1, nums2)
	h = 3.5
	t.Logf("\nnums1=%v \nnums2=%v \nr=%g \nh=%g\n", nums1, nums2, r, h)
	//return

	nums1 = []int{1, 3}
	nums2 = []int{2}
	r = findMedianSortedArrays(nums1, nums2)
	h = 2.0
	t.Logf("\nnums1=%v \nnums2=%v \nr=%g \nh=%g\n", nums1, nums2, r, h)
	//return

	nums1 = []int{1, 2}
	nums2 = []int{1, 2, 3}
	r = findMedianSortedArrays(nums1, nums2)
	h = 2.0
	t.Logf("\nnums1=%v \nnums2=%v \nr=%g \nh=%g\n", nums1, nums2, r, h)
	//return
	r = findMedianSortedArrays(nums1, nums2)
	h = 2.0
	t.Logf("\nnums1=%v \nnums2=%v \nr=%g \nh=%g\n", nums1, nums2, r, h)

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	r = findMedianSortedArrays(nums1, nums2)
	h = 2.5
	t.Logf("\nnums1=%v \nnums2=%v \nr=%g \nh=%g\n", nums1, nums2, r, h)
	//return
	nums1 = []int{}
	nums2 = []int{3, 4}
	r = findMedianSortedArrays(nums1, nums2)
	h = 3.5
	t.Logf("\nnums1=%v \nnums2=%v \nr=%g \nh=%g\n", nums1, nums2, r, h)
	//return
	nums1 = []int{}
	nums2 = []int{}
	r = findMedianSortedArrays(nums1, nums2)
	h = 0.0
	t.Logf("\nnums1=%v \nnums2=%v \nr=%g \nh=%g\n", nums1, nums2, r, h)

	nums1 = []int{1, 2, 3}
	nums2 = []int{1}
	r = findMedianSortedArrays(nums1, nums2)
	h = 1.5
	t.Logf("\nnums1=%v \nnums2=%v \nr=%g \nh=%g\n", nums1, nums2, r, h)

	nums1 = []int{1, 3, 5}
	nums2 = []int{1}
	r = findMedianSortedArrays(nums1, nums2)
	h = 2.0
	t.Logf("\nnums1=%v \nnums2=%v \nr=%g \nh=%g\n", nums1, nums2, r, h)

	nums1 = []int{1, 2}
	nums2 = []int{-1, 3}
	r = findMedianSortedArrays(nums1, nums2)
	h = 1.5
	t.Logf("\nnums1=%v \nnums2=%v \nr=%g \nh=%g\n", nums1, nums2, r, h)
}
