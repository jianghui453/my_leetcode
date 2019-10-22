package longest_parenthese

import (
	"testing"
)

func TestLongestValidParenthese(t *testing.T) {
	var s string
	var hope, ret int

	s = ")()())"
	hope = 4
	ret = longestValidParentheses(s)
	t.Logf("s = %s; hope = %d; ret = %d", s, hope, ret)

	s = "(()"
	hope = 2
	ret = longestValidParentheses(s)
	t.Logf("s = %s; hope = %d; ret = %d", s, hope, ret)

	s = "()(())"
	hope = 6
	ret = longestValidParentheses(s)
	t.Logf("s = %s; hope = %d; ret = %d", s, hope, ret)

	s = "()(()"
	hope = 2
	ret = longestValidParentheses(s)
	t.Logf("s = %s; hope = %d; ret = %d", s, hope, ret)

	s = "(()(((()"
	hope = 2
	ret = longestValidParentheses(s)
	t.Logf("s = %s; hope = %d; ret = %d", s, hope, ret)
}
