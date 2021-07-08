package bit

import "testing"

func TestReverseBits(t *testing.T) {
	tests := []struct{
		In uint32
		Out uint32
	} {
		{
			In: 43261596,
			Out: 964176192,
		},
		{
			In: 4294967293,
			Out: 3221225471,
		},
	}
	for _, test := range tests {
		out := reverseBits(test.In)
		if out == test.Out {
			t.Log("test sucess:", test)
		} else {
			t.Error("test error:", test, out)
		}
	}
}