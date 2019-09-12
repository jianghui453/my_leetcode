package dynamic_programming

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	var s, hope, ret string

	s = "babad"
	hope = "bab"
	ret = longestPalindrome(s)
	t.Logf("s=%s hope=%s ret=%s", s, hope, ret)

	s = "cbbd"
	hope = "bb"
	ret = longestPalindrome(s)
	t.Logf("s=%s hope=%s ret=%s", s, hope, ret)

	s = "a"
	hope = "a"
	ret = longestPalindrome(s)
	t.Logf("s=%s hope=%s ret=%s", s, hope, ret)

	s = "ccc"
	hope = "ccc"
	ret = longestPalindrome(s)
	t.Logf("s=%s hope=%s ret=%s", s, hope, ret)

	s = "aaaa"
	hope = "aaaa"
	ret = longestPalindrome(s)
	t.Logf("s=%s hope=%s ret=%s", s, hope, ret)

	s = "bananas"
	hope = "anana"
	ret = longestPalindrome(s)
	t.Logf("s=%s hope=%s ret=%s", s, hope, ret)

	s = "ababababababa"
	hope = "ababababababa"
	ret = longestPalindrome(s)
	t.Logf("s=%s hope=%s ret=%s", s, hope, ret)
}
