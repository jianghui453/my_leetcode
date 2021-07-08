package math_algorithm

import "testing"

func TestConvertToTitle(t *testing.T) {
	tests := []struct{
		In int
		Out string
	} {
		{
			1,
			"A",
		},
		{
			28,
			"AB",
		},
		{
			701,
			"ZY",
		},
		{
			2147483647,
			"FXSHRXW",
		},
	}
	for _, test := range tests {
		out := convertToTitle(test.In)
		if out == test.Out {
			t.Log("ok:", test)
		} else {
			t.Error("error:", test, out)
		}
	}
}