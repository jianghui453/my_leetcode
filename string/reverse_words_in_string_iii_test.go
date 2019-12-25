package string

import (
	"testing"
)

func Test_reverseWords(t *testing.T) {
	var (
		s, hope, ret string
	)

	s = "Let's take LeetCode contest"
	hope = "s'teL ekat edoCteeL tsetnoc"
	ret = reverseWords(s)
	t.Log(ret == hope, s, hope, ret)

	s = ""
	hope = ""
	ret = reverseWords(s)
	t.Log(ret == hope, s, hope, ret)
	
	s = "L"
	hope = "L"
	ret = reverseWords(s)
	t.Log(ret == hope, s, hope, ret)
}