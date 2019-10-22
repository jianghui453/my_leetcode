package find_sub_string

import (
	"testing"
)

func TestFindSubString(t *testing.T) {
	var s string
	var words []string
	var hope []int
	var ret []int

	s = "barfoothefoobarman"
	words = []string{"foo", "bar"}
	hope = []int{0, 9}
	ret = findSubstring(s, words)
	t.Logf("s = %s; words = %v; hope = %v; ret = %v", s, words, hope, ret)
}
