package find_sub_string

import (
	"testing"
)

func TestFindSubString(t *testing.T) {
	var s string
	var words []string
	var h, r []int

	s = "barfoothefoobarman"
	words = []string{"foo", "bar"}
	h = []int{0, 9}
	r = findSubstring(s, words)
	t.Logf("s=%s words=%v h=%v r=%v", s, words, h, r)

	s = "barfoothefoobarman"
	words = []string{"foo"}
	h = []int{3, 9}
	r = findSubstring(s, words)
	t.Logf("s=%s words=%v h=%v r=%v", s, words, h, r)

	s = "barfoofoobarthefoobarman"
	words = []string{"bar","foo","the"}
	h = []int{6, 9, 12}
	r = findSubstring(s, words)
	t.Logf("s=%s words=%v h=%v r=%v", s, words, h, r)

	s = "wordgoodgoodgoodbestword"
	words = []string{"word","good","best","word"}
	h = []int{}
	r = findSubstring(s, words)
	t.Logf("s=%s words=%v h=%v r=%v", s, words, h, r)
}
