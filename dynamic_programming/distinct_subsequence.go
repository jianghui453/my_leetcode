// 给定一个字符串 S 和一个字符串 T，计算在 S 的子序列中 T 出现的个数。

// 一个字符串的一个子序列是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。（例如，"ACE" 是 "ABCDE" 的一个子序列，而 "AEC" 不是）

// 示例 1:

// 输入: S = "rabbbit", T = "rabbit"
// 输出: 3
// 解释:

// 如下图所示, 有 3 种可以从 S 中得到 "rabbit" 的方案。
// (上箭头符号 ^ 表示选取的字母)

// rabbbit
// ^^^^ ^^
// rabbbit
// ^^ ^^^^
// rabbbit
// ^^^ ^^^
// 示例 2:

// 输入: S = "babgbag", T = "bag"
// 输出: 5
// 解释:

// 如下图所示, 有 5 种可以从 S 中得到 "bag" 的方案。
// (上箭头符号 ^ 表示选取的字母)

// babgbag
// ^^ ^
// babgbag
// ^^    ^
// babgbag
// ^    ^^
// babgbag
//  ^  ^^
// babgbag

package dynamic_programming

import (
	"fmt"
)

func numDistinct(s string, t string) int {
	sLen := len(s)
	tLen := len(t)
	if sLen < tLen || tLen == 0 {
		return 0
	}

	dp := make([][]int, sLen+1)
	for i := 0; i <= sLen; i++ {
		dp[i] = make([]int, tLen+1)
	}

	for j := 1; j <= tLen; j++ {
		for i := j; i <= sLen; i++ {
			dp[i][j] = dp[i-1][j]
			if s[i-1] == t[j-1] {
				if j == 1 {
					dp[i][j]++
				} else {
					dp[i][j] += dp[i-1][j-1]
				}
			}
		}
	}

	return dp[sLen][tLen]
}

//func numDistinct(s string, t string) int {
//    lenS := len(s)
//    lenT := len(t)
//    if lenS < lenT {
//        return 0
//    }
//    if lenT == 0 {
//        return 1
//    }
//    dp := make([][]int, lenS+1)
//    for i := 0; i <= lenS; i ++ {
//        dp[i] = make([]int, lenT+1)
//        dp[i][0] = 1
//        if i > 0 {
//            if i < lenS-1 {
//                dp[i][1] = strings.Count(s[: i+1], string(t[0]))
//            } else {
//                dp[i][1] = strings.Count(s, string(t[0]))
//            }
//        }
//    }
//
//    for j := 1; j <= lenT; j ++ {
//        for i := j; i <= lenS; i ++ {
//            dp[i][j] = dp[i-1][j]
//            if s[i-1] == t[j-1] {
//                dp[i][j] += dp[i-1][j-1]
//            }
//        }
//    }
//    return dp[lenS][lenT]
//}
