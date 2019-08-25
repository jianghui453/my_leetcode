//给定两个单词 word1 和 word2，计算出将 word1 转换成 word2 所使用的最少操作数 。
//
//你可以对一个单词进行如下三种操作：
//
//插入一个字符
//删除一个字符
//替换一个字符
//示例 1:
//
//输入: word1 = "horse", word2 = "ros"
//输出: 3
//解释:
//horse -> rorse (将 'h' 替换为 'r')
//rorse -> rose (删除 'r')
//rose -> ros (删除 'e')
//示例 2:
//
//输入: word1 = "intention", word2 = "execution"
//输出: 5
//解释:
//intention -> inention (删除 't')
//inention -> enention (将 'i' 替换为 'e')
//enention -> exention (将 'n' 替换为 'x')
//exention -> exection (将 'n' 替换为 'c')
//exection -> execution (插入 'u')

package min_distance

import (
    "fmt"
    "math"
)

func minDistance(word1 string, word2 string) int {
    lenWord1 := len(word1)
    lenWord2 := len(word2)
    if lenWord1 == 0 {
        return lenWord2
    }
    if lenWord2 == 0 {
        return lenWord1
    }
    dp := make([][]int, lenWord1)
    for i := 0; i < lenWord1; i ++ {
        dp[i] = make([]int, lenWord2)
    }
    if word1[0] != word2[0] {
        dp[0][0] = 1
        for i := 1; i < lenWord1; i ++ {
            if word1[i] == word2[0] {
                dp[i][0] = dp[i-1][0]
                i ++
                for ; i < lenWord1; i ++ {
                    dp[i][0] = dp[i-1][0]+1

                }
            } else {
                dp[i][0] = dp[i-1][0]+1
            }
        }
        for j := 1; j < lenWord2; j ++ {
            if word2[j] == word1[0] {
                dp[0][j] = dp[0][j-1]
                j ++
                for ; j < lenWord2; j ++ {
                    dp[0][j] = dp[0][j-1]+1

                }
            } else {
                dp[0][j] = dp[0][j-1]+1
            }
        }
    } else {
        for i := 1; i < lenWord1; i ++ {
            dp[i][0] = dp[i-1][0]+1
        }
        for j := 1; j < lenWord2; j ++ {
            dp[0][j] = dp[0][j-1]+1
        }
    }

    for i := 1; i < lenWord1; i ++ {
        for j := 1; j < lenWord2; j ++ {
            min := int(math.Min(float64(dp[i-1][j]+1), float64(dp[i][j-1]+1)))
            if word1[i] == word2[j] {
                if dp[i-1][j-1] < min {
                    min = dp[i-1][j-1]
                }
            } else {
                if dp[i-1][j-1]+1 < min {
                    min = dp[i-1][j-1]+1
                }
            }
            dp[i][j] = min
        }
    }
    fmt.Printf("dp=%v\n", dp)
    return dp[lenWord1-1][lenWord2-1]
}