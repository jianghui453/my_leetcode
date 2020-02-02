// 给定一个整数数组 a，其中1 ≤ a[i] ≤ n （n为数组长度）, 其中有些元素出现两次而其他元素出现一次。

// 找到所有出现两次的元素。

// 你可以不用到任何额外空间并在O(n)时间复杂度内解决这个问题吗？

// 示例：

// 输入:
// [4,3,2,7,8,2,3,1]

// 输出:
// [2,3]

package array

import (
)

func findDuplicates(nums []int) []int {
    var (
		n = len(nums)
		ret []int
	)

	for i := 0; i < n; i++ {
		idx := abs(nums[i])-1
		if nums[idx] < 0 {
			ret = append(ret, idx+1)
		} else {
			nums[idx] = -nums[idx]
		}
	}

	return ret
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}