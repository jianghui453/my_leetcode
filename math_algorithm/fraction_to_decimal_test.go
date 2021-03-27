package math_algorithm

import "testing"

func TestFractionToDecimal(t *testing.T) {
	type in struct{
		Numerator, Denominator int
	}
	tests := []struct{
		In in
		Out string
	} {
		{
			in{1, 2},
			"0.5",
		},
		{
			in{2, 1},
			"2",
		},
		{
			in{2, 3},
			"0.(6)",
		},
		{
			in{4, 333},
			"0.(012)",
		},
		{
			in{1, 5},
			"0.2",
		},
		{
			in{1, 6},
			"0.1(6)",
		},
		{
			in{22, 7},
			"3.(142857)",
		},
		{
			in{-50, 8},
			"-6.25",
		},
		{
			in{-22, -2},
			"11",
		},
	}
	for _, test := range tests {
		if out := fractionToDecimal(test.In.Numerator, test.In.Denominator); out == test.Out {
			t.Log("ok:", test)
		} else {
			t.Error("error:", test, out)
		}
	}
}