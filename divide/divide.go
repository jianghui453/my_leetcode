package divide

import (
	"math"
)

func divide(dividend int, divisor int) int {
	signBool := (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0)
	if dividend < 0 {
		dividend = 0 - dividend
	}
	if divisor < 0 {
		divisor = 0 - divisor
	}
	ret := 0
	for dividend >= divisor {
		var logn uint = 0
		for {
			if dividend < (divisor << logn) {
				ret += int(1 << (logn - 1))
				dividend = dividend - (divisor << (logn - 1))
				break
			}
			logn++
		}
	}
	if !signBool {
		ret = 0 - ret
	}
	if ret > math.MaxInt32 {
		return math.MaxInt32
	}
	if ret < math.MinInt32 {
		return math.MinInt32
	}
	return ret
}
