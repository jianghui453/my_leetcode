package leet_code

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	var s, ret string
	s = "babad"
	ret = "bab"
	if ret == longestPalindrome(s) {
		t.Log(fmt.Sprintf("success: s=%s; ret=%s", s, ret))
	} else {
		t.Error(fmt.Sprintf("success: s=%s; ret=%s", s, ret))
	}

	s = "cbbd"
	ret = "bb"
	if ret == longestPalindrome(s) {
		t.Log(fmt.Sprintf("success: str=%s; ret=%s", s, ret))
	} else {
		t.Error(fmt.Sprintf("success: str=%s; ret=%s", s, ret))
	}

	s = "a"
	ret = "a"
	if ret == longestPalindrome(s) {
		t.Log(fmt.Sprintf("success: str=%s; ret=%s", s, ret))
	} else {
		t.Error(fmt.Sprintf("success: str=%s; ret=%s", s, ret))
	}

	//s = "cbbd"
	//ret = "bb"
	//if ret == longestPalindrome(s) {
	//	t.Log(fmt.Sprintf("success: str=%s; ret=%s", s, ret))
	//} else {
	//	t.Error(fmt.Sprintf("success: str=%s; ret=%s", s, ret))
	//}
}
