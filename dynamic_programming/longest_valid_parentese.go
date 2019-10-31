// 给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。

// 示例 1:

// 输入: "(()"
// 输出: 2
// 解释: 最长有效括号子串为 "()"
// 示例 2:

// 输入: ")()())"
// 输出: 4
// 解释: 最长有效括号子串为 "()()"

package dynamic_programming

import (
	"math"
	// "fmt"
)

func longestValidParentheses(s string) int {
	sLen := len(s)
	if sLen <= 1 {
		return 0
	}

	maxLen := 0.0
	dp := make([][]bool, sLen)
	for i := 0; i < sLen; i++ {
		dp[i] = make([]bool, sLen)
	}

	for i := 1; i < sLen; i++ {
		if s[i] == '(' {
			continue
		}

		for j := i-1; j >= 0; j-- {
			if (j == i-1 && s[j] == '(') || (dp[j+1][i-1] && s[j] == '(') {
				dp[j][i] = true
				for k := j-2; k >= 0; k-- {
					if dp[k][j-1] {
						dp[k][i] = true
						j = k
					}
				}
				maxLen = math.Max(maxLen, float64(i-j+1))
			}
		}
	}
	
	return int(maxLen)
}

// func longestValidParentheses(s string) int {
// 	var left, right, ret int
// 	for i := 0; i < len(s); i++ {
// 		if s[i] == '(' {
// 			left++
// 		} else {
// 			right++
// 		}
// 		if left == right {
// 			if left*2 > ret {
// 				ret = left * 2
// 			}
// 		} else if left <= right {
// 			left, right = 0, 0
// 		}
// 	}
// 	left, right = 0, 0
// 	for i := len(s) - 1; i >= 0; i-- {
// 		if s[i] == '(' {
// 			left++
// 		} else {
// 			right++
// 		}
// 		if left == right {
// 			if left*2 > ret {
// 				ret = left * 2
// 			}
// 		} else if right <= left {
// 			left, right = 0, 0
// 		}
// 		fmt.Printf("left = %d; right = %d\n", left, right)
// 	}
// 	return ret
// }

// func longestValidParentheses(s string) int {
// 	var st = []int{-1}
// 	var ret int
// 	for i := 0; i < len(s); i++ {
// 		if s[i] == '(' {
// 			st = append(st, i)
// 		} else {
// 			if len(st) == 1 {
// 				st = []int{i}
// 			} else {
// 				if i-st[len(st)-2] > ret {
// 					ret = i - st[len(st)-2]
// 				}
// 				st = st[:len(st)-1]
// 			}
// 		}
// 		fmt.Printf("st = %v\n", st)
// 	}
// 	return ret
// }

// func longestValidParentheses(s string) int {
// 	var dp = make([]int, len(s), len(s))
// 	var ret = 0
// 	for i := 1; i < len(s); i++ {
// 		if s[i] == '(' {
// 			continue
// 		}
// 		if s[i-1] == '(' {
// 			if i-2 >= 0 {
// 				dp[i] = dp[i-2] + 2
// 			} else {
// 				dp[i] = 2
// 			}
// 		} else if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
// 			dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-1]
// 			if i-dp[i-1]-2 > 0 {
// 				dp[i] += dp[i-dp[i-1]-2]
// 			}
// 		}
// 		if dp[i] > ret {
// 			ret = dp[i]
// 		}
// 	}
// 	fmt.Printf("db = %v\n", dp)
// 	return ret
// }
