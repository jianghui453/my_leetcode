// 在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

// 示例 1:

// 输入: [3,2,1,5,6,4] 和 k = 2
// 输出: 5
// 示例 2:

// 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
// 输出: 4
// 说明:

// 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。

package array

import (
	"sort"
)

func findKthLargest(nums []int, k int) int {
	s := make([]int, k)
	copy(s, nums[: k])
	sort.Ints(s)
	l := len(nums)
	for i := k; i < l; i++ {
		if nums[i] < s[0] {
			continue
		}
		s[0] = nums[i]
		for j := 1; j < k; j++ {
			if s[j-1] > s[j] {
				s[j-1], s[j] = s[j], s[j-1]
			} else {
				break
			}
		}
	}

	return s[0]
}