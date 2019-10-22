package search_range

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	var min, max = 0, len(nums) - 1
	for min <= max {
		half := (min + max) / 2
		fmt.Printf("min = %d; max = %d; half = %d\n", min, max, half)
		if nums[half] < target {
			min = half + 1
		} else if nums[half] > target {
			max = half - 1
		} else {
			left, right := -1, -1
			for i := min; i <= max; i++ {
				if nums[i] == target {
					if left == -1 {
						left = i
					}
					right = i
				} else if nums[i] > target {
					break
				}
			}
			return []int{left, right}
		}
	}
	return []int{-1, -1}
}

func searchRangeV1(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	var min, max = 0, len(nums) - 1
	for min < max {
		half := (min + max) / 2
		if nums[half] < target {
			min = half + 1
		} else if nums[half] > target {
			max = half
		} else {
			var start, end = half, half
			for {
				if start > 0 && nums[start-1] == target {
					start--
					continue
				}
				if end < len(nums)-1 && nums[end+1] == target {
					end++
					continue
				}
				return []int{start, end}
			}
		}
	}
	if min == max && nums[min] == target {
		return []int{min, min}
	}
	return []int{-1, -1}
}
