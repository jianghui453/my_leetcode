// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

// 你可以假设数组中无重复元素。

// 示例 1:

// 输入: [1,3,5,6], 5
// 输出: 2
// 示例 2:

// 输入: [1,3,5,6], 2
// 输出: 1
// 示例 3:

// 输入: [1,3,5,6], 7
// 输出: 4
// 示例 4:

// 输入: [1,3,5,6], 0
// 输出: 0

package binary_search

func searchInsert(nums []int, target int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}

	min, max := 0, numsLen-1
	var mid int
	for min <= max {
		mid = (min + max) / 2
		if nums[mid] > target {
			max = mid - 1
		} else if nums[mid] < target {
			min = mid + 1
		} else {
			return mid
		}
	}

	if min == numsLen {
		return min
	}
	if max == -1 {
		return 0
	}
	return min
}

// func searchInsert(nums []int, target int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] >= target {
// 			return i
// 		}
// 		if i+1 < len(nums) {
// 			if nums[i+1] > target {
// 				return i + 1
// 			} else {
// 				continue
// 			}
// 		} else {
// 			return i + 1
// 		}
// 	}
// 	return len(nums)
// }

// func searchInsert(nums []int, target int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}
// 	min, max := 0, len(nums)-1
// 	var half int
// 	for min <= max {
// 		half = (min + max) / 2
// 		fmt.Printf("min = %d; max = %d; half = %d.\n", min, max, half)
// 		if nums[half] < target {
// 			min = half + 1
// 		} else if nums[half] > target {
// 			max = half - 1
// 		} else {
// 			return half
// 		}
// 	}
// 	fmt.Printf("min = %d; max = %d; half = %d.\n", min, max, half)
// 	if nums[half] > target {
// 		return half
// 	} else {
// 		return half + 1
// 	}
// }
