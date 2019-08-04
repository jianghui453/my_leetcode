//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

package leet_code

import "math"

func reverse(x int) int {
	positive := x >= 0
	num := 0
	//max := 1<<31 - 1
	if !positive {
		x = -x
	}
	for i := 10; i <= x * 10; i *= 10 {
		num = num * 10 + (x % i) * 10 / i
		if (positive && num > 1<<31 - 1) || (!positive && num > 1<<31 - 1) {
			return 0
		}
	}
	//fmt.Printf("num=%d; positive=%v; max=%d\n", num, positive, max)
	if positive {
		return num
	}
	return 0 - num
}

func reverseV2(x int) int {
	var num int
	for x != 0 {
		num = num * 10 + (x % 10)
		if num < math.MinInt32 || num > math.MaxInt32{
			return 0
		}
		x = x / 10
	}
	return num
}
