// 给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

// 返回被除数 dividend 除以除数 divisor 得到的商。

// 示例 1:

// 输入: dividend = 10, divisor = 3
// 输出: 3
// 示例 2:

// 输入: dividend = 7, divisor = -3
// 输出: -2
// 说明:

// 被除数和除数均为 32 位有符号整数。
// 除数不为 0。
// 假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,  231 − 1]。本题中，如果除法结果溢出，则返回 231 − 1。

package divide

import (
	"math"
	"fmt"
)

func divide(dividend int, divisor int) int {
	var _divide func (_dividend, _divisor, noun int) int
	_divide = func (_dividend, _divisor, noun int) int {
		fmt.Printf("_dividend=%d _divisor=%d noun=%d\n", _dividend, _divisor, noun)

		if math.Abs(float64(_dividend)) < math.Abs(float64(_divisor)) {
			return 0
		}

		sign := (_dividend < 0 && _divisor < 0) || (_dividend > 0 && _divisor > 0)

		_divisor += _divisor
		if sign {
			_dividend -= _divisor
		} else {
			_dividend += _divisor
		}
		noun += noun
		
		ret := _divide(_dividend-_divisor, _divisor, noun)
		if ret > math.MaxInt32 || ret < math.MinInt32 {
			return math.MaxInt32
		}

		if ret == 0 {
			
		}

		if sign {
			return ret+noun
		} else {
			return ret-noun
		}
	}

	ret := _divide(dividend, divisor, 1)
	if ret > math.MaxInt32 || ret < math.MinInt32 {
		return math.MaxInt32
	}

	return ret
}

// func divide(dividend int, divisor int) int {
// 	signBool := (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0)
// 	if dividend < 0 {
// 		dividend = 0 - dividend
// 	}
// 	if divisor < 0 {
// 		divisor = 0 - divisor
// 	}
// 	ret := 0
// 	for dividend >= divisor {
// 		var logn uint = 0
// 		for {
// 			if dividend < (divisor << logn) {
// 				ret += int(1 << (logn - 1))
// 				dividend = dividend - (divisor << (logn - 1))
// 				break
// 			}
// 			logn++
// 		}
// 	}
// 	if !signBool {
// 		ret = 0 - ret
// 	}
// 	if ret > math.MaxInt32 {
// 		return math.MaxInt32
// 	}
// 	if ret < math.MinInt32 {
// 		return math.MinInt32
// 	}
// 	return ret
// }
