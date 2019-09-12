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

//func superEggDrop(K int, N int) int {
//   dp := make([][]int, K+1)
//   for k := 0; k <= K; k ++ {
//       dp[k] = make([]int, N+1)
//       for n := 0; n <= N; n ++ {
//           dp[k][n] = n
//       }
//   }
//   for k := 2; k <= K; k ++ {
//       for n := 2; n <= N; n ++ {
//           min := dp[k][n]
//           for n0 := 1; n0 < n; n0 ++ {
//               max := dp[k-1][n-n0]+1
//               if dp[k][n0-1]+1 > max {
//                   max = dp[k][n0-1]+1
//               }
//               if max < min {
//                   min = max
//               }
//           }
//           dp[k][n] = min
//       }
//   }
//   return dp[K][N]
//}

//func superEggDrop(K int, N int) int {
//  dp := make([][]int, K+1)
//  for k := 0; k <= K; k ++ {
//      dp[k] = make([]int, N+1)
//      for n := 0; n <= N; n ++ {
//          dp[k][n] = n
//      }
//  }
//  for k := 2; k <= K; k ++ {
//      for n := 2; n <= N; n ++ {
//          lo := 1
//          hi := n-1
//          mi := (lo+hi)/2
//          for lo <= hi {
//              mi = (lo+hi)/2
//              if dp[k-1][mi-1] > dp[k][n-mi] {
//                  if hi == mi {
//                      hi--
//                  } else {
//                      hi = mi
//                  }
//              } else {
//                  lo = mi+1
//              }
//          }
//          max1 := dp[k-1][hi-1]+1
//          if dp[k][n-hi]+1 > max1 {
//             max1 = dp[k][n-hi] + 1
//          }
//          max2 := dp[k-1][lo-1]+1
//          if dp[k][n-lo]+1 > max2 {
//             max2 = dp[k][n-lo]+1
//          }
//          dp[k][n] = max1
//          if max2 < max1 {
//             dp[k][n] = max2
//          }
//      }
//  }
//  return dp[K][N]
//}

func superEggDrop(K int, N int) int {
	if K == 1 || N < 2 {
		return N
	}
	if K == 0 {
		return 0
	}
	dp := make([][]int, K+1)
	for i := 0; i <= K; i++ {
		dp[i] = make([]int, N+1)
	}

}
