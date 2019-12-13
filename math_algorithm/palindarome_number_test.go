package leet_code

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	var x int
	var ret, hope bool

	x = 121
	hope = true
	ret = isPalindrome(x)
	t.Log(ret == hope, x, hope, ret)

	x = -121
	hope = false
	ret = isPalindrome(x)
	t.Log(ret == hope, x, hope, ret)

	x = 10
	hope = false
	ret = isPalindrome(x)
	t.Log(ret == hope, x, hope, ret)

	x = 0
	hope = true
	ret = isPalindrome(x)
	t.Log(ret == hope, x, hope, ret)

	x = 1
	hope = true
	ret = isPalindrome(x)
	t.Log(ret == hope, x, hope, ret)
}
