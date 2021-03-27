package sort

import "testing"

func TestMaximumGap(t *testing.T) {
	tests := []struct{
		In []int
		Out int
	} {
		// {
		// 	In: []int{3,6,9,1},
		// 	Out: 3,
		// },
		// {
		// 	In: []int{10},
		// 	Out: 0,
		// },
		// {
		// 	In: []int{1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 3},
		// 	Out: 1,
		// },
		// {
		// 	In: []int{1, 1, 1, 1, 1},
		// 	Out: 0,
		// },
		{
			In: []int{1,3,100},
			Out: 97,
		},
	}
	for _, test := range tests {
		if out := maximumGap(test.In); out == test.Out {
			t.Log("ok:", test)
		} else {
			t.Error("error:", test, out)
		}
	}
}