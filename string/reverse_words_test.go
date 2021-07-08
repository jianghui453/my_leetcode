package string

import "testing"

func TestReverseWords151(t *testing.T) {
	tests := []struct{
		In string
		Out string
	} {
		{
			In: "the sky is blue",
			Out: "blue is sky the",
		},
		{
			In: "  hello world!  ",
			Out: "world! hello",
		},
		{
			In: "a good   example",
			Out: "example good a",
		},
		{
			In: "  Bob    Loves  Alice   ",
			Out: "Alice Loves Bob",
		},
		{
			In: "Alice does not even like bob",
			Out: "bob like even not does Alice",
		},
	}
	for _, test := range tests {
		out := reverseWords151(test.In)
		if out == test.Out {
			t.Log("sucess:", test)
		} else {
			t.Error("error:", test, out)
		}
	}
}
