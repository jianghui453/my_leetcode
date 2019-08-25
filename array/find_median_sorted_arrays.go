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

package array

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	if len1 < len2 {
		len1, len2, nums1, nums2 = len2, len1, nums2, nums1
	}
	fmt.Printf("start: nums1=%v; nums2=%v; len1=%d; len2=%d\n", nums1, nums2, len1, len2)
	if len2 == 0 {
		if len1 == 0 {
			return 0.0
		}
		if len1%2 == 0 {
			return float64(nums1[len1/2]+nums1[len1/2-1]) / 2
		}
		return float64(nums1[len1/2])
	}
	halfLen := int(math.Ceil(float64(len1+len2) / 2))
	iMin, iMax := halfLen-len2, halfLen
	for {
		fmt.Printf("iMin=%d; iMax=%d\n", iMin, iMax)
		offset1 := (iMin + iMax) / 2
		offset2 := halfLen - offset1
		fmt.Printf("offset1=%d; offset2=%d\n", offset1, offset2)
		if offset1 > 0 && offset2 < len2 && nums1[offset1-1] > nums2[offset2] {
			iMax = offset1 - 1
			continue
		}
		if offset2 > 0 && offset1 < len1 && nums2[offset2-1] > nums1[offset1] {
			iMin = offset1 + 1
			continue
		}
		var leftMax float64
		if offset1 == 0 {
			leftMax = float64(nums2[offset2-1])
		} else if offset2 == 0 {
			leftMax = float64(nums1[offset1-1])
		} else {
			leftMax = math.Max(float64(nums1[offset1-1]), float64(nums2[offset2-1]))
		}
		fmt.Printf("leftMax=%f\n", leftMax)
		if (len1+len2)%2 != 0 {
			return leftMax
		}
		var rightMin float64
		if offset1 == len1 {
			rightMin = float64(nums2[offset2])
		} else if offset2 == len2 {
			rightMin = float64(nums1[offset1])
		} else {
			rightMin = math.Min(float64(nums1[offset1]), float64(nums2[offset2]))
		}
		fmt.Printf("rightMin=%f\n", rightMin)
		return (leftMax + rightMin) / 2
	}
}
