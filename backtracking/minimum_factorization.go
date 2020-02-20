// 给定一个正整数 a，找出最小的正整数 b 使得 b 的所有数位相乘恰好等于 a。

// 如果不存在这样的结果或者结果不是 32 位有符号整数，返回 0。

//  

// 样例 1

// 输入：

// 48 
// 输出：

// 68
//  

// 样例 2

// 输入：

// 15
// 输出：

// 35

package backtracking

import (
	"sort"
	"math"
)

func smallestFactorization(a int) int {
	var res int
	
	ok, ret := bt(a)
	if ok {
		sort.Ints(ret)
		for i := range ret {
			res = res * 10 + ret[i]
			if res > math.MaxInt32 {
				return 0
			}
		}
		return res
	}

	return 0
}

func bt(a int) (bool, []int) {
	if a < 10 {
		return true, []int{a}
	}

	for i := 9; i > 1; i-- {
		if a % i == 0 {
			ok, ret := bt(a/i)
			if ok {
				ret = append(ret, i)
				return true, ret
			}
		}
	}

	return false, []int{}
}
