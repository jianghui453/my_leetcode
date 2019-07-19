package is_match

import (
	"testing"
)

func TestIsMatch2(t *testing.T) {
	var s, p string
	var ret, hope bool

	// s = "aa"
	// p = "a"
	// hope = false
	// ret = isMatch2(s, p)
	// t.Logf("s=%s; p = %s; hope = %t; ret = %t.", s, p, hope, ret)

	// s = "aa"
	// p = "*"
	// hope = true
	// ret = isMatch2(s, p)
	// t.Logf("s=%s; p = %s; hope = %t; ret = %t.", s, p, hope, ret)

	// s = "cb"
	// p = "?a"
	// hope = false
	// ret = isMatch2(s, p)
	// t.Logf("s=%s; p = %s; hope = %t; ret = %t.", s, p, hope, ret)

	// s = "acdcb"
	// p = "*a*b"
	// hope = true
	// ret = isMatch2(s, p)
	// t.Logf("s=%s; p = %s; hope = %t; ret = %t.", s, p, hope, ret)

	// s = "acdcb"
	// p = "a*c?b"
	// hope = false
	// ret = isMatch2(s, p)
	// t.Logf("s=%s; p = %s; hope = %t; ret = %t.", s, p, hope, ret)

	// s = ""
	// p = ""
	// hope = true
	// ret = isMatch2(s, p)
	// t.Logf("s=%s; p = %s; hope = %t; ret = %t.", s, p, hope, ret)

	// s = ""
	// p = "*"
	// hope = true
	// ret = isMatch2(s, p)
	// t.Logf("s=%s; p = %s; hope = %t; ret = %t.", s, p, hope, ret)

	// s = ""
	// p = "?"
	// hope = false
	// ret = isMatch2(s, p)
	// t.Logf("s=%s; p = %s; hope = %t; ret = %t.", s, p, hope, ret)

	// s = "q"
	// p = ""
	// hope = false
	// ret = isMatch2(s, p)
	// t.Logf("s=%s; p = %s; hope = %t; ret = %t.", s, p, hope, ret)

	// s = "aaabbbaabaaaaababaabaaabbabbbbbbbbaabababbabbbaaaaba"
	// p = "a*******b"
	// hope = false
	// ret = isMatch2(s, p)
	// t.Logf("s=%s; p = %s; hope = %t; ret = %t.", s, p, hope, ret)

	// s = "aab"
	// p = "c*a*b"
	// hope = false
	// ret = isMatch2(s, p)
	// t.Logf("s=%s; p = %s; hope = %t; ret = %t.", s, p, hope, ret)

	// s = "babbbbaabababaabbababaababaabbaabababbaaababbababaaaaaabbabaaaabababbabbababbbaaaababbbabbbbbbbbbbaabbb"
	// p = "b**bb**a**bba*b**a*bbb**aba***babbb*aa****aabb*bbb***a"
	// hope = false
	// ret = isMatch2(s, p)
	// t.Logf("s=%s; p = %s; hope = %t; ret = %t.", s, p, hope, ret)

	// s = "abefcdgiescdfimde"
	// p = "ab*cd?i*de"
	// hope = true
	// ret = isMatch2(s, p)
	// t.Logf("s=%s; p = %s; hope = %t; ret = %t.", s, p, hope, ret)

	s = "c"
	p = "*?*"
	hope = true
	ret = isMatch2(s, p)
	t.Logf("s=%s; p = %s; hope = %t; ret = %t.", s, p, hope, ret)
}
