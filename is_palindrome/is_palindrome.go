package leet_code

import (
	"fmt"
)

func isPalindrome(x int) bool {
	_x := x
	if x >= 0 {
		ret := 0
		for {
			ret = ret*10 + (x % 10)
			fmt.Printf("ret = %d\n", ret)
			if x < 10 {
				break
			}
			x = x / 10
		}
		if ret == _x {
			return true
		}
	}
	return false
}
