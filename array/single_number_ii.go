//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。
//
//说明：
//
//你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
//
//示例 1:
//
//输入: [2,2,3,2]
//输出: 3
//示例 2:
//
//输入: [0,1,0,1,0,1,99]
//输出: 99

package array

import (
// "fmt"
)

// 	// a b  nums[i]  a b
// 	// 0 0  0        0 0
// 	// 1 0  1        0 0

// 	// 1 0  0        1 0
// 	// 0 1  1        1 0

// 	// 0 1  0        0 1
// 	// 0 0  1        0 1
func singleNumber2(nums []int) int {
	l, a, b := len(nums), 0, 0
	for i := 0; i < l; i++ {
		a, b = (a^b)&(a^nums[i]), (b^nums[i])&(a^-1)
	}
	return b
}
