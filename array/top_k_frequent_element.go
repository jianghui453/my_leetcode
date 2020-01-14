// 给定一个非空的整数数组，返回其中出现频率前 k 高的元素。

// 示例 1:

// 输入: nums = [1,1,1,2,2,3], k = 2
// 输出: [1,2]
// 示例 2:

// 输入: nums = [1], k = 1
// 输出: [1]
// 说明：

// 你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
// 你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。

package array

import (
)

func topKFrequent(nums []int, k int) []int {
    var (
		l int = len(nums)
		ret []int
		hash map[int]int = make(map[int]int)
	)

	for i := 0; i < l; i++ {
		if _, ok := hash[nums[i]]; !ok {
			hash[nums[i]] = 1
			ret = append(ret, nums[i])
		} else {
			hash[nums[i]]++
		}
	}
	if k >= len(ret) {
		return ret
	}

	quicksort(hash, ret)

	return ret[: k]
}

func quicksort(hash map[int]int, ret []int) {
	var (
		l int = len(ret)
	)
	if l <= 1 {
		return
	}

	_ret := make([]int, l)
	copy(_ret, ret)

	equal_ret := make([]int, 0)
	left, right := 0, l-1
	for i := 0; i < l; i++ {
		if hash[_ret[i]] > hash[_ret[0]] {
	
			ret[left] = _ret[i]
			left++
		} else if hash[_ret[i]] < hash[_ret[0]] {
	
			ret[right] = _ret[i]
			right-- 
		} else {
			equal_ret = append(equal_ret, _ret[i])
		}
	}

	quicksort(hash, ret[: left])
	if right < l-1 {
		quicksort(hash, ret[right+1: ])
	}


	for i := left; i <= right; i++ {
		ret[i] = equal_ret[i-left]
	}
}
