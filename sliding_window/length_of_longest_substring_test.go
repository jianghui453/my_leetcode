package sliding_window

import (
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	var s string
	var h, r int

	s = "abcabcbb"
	r = lengthOfLongestSubstring(s)
	h = 3
	t.Logf("%t s=%s h=%d r=%d", r==h, s, h, r)

	s = "bbbbb"
	r = lengthOfLongestSubstring(s)
	h = 1
	t.Logf("%t s=%s h=%d r=%d", r==h, s, h, r)

	s = "pwwkew"
	r = lengthOfLongestSubstring(s)
	h = 3
	t.Logf("%t s=%s h=%d r=%d", r==h, s, h, r)

	s = ""
	r = lengthOfLongestSubstring(s)
	h = 0
	t.Logf("%t s=%s h=%d r=%d", r==h, s, h, r)

	s = "au"
	r = lengthOfLongestSubstring(s)
	h = 2
	t.Logf("%t s=%s h=%d r=%d", r==h, s, h, r)

	s = " "
	r = lengthOfLongestSubstring(s)
	h = 1
	t.Logf("%t s=%s h=%d r=%d", r==h, s, h, r)

	s = "  "
	r = lengthOfLongestSubstring(s)
	h = 1
	t.Logf("%t s=%s h=%d r=%d", r==h, s, h, r)

	s = "dvdf"
	r = lengthOfLongestSubstring(s)
	h = 3
	t.Logf("%t s=%s h=%d r=%d", r==h, s, h, r)

	//s = "你好你"
	//r = lengthOfLongestSubstring(s)
	//if 2 == r {
	//	msg := fmt.Sprintf("success: s=%s; r=%d", s, r)
	//	t.Log(msg)
	//} else {
	//	msg := fmt.Sprintf("fail: s=%s; r=%d", s, r)
	//	t.Log(msg)
	//}
}
