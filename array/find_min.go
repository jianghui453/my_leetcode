package array

import (
	// "fmt"
)

// 153
func findMin153(nums []int) int {
	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	for len(nums) > 2 {
		mid := nums[len(nums)/2]
		if mid >= nums[len(nums)-1] {
			nums = nums[len(nums)/2:]
		} else {
			nums = nums[:len(nums)/2+1]
		}
		// fmt.Println(nums)
	}
	if len(nums) == 2 {
		return min(nums[0], nums[1])
	}
	if len(nums) == 0 {
		return 0
	}
	return nums[0]
}

func findMin154(nums []int) int {
	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	for len(nums) > 2 {
		mid := nums[len(nums)/2]
		if mid > nums[len(nums)-1] {
			nums = nums[len(nums)/2:]
		} else if mid < nums[len(nums)-1] {
			nums = nums[:len(nums)/2+1]
		} else {
			return min(findMin154(nums[: len(nums)/2]), findMin154(nums[len(nums)/2: ]))
		}
		// fmt.Println(nums)
	}
	if len(nums) == 2 {
		return min(nums[0], nums[1])
	}
	if len(nums) == 0 {
		return 0
	}
	return nums[0]
}
