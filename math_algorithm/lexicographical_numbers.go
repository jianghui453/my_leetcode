// 给定一个整数 n, 返回从 1 到 n 的字典顺序。

// 例如，

// 给定 n =1 3，返回 [1,10,11,12,13,2,3,4,5,6,7,8,9] 。

// 请尽可能的优化算法的时间复杂度和空间复杂度。 输入的数据 n 小于等于 5,000,000。

package math_algorithm

import (
	// "fmt"
)

func lexicalOrder(n int) []int {
	ret := make([]int, 0)
	num := 1
	for {
		if num <= n {
			ret = append(ret, num)
			num *= 10
		} else {
			num /= 10
			for num % 10 == 9 {
				num /= 10
			}
			if num == 0 {
				break
			}
			num++
		}
	}

	return ret
}