package dynamic_programming

import "testing"

func TestIsScramble(t *testing.T) {
	var s1, s2 string
	var hope, ret bool

	s1 = "great"
	s2 = "rgeat"
	hope = true
	ret = isScramble(s1, s2)
	t.Logf("%t s1=%s s2=%s hope=%t ret=%t", hope == ret, s1, s2, hope, ret)

	s1 = "abcde"
	s2 = "caebd"
	hope = false
	ret = isScramble(s1, s2)
	t.Logf("%t s1=%s s2=%s hope=%t ret=%t", hope == ret, s1, s2, hope, ret)

	s1 = "ccabcbabcbabbbbcbb"
	s2 = "bbbbabccccbbbabcba"
	hope = false
	ret = isScramble(s1, s2)
	t.Logf("%t s1=%s s2=%s hope=%t ret=%t", hope == ret, s1, s2, hope, ret)
}
