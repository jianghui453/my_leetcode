//给定三个字符串 s1, s2, s3, 验证 s3 是否是由 s1 和 s2 交错组成的。
//
//示例 1:
//
//输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
//输出: true
//示例 2:
//
//输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
//输出: false

package dynamic_programming

import (
	// "fmt"
)

func isInterleave(s1 string, s2 string, s3 string) bool {
	l1, l2, l3 := len(s1), len(s2), len(s3)
	if l3 != l1 + l2 {
		return false
	}

	dp := make([][]bool, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]bool, l2+1)
	}
	dp[0][0] = true

	for i := 0; i <= l1; i++ {
		for j := 0; j <= l2; j++ {
			if i > 0 && dp[i-1][j] && s1[i-1] == s3[i+j-1] {
				dp[i][j] = true
			}
			if j > 0 && dp[i][j-1] && s2[j-1] == s3[i+j-1] {
				dp[i][j] = true
			}
		}
	}

	return dp[l1][l2]
}

// func isInterleave(s1 string, s2 string, s3 string) bool {
// 	len1 := len(s1)
// 	len2 := len(s2)
// 	len3 := len(s3)
// 	if len1+len2 != len3 {
// 		return false
// 	}

// 	dp := make([][]bool, len1+1)
// 	for i := 0; i < len1+1; i++ {
// 		dp[i] = make([]bool, len2+1)
// 	}
// 	dp[0][0] = true
// 	for i := 0; i < len1; i++ {
// 		if s1[i] == s3[i] {
// 			dp[i+1][0] = true
// 		} else {
// 			break
// 		}
// 	}
// 	for i := 0; i < len2; i++ {
// 		if s2[i] == s3[i] {
// 			dp[0][i+1] = true
// 		} else {
// 			break
// 		}
// 	}
// 	for i := 1; i <= len1; i++ {
// 		for j := 1; j <= len2; j++ {
// 			dp[i][j] = (i > 0 && dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (j > 0 && dp[i][j-1] && s2[j-1] == s3[i+j-1])
// 		}
// 	}
// 	return dp[len1][len2]
// }
