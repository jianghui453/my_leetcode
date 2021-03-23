//你将获得 K 个鸡蛋，并可以使用一栋从 1 到 N  共有 N 层楼的建筑。
//
//每个蛋的功能都是一样的，如果一个蛋碎了，你就不能再把它掉下去。
//
//你知道存在楼层 F ，满足 0 <= F <= N 任何从高于 F 的楼层落下的鸡蛋都会碎，从 F 楼层或比它低的楼层落下的鸡蛋都不会破。
//
//每次移动，你可以取一个鸡蛋（如果你有完整的鸡蛋）并把它从任一楼层 X 扔下（满足 1 <= X <= N）。
//
//你的目标是确切地知道 F 的值是多少。
//
//无论 F 的初始值如何，你确定 F 的值的最小移动次数是多少？
//
//
//
//示例 1：
//
//输入：K = 1, N = 2
//输出：2
//解释：
//鸡蛋从 1 楼掉落。如果它碎了，我们肯定知道 F = 0 。
//否则，鸡蛋从 2 楼掉落。如果它碎了，我们肯定知道 F = 1 。
//如果它没碎，那么我们肯定知道 F = 2 。
//因此，在最坏的情况下我们需要移动 2 次以确定 F 是多少。
//示例 2：
//
//输入：K = 2, N = 6
//输出：3
//示例 3：
//
//输入：K = 3, N = 14
//输出：4
//
//
//提示：
//
//1 <= K <= 100
//1 <= N <= 10000

package dynamic_programming

import (
	// "fmt"
)

func superEggDrop(K int, N int) int {
	dp := make([][]int, K+1)
	for k := 0; k <= K; k++ {
		dp[k] = make([]int, N+1)
		for n := 1; n <= N; n++ {
			dp[k][n] = n
		}
	}
	
	for n := 2; n <= N; n++ {
		for k := 2; k <= K; k++ {
			var (
				l, h, m, t1, t2 int
			)

			l, h = 0, n
			for l <= h {
				m = (l+h)/2
				// 两种情况：鸡蛋碎掉和没碎，取其中最大的一种情况（因为都是考虑最坏的情况）
				// 因为这两种情况一个是递增一个是递减，最大情况中的最小值就是它们的相交处，这里可以用二分法来优化查找过程
				t1, t2 = dp[k][n-m], dp[k-1][m-1] // 如果鸡蛋碎了，这一层就不需要继续查找了，如果没碎则需要继续查找
				if t1 > t2 {
					l = m + 1
				} else if t1 < t2 {
					h = m - 1
				} else {
					l, h = m, m
					break
				}
			}
			
			dp[k][n] = min(max(dp[k][n-h], dp[k-1][h-1]), max(dp[k][n-l], dp[k-1][l-1])) + 1
		}
	}
	
	return dp[K][N]
}
