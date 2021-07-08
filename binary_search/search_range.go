// 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

// 你的算法时间复杂度必须是 O(log n) 级别。

// 如果数组中不存在目标值，返回 [-1, -1]。

// 示例 1:

// 输入: nums = [5,7,7,8,8,10], target = 8
// 输出: [3,4]
// 示例 2:

// 输入: nums = [5,7,7,8,8,10], target = 6
// 输出: [-1,-1]

package binary_search

// import (
// 	"fmt"
// )

func searchRange(nums []int, target int) []int {
	numsLen := len(nums)
	if numsLen == 0 {
		return []int{-1, -1}
	}

	min, max := 0, len(nums)-1
	var mid int
	for min <= max {
		mid = (min + max) / 2
		if nums[mid] > target {
			max = mid - 1
		} else if nums[mid] < target {
			min = mid + 1
		} else {
			left, right := mid, mid
			if left > 0 {
				retLeft := searchRange(nums[:mid], target)
				if retLeft[0] != -1 {
					left = retLeft[0]
				}
			}
			if right < numsLen-1 {
				retRight := searchRange(nums[mid+1:], target)
				if retRight[1] != -1 {
					right = retRight[1] + mid + 1
				}
			}
			return []int{left, right}
		}
	}

	return []int{-1, -1}
}

// func searchRange(nums []int, target int) []int {
// 	if len(nums) == 0 {
// 		return []int{-1, -1}
// 	}
// 	var min, max = 0, len(nums) - 1
// 	for min <= max {
// 		half := (min + max) / 2
// 		fmt.Printf("min = %d; max = %d; half = %d\n", min, max, half)
// 		if nums[half] < target {
// 			min = half + 1
// 		} else if nums[half] > target {
// 			max = half - 1
// 		} else {
// 			left, right := -1, -1
// 			for i := min; i <= max; i++ {
// 				if nums[i] == target {
// 					if left == -1 {
// 						left = i
// 					}
// 					right = i
// 				} else if nums[i] > target {
// 					break
// 				}
// 			}
// 			return []int{left, right}
// 		}
// 	}
// 	return []int{-1, -1}
// }

// func searchRange(nums []int, target int) []int {
// 	if len(nums) == 0 {
// 		return []int{-1, -1}
// 	}
// 	var min, max = 0, len(nums) - 1
// 	for min < max {
// 		half := (min + max) / 2
// 		if nums[half] < target {
// 			min = half + 1
// 		} else if nums[half] > target {
// 			max = half
// 		} else {
// 			var start, end = half, half
// 			for {
// 				if start > 0 && nums[start-1] == target {
// 					start--
// 					continue
// 				}
// 				if end < len(nums)-1 && nums[end+1] == target {
// 					end++
// 					continue
// 				}
// 				return []int{start, end}
// 			}
// 		}
// 	}
// 	if min == max && nums[min] == target {
// 		return []int{min, min}
// 	}
// 	return []int{-1, -1}
// }
