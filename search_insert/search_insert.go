package search_insert

import (
	"fmt"
)
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	for i := 0; i < len(nums); i ++ {
		if nums[i] >= target {
			return i
		}
		if i + 1 < len(nums) {
			if nums[i + 1] > target {
				return i + 1
			} else {
				continue
			}
		} else {
			return i + 1
		}
	}
	return len(nums)
}

func searchInsertV1(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	min, max := 0, len(nums) - 1
	var half int
	for min <= max {
		half = (min + max) / 2
		fmt.Printf("min = %d; max = %d; half = %d.\n", min, max, half)
		if nums[half] < target {
			min = half + 1
		} else if nums[half] > target {
			max = half - 1
		} else {
			return half
		}
	}
	fmt.Printf("min = %d; max = %d; half = %d.\n", min, max, half)
	if nums[half] > target {
		return half
	} else {
		return half + 1
	}
}