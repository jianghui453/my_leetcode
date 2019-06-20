package search

import (
	"fmt"
)

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	min, max := 0, len(nums) - 1
	for min <= max{
		fmt.Printf("min = %d; max = %d\n", min, max)
		if min == max {
			if target == nums[min] {
				return min
			} else {
				return -1
			}
		}
		half := (min + max) / 2
		if target > nums[half] {
			if nums[half] < nums[0] {
				if target < nums[0] {
					min = half + 1
				} else {
					max = half
				}
			} else {
				min = half + 1
			}
		} else if target < nums[half] {
			if nums[half] < nums[0] {
				max = half
			} else {
				if target < nums[0] {
					min = half + 1
				} else {
					max = half
				}
			}
		} else {
			return half
		}
	}
	return -1
}