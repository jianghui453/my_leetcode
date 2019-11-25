package dynamic_programming

import (
	"testing"
)

func TestLongestValidParenthese(t *testing.T) {
	var s string
	var h, r int

	s = ")()())"
	h = 4
	r = longestValidParentheses(s)
	t.Logf("%t s=%s h=%d r=%d", r == h, s, h, r)

	s = "(()"
	h = 2
	r = longestValidParentheses(s)
	t.Logf("%t s=%s h=%d r=%d", r == h, s, h, r)

	s = "()(())"
	h = 6
	r = longestValidParentheses(s)
	t.Logf("%t s=%s h=%d r=%d", r == h, s, h, r)

	s = "()(()"
	h = 2
	r = longestValidParentheses(s)
	t.Logf("%t s=%s h=%d r=%d", r == h, s, h, r)

	s = "(()(((()"
	h = 2
	r = longestValidParentheses(s)
	t.Logf("%t s=%s h=%d r=%d", r == h, s, h, r)

	s = "((()))())"
	h = 8
	r = longestValidParentheses(s)
	t.Logf("%t s=%s h=%d r=%d", r == h, s, h, r)
}
