package math_algorithm

import "testing"

func TestTitleToNumber(t *testing.T) {
	tests := []struct{
		In string
		Out int
	} {
		{
			"A",
			1,
		},
		{
			"AB",
			28,
		},
		{
			"ZY",
			701,
		},
		{
			"FXSHRXW",
			2147483647,
		},
	}
	for _, test := range tests {
		out := titleToNumber(test.In)
		if out == test.Out {
			t.Log("ok:", test)
		} else {
			t.Error("error:", test, out)
		}
	}
}
