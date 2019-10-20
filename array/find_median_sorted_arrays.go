//给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
//
//请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
//
//你可以假设 nums1 和 nums2 不会同时为空。
//
//示例 1:
//
//nums1 = [1, 3]
//nums2 = [2]
//
//则中位数是 2.0
//示例 2:
//
//nums1 = [1, 2]
//nums2 = [3, 4]
//
//则中位数是 (2 + 3)/2 = 2.5

//标签：数组、二分、分治

package array

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	if len1 < len2 {
		len1, len2, nums1, nums2 = len2, len1, nums2, nums1
	}
	if len2 == 0 {
		if len1 == 0 {
			return 0.0
		}
		if len1 % 2 == 0 {
			return float64(nums1[len1/2] + nums1[len1/2-1]) / 2
		} else {
			return float64(nums1[len1/2])
		}
	}
	halfLen := (len1+len2+1) / 2
	min1, max1 := halfLen-len2, halfLen
	for min1 <= max1 {
		m1 := (min1 + max1) / 2
		m2 := halfLen - m1
		if m1 > 0 && m2 < len2 && nums1[m1-1] > nums2[m2] {
			max1 = m1 - 1
		} else if m2 > 0 && m1 < len1 && nums2[m2-1] > nums1[m1] {
			min1 = m1 + 1
		} else {
			var min float64
			if m1 == 0 {
				min = float64(nums2[m2-1])
			} else if m2 == 0 {
				min = float64(nums1[m1-1])
			} else {
				min = math.Max(float64(nums1[m1-1]), float64(nums2[m2-1]))
			}
			if (len1+len2) % 2 == 1 {
				return min
			}
			var max float64
			if m1 == len1 {
				max = float64(nums2[m2])
			} else if m2 == len2 {
				max = float64(nums1[m1])
			} else {
				max = math.Min(float64(nums1[m1]), float64(nums2[m2]))
			}
			return (min + max) / 2
		}
	}
	return 0.0
}
