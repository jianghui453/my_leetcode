//格雷编码是一个二进制数字系统，在该系统中，两个连续的数值仅有一个位数的差异。
//
//给定一个代表编码总位数的非负整数 n，打印其格雷编码序列。格雷编码序列必须以 0 开头。
//
//示例 1:
//
//输入: 2
//输出: [0,1,3,2]
//解释:
//00 - 0
//01 - 1
//11 - 3
//10 - 2
//
//对于给定的 n，其格雷编码序列并不唯一。
//例如，[0,2,3,1] 也是一个有效的格雷编码序列。
//
//00 - 0
//10 - 2
//11 - 3
//01 - 1
//示例 2:
//
//输入: 0
//输出: [0]
//解释: 我们定义格雷编码序列必须以 0 开头。
//     给定编码总位数为 n 的格雷编码序列，其长度为 2n。当 n = 0 时，长度为 20 = 1。
//     因此，当 n = 0 时，其格雷编码序列为 [0]。

package backtracking

import (
// "math"
// "fmt"
)

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}

	if n == 1 {
		return []int{0, 1}
	}

	ret := make([]int, 0)
	v := 1 << uint(n-1)
	_ret := grayCode(n - 1)
	ret = append(ret, _ret...)

	for i := len(_ret) - 1; i >= 0; i-- {
		ret = append(ret, _ret[i]+v)
	}

	return ret
}

// func grayCode(n int) []int {
// 	if n < 1 {
// 		return []int{0}
// 	}
// 	ret := []int{0}
// 	his := make([]int, 1<<uint(n))
// 	num := 0
// 	his[0]++
// 	for i := 0; i < n; i++ {
// 		j := 1 << uint(i)
// 		fmt.Printf("i=%d j=%d num=%d\n", i, j, num)
// 		if num&j > 0 && his[num-j] == 0 {
// 			ret = append(ret, num-j)
// 			his[num-j]++
// 			i = -1
// 			num = num - j
// 		} else if num&j == 0 && his[num+j] == 0 {
// 			ret = append(ret, num+j)
// 			his[num+j]++
// 			i = -1
// 			num += j
// 		}
// 		fmt.Printf("i=%d j=%d num=%d his=%v ret=%v\n", i, j, num, his, ret)
// 	}
// 	return ret
// }
