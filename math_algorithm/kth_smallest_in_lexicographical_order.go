// 给定整数 n 和 k，找到 1 到 n 中字典序第 k 个的数字。

// 注意：1 ≤ k ≤ n ≤ 10^9。

// 示例 :

// 输入:
// n: 13   k: 2

// 输出:
// 10

// 解释:
// 字典序的排列是 [1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9]，所以第二个数字是 10。

package math_algorithm

import (
	// "fmt"
	// "math"
)

func findKthNumber(n int, k int) int {
	num := 1
	for k > 1 {
		cnt := calCnt(num, n)
		
		if k > cnt {
			num++
			k -= cnt
		} else {
			num *= 10
			k--
		}
	}

	return num
}

func calCnt(num, n int) int {
	magnitudes := 1
	var cnt int
	for num*magnitudes <= n {
		if (num+1)*magnitudes > n {
			return cnt + n - num*magnitudes + 1
		}
		cnt += magnitudes
		magnitudes *= 10
	}
	return cnt
}