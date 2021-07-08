package array

import "testing"

func TestFindMin153(t *testing.T) {
	tests := []struct{
		In []int
		Out int
	} {
		{
			In: []int{3,4,5,1,2},
			Out: 1,
		},
		{
			In: []int{4,5,6,7,0,1,2},
			Out: 0,
		},
		{
			In: []int{1},
			Out: 1,
		},
		{
			In: []int{11,13,15,17},
			Out: 11,
		},
	}
	for _, test := range tests {
		out := findMin153(test.In)
		if out == test.Out {
			t.Log("sucess:", test)
		} else {
			t.Error("error:", test, out)
		}
	}
}

func TestFindMin154(t *testing.T) {
	tests := []struct{
		In []int
		Out int
	} {
		{
			In: []int{1,3,5},
			Out: 1,
		},
		{
			In: []int{2,2,2,0,1},
			Out: 0,
		},
		{
			In: []int{2,2,2,2,2},
			Out: 2,
		},
	}
	for _, test := range tests {
		out := findMin154(test.In)
		if out == test.Out {
			t.Log("sucess:", test)
		} else {
			t.Error("error:", test, out)
		}
	}
}