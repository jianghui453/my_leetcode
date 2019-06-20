package leet_code

import (
	"fmt"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	var s string
	var max int
	s = "abcabcbb"
	max = lengthOfLongestSubstring(s)
	if 3 == max {
		msg := fmt.Sprintf("success: s=%s; max=%d", s, max)
		t.Log(msg)
	} else {
		msg := fmt.Sprintf("fail: s=%s; max=%d", s, max)
		t.Log(msg)
	}
	s = "bbbbb"
	max = lengthOfLongestSubstring(s)
	if 1 == max {
		msg := fmt.Sprintf("success: s=%s; max=%d", s, max)
		t.Log(msg)
	} else {
		msg := fmt.Sprintf("fail: s=%s; max=%d", s, max)
		t.Log(msg)
	}
	s = "pwwkew"
	max = lengthOfLongestSubstring(s)
	if 3 == max {
		msg := fmt.Sprintf("success: s=%s; max=%d", s, max)
		t.Log(msg)
	} else {
		msg := fmt.Sprintf("fail: s=%s; max=%d", s, max)
		t.Log(msg)
	}
	s = ""
	max = lengthOfLongestSubstring(s)
	if 0 == max {
		msg := fmt.Sprintf("success: s=%s; max=%d", s, max)
		t.Log(msg)
	} else {
		msg := fmt.Sprintf("fail: s=%s; max=%d", s, max)
		t.Log(msg)
	}
	s = "au"
	max = lengthOfLongestSubstring(s)
	if 2 == max {
		msg := fmt.Sprintf("success: s=%s; max=%d", s, max)
		t.Log(msg)
	} else {
		msg := fmt.Sprintf("fail: s=%s; max=%d", s, max)
		t.Log(msg)
	}
	s = " "
	max = lengthOfLongestSubstring(s)
	if 1 == max {
		msg := fmt.Sprintf("success: s=%s; max=%d", s, max)
		t.Log(msg)
	} else {
		msg := fmt.Sprintf("fail: s=%s; max=%d", s, max)
		t.Log(msg)
	}
	s = "dvdf"
	max = lengthOfLongestSubstring(s)
	if 3 == max {
		msg := fmt.Sprintf("success: s=%s; max=%d", s, max)
		t.Log(msg)
	} else {
		msg := fmt.Sprintf("fail: s=%s; max=%d", s, max)
		t.Log(msg)
	}
	//s = "你好你"
	//max = lengthOfLongestSubstring(s)
	//if 2 == max {
	//	msg := fmt.Sprintf("success: s=%s; max=%d", s, max)
	//	t.Log(msg)
	//} else {
	//	msg := fmt.Sprintf("fail: s=%s; max=%d", s, max)
	//	t.Log(msg)
	//}
}
