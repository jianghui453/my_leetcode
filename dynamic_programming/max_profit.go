package dynamic_programming

import (
	// "fmt"
	// "encoding/json"
)

func maxProfit188(k int, prices []int) int {
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	m := map[int]map[int]int{}
	for k0 := 1; k0 <= k; k0++ {
		m[k0] = map[int]int{}
		for i := 0; i < len(prices); i++ {
			m[k0][i] = m[k0][i-1]
			for j := i-1; j >= 0; j-- {
				if prices[j] >= prices[i] {
					break
				}
				if j == 0 {
					m[k0][i] = max(m[k0][i], prices[i] - prices[j])
				} else {
					m[k0][i] = max(m[k0][i], m[k0-1][j-1] + prices[i] - prices[j])
				}				
			}
		}
	}
	// mJson, _ := json.MarshalIndent(m, "", "\t")
	// fmt.Println(string(mJson))
	return m[k][len(prices)-1]
}
