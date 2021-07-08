//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
//
//示例 1:
//
//输入: 123
//输出: 321
// 示例 2:
//
//输入: -123
//输出: -321
//示例 3:
//
//输入: 120
//输出: 21
//注意:
//
//假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
//
//标签：数学

package math_algorithm

import "math"

// import "fmt"

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	ret := 0
	signed := x / int(math.Abs(float64(x)))
	x = int(math.Abs(float64(x)))
	for x != 0 {
		reminder := x % 10
		ret = ret*10 + reminder
		if ret > math.MaxInt32 || ret < math.MinInt32 {
			return 0
		}
		x = x / 10
	}
	return ret * signed
}

// func reverse(x int) int {
// 	positive := x >= 0
// 	num := 0
// 	if !positive {
// 		x = -x
// 	}
// 	for i := 10; i <= x * 10; i *= 10 {
// 		num = num * 10 + (x % i) * 10 / i
// 		if (positive && num > 1<<31 - 1) || (!positive && num > 1<<31 - 1) {
// 			return 0
// 		}
// 	}
// 	if positive {
// 		return num
// 	}
// 	return 0 - num
// }

//func reverse(x int) int {
//	var num int
//	for x != 0 {
//		num = num * 10 + (x % 10)
//		if num < math.MinInt32 || num > math.MaxInt32{
//			return 0
//		}
//		x = x / 10
//	}
//	return num
//}
