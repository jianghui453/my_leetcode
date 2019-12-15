package string

import (
	"testing"
)

func Test_decodeString(t *testing.T) {
	var s, hope, ret string

	s = "3[a]2[bc]"
	hope = "aaabcbc"
	ret = decodeString(s)
	t.Log(ret == hope, s, hope, ret)

	s = "3[a2[c]]"
	hope = "accaccacc"
	ret = decodeString(s)
	t.Log(ret == hope, s, hope, ret)

	s = "2[abc]3[cd]ef"
	hope = "abcabccdcdcdef"
	ret = decodeString(s)
	t.Log(ret == hope, s, hope, ret)

	s = "2[abc2[a]b2[c]]3[cd]ef"
	hope = "abcaabccabcaabcccdcdcdef"
	ret = decodeString(s)
	t.Log(ret == hope, s, hope, ret)

	s = "a"
	hope = "a"
	ret = decodeString(s)
	t.Log(ret == hope, s, hope, ret)

	s = ""
	hope = ""
	ret = decodeString(s)
	t.Log(ret == hope, s, hope, ret)
}