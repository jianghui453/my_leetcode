package dynamic_programming

import (
	"fmt"
	"math"
)

func calculateMinimumHP0(dungeon [][]int) int {
	if len(dungeon) == 0 || len(dungeon[0]) == 0 {
		return 0
	}

	m, n := len(dungeon), len(dungeon[0])
	// 倒序动态规划。
	// dp 里记录每一个需要的最小生命值，生命值必须大于等于 1。
	// 通过公主所在位置后必须还有至少为 1 的生命值。
	// 如果是正序动态规划，每一步都有可能受到后面的影响，因为这里需要考虑两个地方：
	// 总的生命值以及过程中出现过的最小生命值。
	// 总生命值和最小生命值的选择可能是不一致的。
	// 后面可能需要高的总生命值，这时候就会导致前面的选择出错。
	// 如果选择总生命值高的路线，则会导致最小生命值不可控。
	// 如果使用倒序动态规划，只需要关系勇士的生命值，这一步所需要的最小生命值 + 上一步所需要的最小生命值。
	// 所以 dp[0][0] 就是勇士初始所需要的最小生命值。
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		// 这一步是避免在选择上一步所需要的最小生命值时出现 0 的情况。
		for j := 0; j <= n; j++ {
			dp[i][j] = math.MaxInt32
		}
	}

	dp[m][n-1], dp[m-1][n] = 1, 1

	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}

	for i := m-1; i >= 0; i-- {
		for j := n-1; j >= 0; j-- {
			// 选择上一步所需要的最小生命值
			minn := min(dp[i+1][j], dp[i][j+1])
			// 这一步所需要的最小生命值是上一步的最小生命值 + 这一步会消耗的生命值。
			// 最小生命值不能小于 1。
			dp[i][j] = max(minn - dungeon[i][j], 1)
		}
	}

	fmt.Println(dp)
	return dp[0][0]
}

func calculateMinimumHP1(dungeon [][]int) int {

	min := func (x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	max := func (x, y int) int {
		if x > y {
			return x
		}
		return y
	}
    n, m := len(dungeon), len(dungeon[0])
    dp := make([][]int, n + 1)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, m + 1)
        for j := 0; j < len(dp[i]); j++ {
            dp[i][j] = math.MaxInt32
        }
    }
    dp[n][m - 1], dp[n - 1][m] = 1, 1
    for i := n - 1; i >= 0; i-- {
        for j := m - 1; j >= 0; j-- {
            minn := min(dp[i+1][j], dp[i][j+1])
            dp[i][j] = max(minn - dungeon[i][j], 1)
        }
    }

	fmt.Println(dp)
    return dp[0][0]
}

