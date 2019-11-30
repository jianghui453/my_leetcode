//给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
//
//如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。
//
//注意你不能在买入股票前卖出股票。
//
//示例 1:
//
//输入: [7,1,5,3,6,4]
//输出: 5
//解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
//示例 2:
//
//输入: [7,6,4,3,1]
//输出: 0
//解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。

package array

import (
	"math"
)

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	ret := 0
	min, max := math.MaxInt64, 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min, max = prices[i], prices[i]
		} else if prices[i] > max {
			max = prices[i]
			ret = int(math.Max(float64(ret), float64(max-min)))
		}
	}

	return ret
}

//func maxProfit(prices []int) int {
//	minv := math.MaxInt64
//	maxp := 0
//	pl := len(prices)
//	for i := 0; i < pl; i ++ {
//		if prices[i] < minv {
//			minv = prices[i]
//		} else if prices[i] - minv > maxp {
//			maxp = prices[i] - minv
//		}
//	}
//	return maxp
//}

//func maxProfit(prices []int) int {
//	pl := len(prices)
//	max := 0
//	for i := 0; i < pl-1; i ++ {
//		for j := i+1; j < pl; j ++ {
//			if prices[j] - prices[i] > max {
//				max = prices[j] - prices[i]
//			}
//		}
//	}
//	return max
//}
