package bit

import "testing"

func TestHammingWeight(t *testing.T) {
	tests := []struct{
		In uint32
		Out int
	} {
		{
			In: 11,
			Out: 3,
		},
		{
			In: 128,
			Out: 1,
		},
		{
			In: 4294967293,
			Out: 31,
		},
	}
	for _, test := range tests {
		out := hammingWeight(test.In)
		if out == test.Out {
			t.Log("test sucess:", test)
		} else {
			t.Error("test error:", test, out)
		}
	}
}
