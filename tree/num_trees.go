package tree

func numTrees(n int) int {
    if n < 2 {
        return 1
    }
    dp := make([]int, n+1)
    dp[0] = 1
    dp[1] = 1
    for i := 2; i <= n; i ++ {
        for j := 0; j < i; j ++ {
            dp[i] += dp[j]*dp[i-1-j]
        }
    }
    return dp[n]
}
