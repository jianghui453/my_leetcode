// 给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。

// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。

// 示例 1:

// 输入: [3,2,3]
// 输出: 3
// 示例 2:

// 输入: [2,2,1,1,1,2,2]
// 输出: 2

package array

import (
// "fmt"
// "math"
)

func majorityElement(nums []int) int {
	l := len(nums)
	m := make(map[int]int, 0)
	for i := 0; i < l; i++ {
		if _, ok := m[nums[i]]; ok {
			m[nums[i]]++
		} else {
			m[nums[i]] = 1
		}
		if m[nums[i]] > l/2 {
			return nums[i]
		}
	}
	return 0
}
