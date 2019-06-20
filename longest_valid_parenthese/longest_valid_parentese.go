package longest_parenthese

import (
	"fmt"
)

func longestValidParentheses(s string) int {
	var left, right, ret int
	for i := 0; i < len(s); i ++ {
		if s[i] == '(' {
			left ++
		} else {
			right ++
		}
		if left == right {
			if left * 2 > ret {
				ret = left * 2
			}
		} else if left <= right {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i -- {
		if s[i] == '(' {
			left ++
		} else {
			right ++
		}
		if left == right {
			if left * 2 > ret {
				ret = left * 2
			}
		} else if right <= left {
			left, right = 0, 0
		}
		fmt.Printf("left = %d; right = %d\n", left, right)
	}
	return ret
}

func longestValidParenthesesV1(s string) int {
	var st = []int{-1}
	var ret int
	for i := 0; i < len(s); i ++ {
		if s[i] == '(' {
			st = append(st, i)
		} else {
			if len(st) == 1 {
				st = []int{i}
			} else {
				if i - st[len(st) - 2] > ret {
					ret = i - st[len(st) - 2]
				}
				st = st[: len(st) - 1]
			}
		}
		fmt.Printf("st = %v\n", st)
	}
	return ret
}

func longestValidParenthesesV0(s string) int {
	var dp = make([]int, len(s), len(s))
	var ret = 0
	for i := 1; i < len(s); i ++ {
		if s[i] == '(' {
			continue
		}
		if s[i - 1] == '(' {
			if i - 2 >= 0 {
				dp[i] = dp[i - 2] + 2
			} else {
				dp[i] = 2
			}
		} else if i - dp[i - 1] - 1 >= 0 && s[i - dp[i - 1] - 1] == '(' {
			dp[i] = dp[i - 1] + 2 + dp[i - dp[i - 1] - 1]
			if i - dp[i - 1] - 2 > 0 {
				dp[i] += dp[i - dp[i - 1] - 2]
			}
		}
		if dp[i] > ret {
			ret = dp[i]
		}
	}
	fmt.Printf("db = %v\n", dp)
	return ret
}