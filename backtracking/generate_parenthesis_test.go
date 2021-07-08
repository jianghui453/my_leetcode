package backtracking

import (
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	var n int
	var h, r []string

	n = 3
	h = []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	r = generateParenthesis(n)
	t.Logf("n=%d \nh=%v \nr=%v", n, h, r)
}
