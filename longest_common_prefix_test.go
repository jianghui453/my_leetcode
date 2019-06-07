package leet_code

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	var strs []string
	var hope, ret string

	strs = []string{"flower","flow","flight"}
	hope = "fl"
	ret = longestCommonPrefix(strs)
	if hope == ret {
		t.Logf("success: strs = %v; hope = %s; ret = %s\n", strs, hope, ret)
	} else {
		t.Errorf("fail: strs = %v; hope = %s; ret = %s\n", strs, hope, ret)
	}

	strs = []string{"dog","racecar","car"}
	hope = ""
	ret = longestCommonPrefix(strs)
	if hope == ret {
		t.Logf("success: strs = %v; hope = %s; ret = %s\n", strs, hope, ret)
	} else {
		t.Errorf("fail: strs = %v; hope = %s; ret = %s\n", strs, hope, ret)
	}

	strs = []string{}
	hope = ""
	ret = longestCommonPrefix(strs)
	if hope == ret {
		t.Logf("success: strs = %v; hope = %s; ret = %s\n", strs, hope, ret)
	} else {
		t.Errorf("fail: strs = %v; hope = %s; ret = %s\n", strs, hope, ret)
	}
}