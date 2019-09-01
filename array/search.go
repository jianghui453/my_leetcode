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

package array

import "fmt"

func search(nums []int, target int) bool {
	lenN := len(nums)
	if lenN < 1 {
		return false
	}
	left := 0
	right := lenN - 1
	start := 0
	for left < lenN && right >= 0 && left < right {
		if left < lenN-1 && nums[left] > nums[left+1] {
			start = left + 1
			break
		}
		if right > 0 && nums[right] < nums[right-1] {
			start = right
			break
		}
		left++
		right--
	}
	if start == 0 {
		left = 0
		right = lenN - 1
	} else if target > nums[0] {
		left = 0
		right = start - 1
	} else if target < nums[lenN-1] {
		left = start
		right = lenN - 1
	} else {
		return target == nums[0] || target == nums[lenN-1]
	}
	fmt.Printf("nums=%v target=%d start=%d left=%d right=%d\n", nums, target, start, left, right)
	for left <= right {
		mid := (left + right) / 2
		fmt.Printf("left=%d right=%d mid=%d\n", left, right, mid)
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			if mid == right {
				right--
			} else {
				right = mid
			}
		} else {
			return true
		}
	}
	return false
}
