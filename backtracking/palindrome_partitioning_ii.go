//给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
//
//返回符合要求的最少分割次数。
//
//示例:
//
//输入: "aab"
//输出: 1
//解释: 进行一次分割就可将 s 分割成 ["aa","b"] 这样两个回文子串。

package backtracking

import (
	"math"
)

func minCut(s string) int {
	memo := make(map[string]int)
	var backtracking func(str string) int
	backtracking = func(str string) int {
		l := len(str)
		if l <= 1 {
			return 0
		}
		if _, ok := memo[str]; ok {
			return memo[str]
		}
		if isParlindrome(str) {
			memo[str] = 0
			return memo[str]
		}

		ret := math.MaxInt64-1
		for i := l-1; i > 0; i-- {
			if isParlindrome(str[: i]) {
				ret = int(math.Min(float64(ret), float64(backtracking(str[i: ])+1)))
			}
		}
		memo[str] = ret
		return ret
	}

	r := backtracking(s)
	return r
}

func isParlindrome(s string) bool {
	l := len(s)
	left, right := 0, l-1
	for left <= right {
		if s[left] != s[right] {
			return false
		}
		left, right = left+1, right-1
	}
	return true
}

// func minCut(s string) int {
// 	lenS := len(s)
// 	if lenS < 2 {
// 		return 0
// 	}

// 	dp := make([][]bool, lenS)
// 	for i := 0; i < lenS; i++ {
// 		dp[i] = make([]bool, lenS)
// 	}

// 	for i := lenS - 1; i >= 0; i-- {
// 		j := i
// 		for ; j < lenS; j++ {
// 			if s[j] == s[i] {
// 				dp[i][j] = true
// 			} else {
// 				break
// 			}
// 		}
// 		for ; j < lenS; j++ {
// 			if s[j] == s[i] && (i+1 < lenS && j-1 >= 0 && i+1 <= j-1 && dp[i+1][j-1]) {
// 				dp[i][j] = true
// 			}
// 		}
// 	}

// 	record := make([]bool, lenS)
// 	his := make([][]int, 0)
// 	for i := lenS - 1; i >= 0; i-- {
// 		if dp[0][i] {
// 			his = append(his, []int{i})
// 			record[i] = true
// 			if i == lenS-1 {
// 				return 0
// 			}
// 		}
// 	}

// 	level := 1
// 	for len(his) > 0 {
// 		lenHis := len(his)
// 		newHis := make([][]int, 0)
// 		for i := 0; i < lenHis; i++ {
// 			idx := his[i][level-1]
// 			for j := lenS - 1; j > idx; j-- {
// 				if record[j] {
// 					continue
// 				}
// 				if dp[idx+1][j] {
// 					hisItem := make([]int, level)
// 					copy(hisItem, his[i])
// 					hisItem = append(hisItem, j)
// 					record[j] = true
// 					newHis = append(newHis, hisItem)
// 					if j == lenS-1 {
// 						return level
// 					}
// 				}
// 			}
// 		}
// 		level++
// 		his = newHis
// 	}

// 	return level
// }
