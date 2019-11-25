package dynamic_programming

import (
	"testing"
)

func TestIsMatch(t *testing.T) {
	var s, p string
	var r, h bool

	s = "aa"
	p = "a"
	h = false
	r = isMatch(s, p)
	t.Logf("%t s=%s p=%s h=%t r=%t", r == h, s, p, h, r)

	s = "aa"
	p = "*"
	h = true
	r = isMatch(s, p)
	t.Logf("%t s=%s p=%s h=%t r=%t", r == h, s, p, h, r)

	s = "cb"
	p = "?a"
	h = false
	r = isMatch(s, p)
	t.Logf("%t s=%s p=%s h=%t r=%t", r == h, s, p, h, r)

	s = "acdcb"
	p = "*a*b"
	h = true
	r = isMatch(s, p)
	t.Logf("%t s=%s p=%s h=%t r=%t", r == h, s, p, h, r)

	s = "acdcb"
	p = "a*c?b"
	h = false
	r = isMatch(s, p)
	t.Logf("%t s=%s p=%s h=%t r=%t", r == h, s, p, h, r)

	s = ""
	p = ""
	h = true
	r = isMatch(s, p)
	t.Logf("%t s=%s p=%s h=%t r=%t", r == h, s, p, h, r)

	s = ""
	p = "*"
	h = true
	r = isMatch(s, p)
	t.Logf("%t s=%s p=%s h=%t r=%t", r == h, s, p, h, r)

	s = ""
	p = "?"
	h = false
	r = isMatch(s, p)
	t.Logf("%t s=%s p=%s h=%t r=%t", r == h, s, p, h, r)

	s = "q"
	p = ""
	h = false
	r = isMatch(s, p)
	t.Logf("%t s=%s p=%s h=%t r=%t", r == h, s, p, h, r)

	s = "aaabbbaabaaaaababaabaaabbabbbbbbbbaabababbabbbaaaaba"
	p = "a*******b"
	h = false
	r = isMatch(s, p)
	t.Logf("%t s=%s p=%s h=%t r=%t", r == h, s, p, h, r)

	s = "aab"
	p = "c*a*b"
	h = false
	r = isMatch(s, p)
	t.Logf("%t s=%s p=%s h=%t r=%t", r == h, s, p, h, r)

	s = "babbbbaabababaabbababaababaabbaabababbaaababbababaaaaaabbabaaaabababbabbababbbaaaababbbabbbbbbbbbbaabbb"
	p = "b**bb**a**bba*b**a*bbb**aba***babbb*aa****aabb*bbb***a"
	h = false
	r = isMatch(s, p)
	t.Logf("%t s=%s p=%s h=%t r=%t", r == h, s, p, h, r)

	s = "abefcdgiescdfimde"
	p = "ab*cd?i*de"
	h = true
	r = isMatch(s, p)
	t.Logf("%t s=%s p=%s h=%t r=%t", r == h, s, p, h, r)

	s = "c"
	p = "*?*"
	h = true
	r = isMatch(s, p)
	t.Logf("%t s=%s p=%s h=%t r=%t", r == h, s, p, h, r)
}
