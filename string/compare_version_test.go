package string

import "testing"

func TestCompareVersion(t *testing.T) {
	tests := []struct{
		In [2]string
		Out int
	} {
		{
			[2]string{"1.01", "1.001"},
			0,
		},
		{
			[2]string{"1.0", "1.0.0"},
			0,
		},
		{
			[2]string{"0.1", "1.1"},
			-1,
		},
		{
			[2]string{"1.0.1", "1"},
			1,
		},
		{
			[2]string{"7.5.2.4", "7.5.3"},
			-1,
		},
	}

	for _, test := range tests {
		if out := compareVersion(test.In[0], test.In[1]); out == test.Out {
			t.Log("ok:", test)
		} else {
			t.Error("error:", test, out)
		}
	}
}