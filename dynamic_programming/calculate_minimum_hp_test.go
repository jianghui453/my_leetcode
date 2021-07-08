package dynamic_programming

import "testing"

func TestCalculateMinimumHP(t *testing.T) {
	tests := []struct{
		In [][]int
		Out int
	} {
		{
			[][]int{
				{-2, -3, 3},
				{-5, -10, 1},
				{10, 30, -5},
			},
			7,
		},
		{
			[][]int{
				{-3, 5},
			},
			4,
		},
	}
	for _, test := range tests {
		out := calculateMinimumHP0(test.In)
		if out == test.Out {
			t.Log("ok:", test)
		} else {
			t.Error("error:", test, out)
		}
	}
}
