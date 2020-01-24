package string

import (
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	var (
		s, strt string
		hope, ret bool
	)

	s = "ace"
	strt = "abcde"
	hope = true
	ret = isSubsequence(s, strt)
	t.Log(ret == hope, s, strt, hope, ret)

	s = "aec"
	strt = "abcde"
	hope = false
	ret = isSubsequence(s, strt)
	t.Log(ret == hope, s, strt, hope, ret)
}