package leet_code

import (
	"fmt"
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T) {
	//nums := []int{1, 2}
	//fmt.Printf("nums[:1]=%v; nums[1:]=%v", nums[:1], nums[1:])
	//i := 3
	//fmt.Println(i / 2)
	//return

	var nums1, nums2 []int
	var m float64

	nums1 = []int{1, 3}
	nums2 = []int{2, 4, 5, 6}
	m = findMedianSortedArrays(nums1, nums2)
	if m == 3.5 {
		t.Log(fmt.Sprintf("success: nums1=%v; nums2=%v; m=%f", nums1, nums2, m))
	} else {
		t.Error(fmt.Sprintf("fail: nums1=%v; nums2=%v; m=%f", nums1, nums2, m))
	}
	return

	nums1 = []int{1, 2}
	nums2 = []int{3, 4, 5}
	m = findMedianSortedArrays(nums1, nums2)
	if m == 3.0 {
		t.Log(fmt.Sprintf("success: nums1=%v; nums2=%v; m=%f", nums1, nums2, m))
	} else {
		t.Error(fmt.Sprintf("fail: nums1=%v; nums2=%v; m=%f", nums1, nums2, m))
	}
	return

	nums1 = []int{1, 2}
	nums2 = []int{3, 4, 5, 6}
	m = findMedianSortedArrays(nums1, nums2)
	if m == 3.5 {
		t.Log(fmt.Sprintf("success: nums1=%v; nums2=%v; m=%f", nums1, nums2, m))
	} else {
		t.Error(fmt.Sprintf("fail: nums1=%v; nums2=%v; m=%f", nums1, nums2, m))
	}
	//return

	nums1 = []int{1, 3}
	nums2 = []int{2}
	m = findMedianSortedArrays(nums1, nums2)
	if m == 2.0 {
		t.Log(fmt.Sprintf("success: nums1=%v; nums2=%v; m=%f", nums1, nums2, m))
	} else {
		t.Error(fmt.Sprintf("fail: nums1=%v; nums2=%v; m=%f", nums1, nums2, m))
	}
	//return

	nums1 = []int{1, 2}
	nums2 = []int{1, 2, 3}
	m = findMedianSortedArrays(nums1, nums2)
	if m == 2.0 {
		t.Log(fmt.Sprintf("success: nums1=%v; nums2=%v; m=%f", nums1, nums2, m))
	} else {
		t.Error(fmt.Sprintf("fail: nums1=%v; nums2=%v; m=%f", nums1, nums2, m))
	}
	//return
	m = findMedianSortedArrays(nums1, nums2)
	if m == 2.0 {
		t.Log(fmt.Sprintf("success: nums1=%v; nums2=%v; m=%f", nums1, nums2, m))
	} else {
		t.Error(fmt.Sprintf("fail: nums1=%v; nums2=%v; m=%f", nums1, nums2, m))
	}

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	m = findMedianSortedArrays(nums1, nums2)
	if m == 2.5 {
		t.Log(fmt.Sprintf("success: nums1=%v; nums2=%v; m=%f", nums1, nums2, m))
	} else {
		t.Error(fmt.Sprintf("fail: nums1=%v; nums2=%v; m=%f", nums1, nums2, m))
	}
	//return
	nums1 = []int{}
	nums2 = []int{3, 4}
	m = findMedianSortedArrays(nums1, nums2)
	if m == 3.5 {
		t.Log(fmt.Sprintf("success: nums1=%v; nums2=%v; m=%f", nums1, nums2, m))
	} else {
		t.Error(fmt.Sprintf("fail: nums1=%v; nums2=%v; m=%f", nums1, nums2, m))
	}

	nums1 = []int{}
	nums2 = []int{}
	m = findMedianSortedArrays(nums1, nums2)
	if m == 0.0 {
		t.Log(fmt.Sprintf("success: nums1=%v; nums2=%v; m=%f", nums1, nums2, m))
	} else {
		t.Error(fmt.Sprintf("fail: nums1=%v; nums2=%v; m=%f", nums1, nums2, m))
	}

	nums1 = []int{1, 2, 3}
	nums2 = []int{1}
	m = findMedianSortedArrays(nums1, nums2)
	if m == 1.5 {
		t.Log(fmt.Sprintf("success: nums1=%v; nums2=%v; m=%f", nums1, nums2, m))
	} else {
		t.Error(fmt.Sprintf("fail: nums1=%v; nums2=%v; m=%f", nums1, nums2, m))
	}

	nums1 = []int{1, 3, 5}
	nums2 = []int{1}
	m = findMedianSortedArrays(nums1, nums2)
	if m == 2.0 {
		t.Log(fmt.Sprintf("success: nums1=%v; nums2=%v; m=%f", nums1, nums2, m))
	} else {
		t.Error(fmt.Sprintf("fail: nums1=%v; nums2=%v; m=%f", nums1, nums2, m))
	}

	nums1 = []int{1, 2}
	nums2 = []int{-1, 3}
	m = findMedianSortedArrays(nums1, nums2)
	if m == 1.5 {
		t.Log(fmt.Sprintf("success: nums1=%v; nums2=%v; m=%f", nums1, nums2, m))
	} else {
		t.Error(fmt.Sprintf("fail: nums1=%v; nums2=%v; m=%f", nums1, nums2, m))
	}
}
