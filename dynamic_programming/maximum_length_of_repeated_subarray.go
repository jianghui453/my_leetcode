// 给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度。

// 示例 1:

// 输入:
// A: [1,2,3,2,1]
// B: [3,2,1,4,7]
// 输出: 3
// 解释: 
// 长度最长的公共子数组是 [3, 2, 1]。
// 说明:

// 1 <= len(A), len(B) <= 1000
// 0 <= A[i], B[i] < 100

package dynamic_programming

import (
	"fmt"
)

func findLength(A []int, B []int) int {
	var (
		lenA, lenB = len(A), len(B)
		dp = make([][]int, lenA)
		ret int
	)

	for i := 0; i < lenA; i++ {
		dp[i] = make([]int, lenB)
	}

	for i := 0; i < lenA; i++ {
		for j := 0; j < lenB; j++ {
			if A[i] == B[j] {
				dp[i][j] = 1
				if i > 0 && j > 0 {
					dp[i][j] += dp[i-1][j-1]
				}
				ret = max(ret, dp[i][j])
			}
		}
	}

	return ret
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
