//给定一个字符串 S 和一个字符串 T，计算在 S 的子序列中 T 出现的个数。
//
//一个字符串的一个子序列是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。（例如，"ACE" 是 "ABCDE" 的一个子序列，而 "AEC" 不是）
//
//示例 1:
//
//输入: S = "rabbbit", T = "rabbit"
//输出: 3
//解释:
//
//如下图所示, 有 3 种可以从 S 中得到 "rabbit" 的方案。
//(上箭头符号 ^ 表示选取的字母)
//
//rabbbit
//^^^^ ^^
//rabbbit
//^^ ^^^^
//rabbbit
//^^^ ^^^
//示例 2:
//
//输入: S = "babgbag", T = "bag"
//输出: 5
//解释:
//
//如下图所示, 有 5 种可以从 S 中得到 "bag" 的方案。
//(上箭头符号 ^ 表示选取的字母)
//
//babgbag
//^^ ^
//babgbag
//^^    ^
//babgbag
//^    ^^
//babgbag
//  ^  ^^
//babgbag

package dynamic_programming

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

func numDistinct(s string, t string) int {
	len_s, len_t := len(s), len(t)
	dp := make([][]int, len_s+1)
	for i := 0; i <= len_s; i++ {
		dp[i] = make([]int, len_t+1)
	}
	dp[0][0] = 1
	for i := 1; i < len_s; i++ {
		dp[i][0] = 1
	}
	for j := 1; j < len_t; j++ {
		dp[0][j] = 0
	}

	for i := 1; i <= len_s; i++ {
		for j := 1; j <= len_t; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len_s][len_t]
}
