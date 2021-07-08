package string

import (
	"testing"
)

func TestStrongPasswordChecker(t *testing.T) {
	var (
		s string
		hope, ret int
	)

	s = "aaa"
	hope = 3
	ret = strongPasswordChecker(s)
	t.Log(ret == hope, s, hope, ret)

	s = "aaa111"
	hope = 2
	ret = strongPasswordChecker(s)
	t.Log(ret == hope, s, hope, ret)

	s = "aaaaaaaaaaaaaaaaaaaaa"
	hope = 7
	ret = strongPasswordChecker(s)
	t.Log(ret == hope, s, hope, ret)
}