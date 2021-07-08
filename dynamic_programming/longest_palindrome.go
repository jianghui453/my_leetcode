//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
//
//示例 1：
//
//输入: "babad"
//输出: "bab"
//注意: "aba" 也是一个有效答案。
//示例 2：
//
//输入: "cbbd"
//输出: "bb"

//标签：字符串，动态规划

package dynamic_programming

// import "fmt"

func longestPalindrome(s string) string {
	sLen := len(s)
	if sLen < 2 {
		return s
	}
	dp := make([][]bool, sLen)
	for i := 0; i < sLen; i++ {
		dp[i] = make([]bool, sLen)
		dp[i][i] = true
	}
	min := 0
	max := 0
	for i := 1; i < sLen; i++ {
		j := i - 1
		for ; j >= 0; j-- {
			if s[j] == s[i] {
				dp[j][i] = true
				if i-j > max-min {
					min, max = j, i
				}
			} else {
				break
			}
		}
		for ; j >= 0; j-- {
			if dp[j+1][i-1] && s[j] == s[i] {
				dp[j][i] = true
				if i-j > max-min {
					min, max = j, i
				}
			}
		}
	}
	if max == sLen-1 {
		return s[min:]
	}
	return s[min : max+1]
}

// func longestPalindrome(s string) string {
// 	lenS := len(s)
// 	if lenS < 2 {
// 		return s
// 	}
// 	dp := make([][]bool, lenS)
// 	for i := 0; i < lenS; i++ {
// 		dp[i] = make([]bool, lenS)
// 		dp[i][i] = true
// 	}
// 	maxLeft := 0
// 	maxRight := 0
// 	for i := 1; i < lenS; i++ {
// 		if s[i] == s[i-1] {
// 			dp[i-1][i] = true
// 			if maxRight-maxLeft < 1 {
// 				maxRight = i
// 				maxLeft = i - 1
// 			}
// 		}
// 	loop:
// 		for j := i - 2; j >= 0; j-- {
// 			if dp[j+1][i-1] == true && s[j] == s[i] {
// 				dp[j][i] = true
// 				if i-j > maxRight-maxLeft {
// 					maxRight = i
// 					maxLeft = j
// 				}
// 				continue
// 			}
// 			for k := j; k < i; k++ {
// 				if s[k] != s[i] {
// 					continue loop
// 				}
// 			}
// 			dp[j][i] = true
// 			if i-j > maxRight-maxLeft {
// 				maxRight = i
// 				maxLeft = j
// 			}
// 		}
// 	}
// 	if maxRight == lenS-1 {
// 		return s[maxLeft:]
// 	}
// 	return s[maxLeft : maxRight+1]
// }

//func longestPalindrome(s string) string {
//	ret := ""
//	strLen := len(s)
//	start := 0
//	for {
//		end := start
//		for {
//			if end < strLen && s[end] == s[start] {
//				end++
//				continue
//			}
//			end--
//			break
//		}
//		i := 1
//		for {
//			if start-i >= 0 && end+i < strLen && s[start-i] == s[end+i] {
//				i++
//				continue
//			}
//			i--
//			break
//		}
//		if end+i-(start-i)+1 > len(ret) {
//			ret = s[start-i : end+i+1]
//		}
//		if end+1 < strLen {
//			start = end + 1
//			continue
//		}
//		break
//	}
//	return ret
//}
