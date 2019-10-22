package string

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	var strs []string
	var h, r string

	strs = []string{"flower", "flow", "flight"}
	h = "fl"
	r = longestCommonPrefix(strs)
	t.Logf("%t strs=%v h=%s r=%s", r == h, strs, h, r)

	strs = []string{"dog", "racecar", "car"}
	h = ""
	r = longestCommonPrefix(strs)
	t.Logf("%t strs=%v h=%s r=%s", r == h, strs, h, r)

	strs = []string{}
	h = ""
	r = longestCommonPrefix(strs)
	t.Logf("%t strs=%v h=%s r=%s", r == h, strs, h, r)
}
