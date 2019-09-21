//给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
//
//返回 s 所有可能的分割方案。
//
//示例:
//
//输入: "aab"
//输出:
//[
//  ["aa","b"],
//  ["a","a","b"]
//]

package backtracking

func partition(s string) [][]string {
    result := [][]string{}
    current := make([]string, 0, len(s))
    dfs(s, 0, current, &result)
    return result
}

func dfs(s string, i int, cur []string, result *[][]string) {
    if i == len(s) {
        tmp := make([]string, len(cur))
        copy(tmp, cur)
        *result = append(*result, tmp)
        return
    }

    for j := i; j < len(s); j++ {
        // i == 0 时，
        // 按照 len(cur[0]) 的不同，来划分 res
        // 并以此类推
        if par(s[i : j+1]) {
            dfs(s, j+1, append(cur, s[i:j+1]), result)
        }
    }
}

// s 为回文，则返回 true
func par(s string) bool {
    if len(s) <= 1 {
        return true
    }
    a, b := 0, len(s)-1
    for a < b {
        if s[a] != s[b] {
            return false
        }
        a++
        b--
    }
    return true
}

//func partition(s string) [][]string {
//    sLen := len(s)
//    ret := make([][]string, 0)
//    if sLen == 0 {
//        return ret
//    }
//
//    dp := make([][]bool, sLen)
//    for i := 0; i < sLen; i ++ {
//        dp[i] = make([]bool, sLen)
//    }
//
//    for i := sLen-1; i >= 0; i -- {
//        j := i
//        for ; j < sLen; j ++ {
//            if s[j] == s[i] {
//                dp[i][j] = true
//            } else {
//                break
//            }
//        }
//        for ; j < sLen; j ++ {
//            if s[j] == s[i] && (i+1 < sLen && j-1 >= 0 && i+1 <= j-1 && dp[i+1][j-1]) {
//                dp[i][j] = true
//            }
//        }
//    }
//
//    var f func(his []string, idx int)
//    f = func(his []string, idx int) {
//        lenHis := len(his)
//        for j := idx; j < sLen; j ++ {
//            if dp[idx][j] {
//                newHis := make([]string, lenHis)
//                copy(newHis, his)
//                if j < sLen-1 {
//                    newHis = append(newHis, s[idx: j+1])
//                    f(newHis, j+1)
//                } else {
//                    newHis = append(newHis, s[idx: ])
//                    ret = append(ret, newHis)
//                }
//            }
//        }
//    }
//    f([]string{}, 0)
//
//    return ret
//}
