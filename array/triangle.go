//给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
//
//例如，给定三角形：
//
//[
//     [2],
//    [3,4],
//   [6,5,7],
//  [4,1,8,3]
//]
//自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
//
//说明：
//
//如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。

package array

import (
	"math"
)

func minimumTotal(triangle [][]int) int {
	triLen := len(triangle)
	if triLen == 0 {
		return 0
	}
	if triLen == 1 {
		return triangle[0][0]
	}

	ret := math.MaxInt64
	for i := 1; i < triLen; i++ {
		triangle[i][0] += triangle[i-1][0]
		for j := 1; j < i; j++ {
			triangle[i][j] += int(math.Min(float64(triangle[i-1][j-1]), float64(triangle[i-1][j])))
		}
		triangle[i][i] += triangle[i-1][i-1]
	}

	for i := 0; i < triLen; i++ {
		ret = int(math.Min(float64(triangle[triLen-1][i]), float64(ret)))
	}
	return ret
}

// func minimumTotal(triangle [][]int) int {
// 	lt := len(triangle)
// 	if lt == 0 {
// 		return 0
// 	}
// 	r := []int{triangle[0][0]}
// 	lr := 1
// 	for i := 1; i < lt; i++ {
// 		rn := make([]int, lr+1)
// 		for j := 0; j <= lr; j++ {
// 			if j == 0 {
// 				rn[j] = triangle[i][j] + r[j]
// 			} else if j < lr {
// 				if r[j-1] < r[j] {
// 					rn[j] = triangle[i][j] + r[j-1]
// 				} else {
// 					rn[j] = triangle[i][j] + r[j]
// 				}
// 			} else {
// 				rn[j] = triangle[i][j] + r[j-1]
// 			}
// 		}
// 		r = rn
// 		lr++
// 	}
// 	min := math.MaxInt64
// 	for i := 0; i < lr; i++ {
// 		if r[i] < min {
// 			min = r[i]
// 		}
// 	}
// 	return min
// }
