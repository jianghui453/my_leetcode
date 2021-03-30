package math_algorithm

import "testing"

func TestTrailingZeroes(t *testing.T) {
	tests := []struct{
		In int
		Out int
	} {
		{
			3,
			0,
		},
		{
			5,
			1,
		},
		{
			7,
			1,
		},
		{
			30,
			7,
		},
	}
	for _, test := range tests {
		out := trailingZeroes(test.In)
		if out == test.Out {
			t.Log("ok:", test)
		} else {
			t.Error("error:", test, out)
		}
	}
}