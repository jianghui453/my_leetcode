//实现 int sqrt(int x) 函数。
//
//计算并返回 x 的平方根，其中 x 是非负整数。
//
//由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
//
//示例 1:
//
//输入: 4
//输出: 2
//示例 2:
//
//输入: 8
//输出: 2
//说明: 8 的平方根是 2.82842...,
//     由于返回类型是整数，小数部分将被舍去。

package binary_search

import (
	// "math"
)

func mySqrt(x int) int {
	min, max := 0, x
	for min <= max {
		mid := (min+max)/2
		if mid * mid > x {
			max = mid-1
		} else if mid * mid < x {
			min = mid+1
		} else {
			return mid
		}
	}
	return max
}

// func mySqrt(x int) int {
// 	if x == 0 {
// 		return 0
// 	}
// 	if x < 4 {
// 		return 1
// 	}
// 	y := x / 2
// 	yNext := (y + x/y) / 2
// 	for y*y > x || math.Abs(float64(y-yNext)) >= 1 {
// 		if math.Abs(float64(y-yNext)) == 1 && y*y < x || yNext*yNext < x {
// 			return int(math.Min(float64(y), float64(yNext)))
// 		}
// 		y = yNext
// 		yNext = (y + x/y) / 2
// 	}
// 	return y
// }
