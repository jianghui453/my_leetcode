//假设按照升序排序的数组在预先未知的某个点上进行了旋转。
//
//( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。
//
//编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。
//
//示例 1:
//
//输入: nums = [2,5,6,0,0,1,2], target = 0
//输出: true
//示例 2:
//
//输入: nums = [2,5,6,0,0,1,2], target = 3
//输出: false
//进阶:
//
//这是 搜索旋转排序数组 的延伸题目，本题中的 nums  可能包含重复元素。
//这会影响到程序的时间复杂度吗？会有怎样的影响，为什么？

package binary_search

import (
// "fmt"
)

func search(nums []int, target int) bool {
	l := len(nums)

	for l > 1 {
		if nums[l-1] != nums[0] {
			break
		}
		l--
	}

	if l == 0 {
		return false
	}

	if l == 1 {
		return nums[0] == target
	}

	min, max := 0, l-1
	for min <= max {
		mid := (min + max) / 2
		if target == nums[0] || target == nums[l-1] || target == nums[mid] {
			return true
		}

		if nums[mid] > target {
			if target >= nums[0] {
				max = mid - 1
			} else {
				if nums[mid] >= nums[0] {
					min = mid + 1
				} else {
					max = mid - 1
				}
			}
		} else {
			if target >= nums[0] {
				if nums[mid] >= nums[0] {
					min = mid + 1
				} else {
					max = mid - 1
				}
			} else {
				min = mid + 1
			}
		}
	}

	return false
}

// import "fmt"

// func search(nums []int, target int) bool {
// 	numsLen := len(nums)
// 	if numsLen == 0 {
// 		return false
// 	}
// 	if numsLen == 1 {
// 		return nums[0] == target
// 	}

// 	binarySearch := func (_nums []int) bool {
// 		_numsLen := len(_nums)
// 		min, max := 0, _numsLen-1

// 		for min <= max {
// 			_mid := (min+max)/2
// 			if target > _nums[_mid] {
// 				min = _mid+1
// 			} else if target < _nums[_mid] {
// 				max = _mid-1
// 			} else {
// 				return true
// 			}
// 		}

// 		return false
// 	}

// 	if nums[0] < nums[numsLen-1] {
// 		return binarySearch(nums)
// 	}

// 	left, right := 0, numsLen-1
// 	mid := (left+right)/2
// 	if target == nums[mid] || target == nums[0] || target == nums[numsLen-1] {
// 		return true
// 	}

// 	if mid+1 == numsLen {
// 		return false
// 	}

// 	if target > nums[numsLen-1] {
// 		if target < nums[mid] {
// 			return binarySearch(nums[: mid])
// 		}
// 		if nums[mid] > nums[0] {
// 			return search(nums[mid+1: ], target)
// 		}
// 		return binarySearch(nums[: mid])
// 	} else {
// 		if target > nums[mid] {
// 			return binarySearch(nums[mid+1: ])
// 		}
// 		if nums[mid] > nums[0] {
// 			return search(nums[mid+1: ], target)
// 		}
// 		return search(nums[: mid], target)
// 	}
// }

// func search(nums []int, target int) bool {
// 	lenN := len(nums)
// 	if lenN < 1 {
// 		return false
// 	}
// 	left := 0
// 	right := lenN - 1
// 	start := 0
// 	for left < lenN && right >= 0 && left < right {
// 		if left < lenN-1 && nums[left] > nums[left+1] {
// 			start = left + 1
// 			break
// 		}
// 		if right > 0 && nums[right] < nums[right-1] {
// 			start = right
// 			break
// 		}
// 		left++
// 		right--
// 	}
// 	if start == 0 {
// 		left = 0
// 		right = lenN - 1
// 	} else if target > nums[0] {
// 		left = 0
// 		right = start - 1
// 	} else if target < nums[lenN-1] {
// 		left = start
// 		right = lenN - 1
// 	} else {
// 		return target == nums[0] || target == nums[lenN-1]
// 	}
// 	fmt.Printf("nums=%v target=%d start=%d left=%d right=%d\n", nums, target, start, left, right)
// 	for left <= right {
// 		mid := (left + right) / 2
// 		fmt.Printf("left=%d right=%d mid=%d\n", left, right, mid)
// 		if nums[mid] < target {
// 			left = mid + 1
// 		} else if nums[mid] > target {
// 			if mid == right {
// 				right--
// 			} else {
// 				right = mid
// 			}
// 		} else {
// 			return true
// 		}
// 	}
// 	return false
// }

// func search(nums []int, target int) int {
// 	if len(nums) == 0 {
// 		return -1
// 	}
// 	min, max := 0, len(nums)-1
// 	for min <= max {
// 		fmt.Printf("min = %d; max = %d\n", min, max)
// 		if min == max {
// 			if target == nums[min] {
// 				return min
// 			} else {
// 				return -1
// 			}
// 		}
// 		half := (min + max) / 2
// 		if target > nums[half] {
// 			if nums[half] < nums[0] {
// 				if target < nums[0] {
// 					min = half + 1
// 				} else {
// 					max = half
// 				}
// 			} else {
// 				min = half + 1
// 			}
// 		} else if target < nums[half] {
// 			if nums[half] < nums[0] {
// 				max = half
// 			} else {
// 				if target < nums[0] {
// 					min = half + 1
// 				} else {
// 					max = half
// 				}
// 			}
// 		} else {
// 			return half
// 		}
// 	}
// 	return -1
// }
