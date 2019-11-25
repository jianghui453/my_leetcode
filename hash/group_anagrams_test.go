package hash

import "testing"

func TestGroupAnagrams(t *testing.T) {
	var strs []string
	var r, h [][]string

	strs = []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	h = [][]string{
		{"ate", "eat", "tea"},
		{"nat", "tan"},
		{"bat"},
	}
	r = groupAnagrams(strs)
	t.Logf("strs=%v \nr=%v \nh=%v", strs, r, h)
}
