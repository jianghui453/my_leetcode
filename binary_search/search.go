// 假设按照升序排序的数组在预先未知的某个点上进行了旋转。

// ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

// 搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

// 你可以假设数组中不存在重复的元素。

// 你的算法时间复杂度必须是 O(log n) 级别。

// 示例 1:

// 输入: nums = [4,5,6,7,0,1,2], target = 0
// 输出: 4
// 示例 2:

// 输入: nums = [4,5,6,7,0,1,2], target = 3
// 输出: -1

package binary_search

func search(nums []int, target int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return -1
	}
	if numsLen == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	binarySearch := func (_nums []int) int {
		_numsLen := len(_nums)
		min, max := 0, _numsLen-1
		
		for min <= max {
			_mid := (min+max)/2
			if target > _nums[_mid] {
				min = _mid+1
			} else if target < _nums[_mid] {
				max = _mid-1
			} else {
				return _mid
			}
		}

		return -1
	}

	if nums[0] < nums[numsLen-1] {
		return binarySearch(nums)
	}

	left, right := 0, numsLen-1
	mid := (left+right)/2

	if target > nums[mid] {
		if nums[mid] > nums[numsLen-1] {
			if mid == numsLen-1 {
				return -1
			}
			ret := search(nums[mid+1: ], target)
			if ret == -1 {
				return -1
			}
			return ret+mid+1
		}

		if target > nums[numsLen-1] {
			return search(nums[: mid], target)
		} 

		if mid == numsLen-1 {
			return -1
		}
		ret := binarySearch(nums[mid+1: ])
		if ret == -1 {
			return -1
		}
		return ret+mid+1
	} else if target < nums[mid] {
		if nums[mid] < nums[0] {
			return search(nums[: mid], target)
		}

		if  target > nums[numsLen-1] {
			return binarySearch(nums[: mid])
		}

		if mid == numsLen-1 {
			return -1
		}
		ret := search(nums[mid+1: ], target)
		if ret == -1 {
			return -1
		}
		return ret+mid+1
	} else {
		return mid
	}
}