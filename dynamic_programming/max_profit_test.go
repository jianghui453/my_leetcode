package dynamic_programming

import "testing"
func TestMaxProfit188(t *testing.T) {
	type in struct{
		K int
		Prices []int
	}
	tests := []struct{
		In in
		Out int
	} {
		// {
		// 	In: in{
		// 		K: 2,
		// 		Prices: []int{2, 4, 1},
		// 	},
		// 	Out: 2,
		// },
		// {
		// 	In: in{
		// 		K: 2,
		// 		Prices: []int{3,2,6,5,0,3},
		// 	},
		// 	Out: 7,
		// },
		// {
		// 	In: in{
		// 		K: 2,
		// 		Prices: []int{6,1,3,2,4,7},
		// 	},
		// 	Out: 7,
		// },
		{
			In: in{
				K: 2,
				Prices: []int{1,2,4,2,5,7,2,4,9,0},
			},
			Out: 13,
		},
	}

	for _, test := range tests {
		out := maxProfit188(test.In.K, test.In.Prices)
		if out == test.Out {
			t.Log("test sucess:", test)
		} else {
			t.Error("test error:", test, out)
		}
	}
}