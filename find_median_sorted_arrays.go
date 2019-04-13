package leet_code

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	for {
		if len(nums1) == 0 {
			return getMedian(nums2)
		}
		if len(nums2) == 0 {
			return getMedian(nums1)
		}
		if len(nums1) == 1 {
			return getMedianFromArrayAndInt(nums2, nums1[0])
		}
		if len(nums2) == 1 {
			return getMedianFromArrayAndInt(nums1, nums2[0])
		}
		if getMedian(nums1) == getMedian(nums2) {
			return getMedian(nums1)
		}
		if nums1[0] > nums2[len(nums2)-1] {
			nums1 = append(nums2, nums1...)
			if len(nums1)%2 == 0 {
				return float64(nums1[len(nums1)/2]+nums1[len(nums1)/2-1]) / 2
			}
			return float64(nums1[len(nums1)/2])
		}
		if nums2[0] > nums1[len(nums1)-1] {
			nums1 = append(nums1, nums2...)
			if len(nums1)%2 == 0 {
				return float64(nums1[len(nums1)/2]+nums1[len(nums1)/2-1]) / 2
			}
			return float64(nums1[len(nums1)/2])
		}
		nums1, nums2 = splitSlice(nums1, nums2)
	}
}

func splitSlice(nums1, nums2 []int) ([]int, []int) {
	nlist := [4][]int{}
	if getMedian(nums1) < getMedian(nums2) {
		nums1, nums2 = nums2, nums1
	}
	halflen1, halflen2 := len(nums1)/2, len(nums2)/2
	if float64(nums1[halflen1]) == getMedian(nums1) && len(nums1)%2 != 0 {
		nlist[0] = nums1[:halflen1+1]
		nlist[1] = nums1[halflen1+1:]
	} else {
		nlist[0] = nums1[:halflen1]
		nlist[1] = nums1[halflen1:]
	}
	if float64(nums2[halflen2]) == getMedian(nums2) && len(nums2)%2 == 0 {
		nlist[2] = nums2[:halflen2-1]
		nlist[3] = nums2[halflen2-1:]
	} else {
		nlist[2] = nums2[:halflen2]
		nlist[3] = nums2[halflen2:]
	}
	for i := 0; i < 4; i++ {
		for j := i - 1; j >= 0; j-- {
			if getMedian(nlist[j]) > getMedian(nlist[j+1]) {
				nlist[j], nlist[j+1] = nlist[j+1], nlist[j]
			}
		}
	}
	fmt.Printf("nums1=%v; nums2=%v; nlist=%v\n", nums1, nums2, nlist)
	return nlist[1], nlist[2]
}

func getMedianFromArrayAndInt(nums []int, i int) float64 {
	if len(nums) == 1 {
		return float64(nums[0]+i) / 2
	}
	halflen := len(nums) / 2
	if len(nums)%2 == 0 {
		if i <= nums[halflen-1] {
			return float64(nums[halflen-1])
		}
		if i >= nums[halflen] {
			return float64(nums[halflen])
		}
		return float64(i)
	}
	if i <= nums[halflen-1] {
		return float64(nums[halflen-1]+nums[halflen]) / 2
	}
	if i <= nums[halflen+1] {
		return float64(nums[halflen]+i) / 2
	}
	return float64(nums[halflen]+nums[halflen+1]) / 2
}

func getMedian(nums []int) float64 {
	if len(nums) == 0 {
		return 0.0
	}
	if len(nums)%2 != 0 {
		return float64(nums[len(nums)/2])
	}
	return float64(nums[len(nums)/2]+nums[len(nums)/2-1]) / 2
}
