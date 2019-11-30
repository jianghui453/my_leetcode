//给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
//
//设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
//
//注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
//示例 1:
//
//输入: [3,3,5,0,0,3,1,4]
//输出: 6
//解释: 在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。
//     随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3 。
//示例 2:
//
//输入: [1,2,3,4,5]
//输出: 4
//解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
//     注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
//     因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
//示例 3:
//
//输入: [7,6,4,3,1]
//输出: 0
//解释: 在这个情况下, 没有交易完成, 所以最大利润为 0。

package array

func maxProfit(prices []int) int {}

//func maxProfit(prices []int) int {
//	pl := len(prices)
//	if pl == 0 {
//		return 0
//	}
//	max := 0
//	pricesNew := make([]int, 0)
//	minV := prices[0]
//	maxV := prices[0]
//	plNew := 0
//
//	// optimize
//	for i := 0; i < pl; i ++ {
//		if prices[i] >= maxV {
//			maxV = prices[i]
//		} else {
//			if minV < maxV {
//				pricesNew = append(pricesNew, minV)
//				pricesNew = append(pricesNew, maxV)
//				plNew += 2
//			}
//			minV = prices[i]
//			maxV = prices[i]
//		}
//	}
//	if minV < maxV {
//		pricesNew = append(pricesNew, minV)
//		pricesNew = append(pricesNew, maxV)
//		plNew += 2
//	}
//	for i := 1; i < plNew; i ++ {
//		if pricesNew[i] > pricesNew[i-1] {
//			var max1, max2 int
//			if i < plNew-1 {
//				max1 = getMax(pricesNew[: i+1], i+1)
//				max2 = getMax(pricesNew[i+1: ], plNew-i-1)
//			} else {
//				max1 = getMax(pricesNew, plNew)
//			}
//
//			if max1+max2 > max {
//				max = max1+max2
//			}
//		}
//	}
//	return max
//}
//
//func getMax(prices []int, pl int) int {
//	if pl == 0 {
//		return 0
//	}
//	max := 0
//	min := prices[0]
//	for i := 1; i < pl; i ++ {
//		if prices[i] < min {
//			min = prices[i]
//		} else if prices[i] - min > max {
//			max = prices[i] - min
//		}
//	}
//	return max
//}

// func maxProfit(prices []int) int {
// 	pl := len(prices)
// 	if pl < 2 {
// 		return 0
// 	}
// 	d := 2
// 	// 三个变量：交易天数（i）、已交易次数(j)、当前交易状态(k，存在两种情况：持有、未持有)
// 	// 买入不算交易次数
// 	dp := make([][][2]int, pl+1)
// 	for i := 0; i <= pl; i++ {
// 		dp[i] = make([][2]int, d+1)
// 		if i > 0 {
// 			for j := 1; j <= d; j++ {
// 				if i == 1 {
// 					dp[i][j][0] = 0
// 					dp[i][j][1] = -prices[0]
// 				} else {
// 					dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i-1])
// 					dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i-1])
// 				}
// 			}
// 		}
// 	}
// 	return dp[pl][d][0]
// }

//func maxProfit(prices []int) int {
//	k := 2
//	if len(prices) <= 1 {
//		return 0
//	}
//	if k > len(prices)/2 {
//		return maxProfitpt(prices)
//	}
//	var dp = make([][][2]int, len(prices))
//	for i := 0; i < len(prices); i++ {
//		dp[i] = make([][2]int, k+1)
//	}
//	for i := 0; i < len(prices); i++ {
//		for j := k; j >= 1; j-- {
//			if i == 0 {
//				dp[i][j][0] = 0
//				dp[i][j][1] = -prices[0]
//			} else {
//				dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
//				dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
//			}
//		}
//	}
//	return dp[len(prices)-1][k][0]
//}
//
//func maxProfitpt(prices []int) int {
//	if len(prices) == 0 {
//		return 0
//	}
//	var dp0, dp1 = 0, -prices[0]
//	for i := 1; i < len(prices); i++ {
//		dp0 = max(dp0, dp1+prices[i])
//		dp1 = max(dp1, dp0-prices[i])
//	}
//	return dp0
//}
//
// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }
