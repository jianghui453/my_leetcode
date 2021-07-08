package hash_table

import "testing"

func TestFindRepeatedDnaSequences(t *testing.T) {
	tests := []struct{
		In string
		Out []string
	}{
		{
			In: "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT",
			Out: []string{"AAAAACCCCC","CCCCCAAAAA"},
		},
		{
			In: "AAAAAAAAAAAAA",
			Out: []string{"AAAAAAAAAA"},
		},
		{
			In: "AAAAAAAAAAA",
			Out: []string{"AAAAAAAAAA"},
		},
	}

	for _, test := range tests {
		out := findRepeatedDnaSequences(test.In)
		outMap := make(map[string]bool)
		for _, item := range out {
			outMap[item] = true
		}
		ok := true
		for _, item := range test.Out {
			if !outMap[item] {
				t.Error("error test:", test)
				ok = false
				break
			}
		}
		if ok {
			t.Log("sucess test:", test)
		}
		
	}
}