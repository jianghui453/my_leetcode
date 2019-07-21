package group_anagrams

import "testing"

func TestGroupAnagrams(t *testing.T) {
	var strs []string
	var ret [][]string

	strs = []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	ret = groupAnagrams(strs)
	t.Logf("strs=%v ret=%v", strs, ret)
}