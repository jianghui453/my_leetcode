package leet_code

import (
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	var n int
	var hope, ret []string

	n = 3
	hope = []string{
		"((()))",
		"(()())",
		"(())()",
		"()(())",
		"()()()",
	}
	ret = generateParenthesis(n)
	t.Logf("n = %d; hope = %v; ret = %v", n, hope, ret)
}