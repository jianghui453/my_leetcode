package dynamic_programming

import (
	"encoding/json"
	"fmt"
)

func maxProfit188(k int, prices []int) int {
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	// map[k][i][j]
	m := map[int]map[int]map[int]int{}
	for k0 := 1; k0 <= k; k0++ {
		m[k0] = map[int]map[int]int{}
		for i := 0; i < len(prices); i++ {
			m[k0][i] = map[int]int{}
			for j := i+1; j < len(prices); j++ {
				// fmt.Println(k0, i, j)
				m[k0][i][j] = m[k0][i][j-1]
				for j0 := j-1; j0 >= i; j0-- {
					if prices[j0] > prices[j0+1] {
						m[k0][i][j] = max(m[k0][i][j], m[k0-1][i][j0] + (prices[j] - prices[j0+1]))
						break
					}
					if j0 == i {
						m[k0][i][j] = max(m[k0][i][j], m[k0-1][i][j0] + (prices[j] - prices[j0]))
						break
					}
				}
			}
		}
	}
	mJson, _ := json.MarshalIndent(m, "", "\t")
	fmt.Println(string(mJson))
	return m[k][0][len(prices)-1]
}