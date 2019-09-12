//给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
//
//'.' 匹配任意单个字符
//'*' 匹配零个或多个前面的那一个元素
//所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
//
//说明:
//
//s 可能为空，且只包含从 a-z 的小写字母。
//p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
//示例 1:
//
//输入:
//s = "aa"
//p = "a"
//输出: false
//解释: "a" 无法匹配 "aa" 整个字符串。
//示例 2:
//
//输入:
//s = "aa"
//p = "a*"
//输出: true
//解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
//示例 3:
//
//输入:
//s = "ab"
//p = ".*"
//输出: true
//解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
//示例 4:
//
//输入:
//s = "aab"
//p = "c*a*b"
//输出: true
//解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
//示例 5:
//
//输入:
//s = "mississippi"
//p = "mis*is*p*."
//输出: false

package dynamic_programming

import "fmt"

func isMatch(s string, p string) bool {
	lenS := len(s)
	lenP := len(p)
	if lenS == 0 {
		cnt := 0
		for i := 0; i < lenP; i++ {
			if p[i] != '*' {
				cnt++
			} else if i > 0 && p[i-1] != '*' {
				cnt--
			}
		}
		return cnt == 0
	}
	if lenP == 0 {
		return false
	}
	dp := make([][]bool, lenS)
	for i := 0; i < lenS; i++ {
		dp[i] = make([]bool, lenP)
	}
	cnt := 0
	for i := 0; i < lenP; i++ {
		if (cnt == 0 && (p[i] == s[0] || p[i] == '.')) || ((i == 0 || dp[0][i-1]) && p[i] == '*') {
			dp[0][i] = true
		}
		if p[i] != '*' {
			cnt++
		} else if i > 0 && p[i-1] != '*' {
			cnt--
		}
	}
	for i := 1; i < lenP; i++ {
		for j := 0; j < lenS; j++ {
			if (j > 0 && dp[j-1][i-1] && (s[j] == p[i] || p[i] == '.' || (p[i] == '*' && (p[i-1] == s[j] || p[i-1] == '.')))) ||
				(j > 0 && dp[j-1][i] && (p[i] == '*' && (s[j] == p[i-1] || p[i-1] == '.'))) ||
				(dp[j][i-1] && (p[i] == '*')) ||
				(i > 1 && dp[j][i-2] && (p[i] == '*')) {
				dp[j][i] = true
			}
		}
	}
	fmt.Printf("dp=%v\n", dp)
	return dp[lenS-1][lenP-1]
}

//var memo = make(map[string]map[string]bool)
//
//func isMatch(s string, p string) bool {
//    if ret, ok := memo[s][p]; ok {
//        return ret
//    }
//    fmt.Printf("s = %s; p = %s; memo = %v\n", s, p, memo)
//    memo[s] = make(map[string]bool)
//    if len(p) == 0 {
//        memo[s][p] = len(s) == 0
//        return memo[s][p]
//    }
//    if len(s) == 0 {
//        memo[s][p] = (len(p) == 2 && p[1] == '*') || (len(p) > 2 && p[1] == '*' && isMatch("", p[2:]))
//        return memo[s][p]
//    }
//    // len(p) > 0 && len(s) > 0
//    firstMatch := p[0] == s[0] || p[0] == '.'
//    var s1, p1, p2 string
//    if len(s) == 1 {
//        s1 = ""
//    } else {
//        s1 = s[1:]
//    }
//    if len(p) == 1 {
//        p1 = ""
//    } else {
//        p1 = p[1:]
//    }
//    if len(p) >= 2 && p[1] == '*' {
//        if len(p) == 2 {
//            p2 = ""
//        } else {
//            p2 = p[2:]
//        }
//        memo[s][p] = isMatch(s, p2) || (firstMatch && isMatch(s1, p2)) || (firstMatch && isMatch(s1, p))
//        return  memo[s][p]
//    }
//    memo[s][p] = firstMatch && isMatch(s1, p1)
//    return memo[s][p]
//}
