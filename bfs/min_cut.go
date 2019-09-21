//给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
//
//返回符合要求的最少分割次数。
//
//示例:
//
//输入: "aab"
//输出: 1
//解释: 进行一次分割就可将 s 分割成 ["aa","b"] 这样两个回文子串。

package bfs

func minCut(s string) int {
	lenS := len(s)
	if lenS < 2 {
		return 0
	}

	dp := make([][]bool, lenS)
	for i := 0; i < lenS; i++ {
		dp[i] = make([]bool, lenS)
	}

	for i := lenS - 1; i >= 0; i-- {
		j := i
		for ; j < lenS; j++ {
			if s[j] == s[i] {
				dp[i][j] = true
			} else {
				break
			}
		}
		for ; j < lenS; j++ {
			if s[j] == s[i] && (i+1 < lenS && j-1 >= 0 && i+1 <= j-1 && dp[i+1][j-1]) {
				dp[i][j] = true
			}
		}
	}

	record := make([]bool, lenS)
	his := make([][]int, 0)
	for i := lenS - 1; i >= 0; i-- {
		if dp[0][i] {
			his = append(his, []int{i})
			record[i] = true
			if i == lenS-1 {
				return 0
			}
		}
	}

	level := 1
	for len(his) > 0 {
		lenHis := len(his)
		newHis := make([][]int, 0)
		for i := 0; i < lenHis; i++ {
			idx := his[i][level-1]
			for j := lenS - 1; j > idx; j-- {
				if record[j] {
					continue
				}
				if dp[idx+1][j] {
					hisItem := make([]int, level)
					copy(hisItem, his[i])
					hisItem = append(hisItem, j)
					record[j] = true
					newHis = append(newHis, hisItem)
					if j == lenS-1 {
						return level
					}
				}
			}
		}
		level++
		his = newHis
	}

	return level
}
