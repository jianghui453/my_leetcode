package string

import (
	"testing"
)

func TestCanConvert(t *testing.T) {
	var (
		str1, str2 string
		ret, hope bool
	)

	str1 = "aabcc"
	str2 = "ccdee"
	hope = true
	ret = canConvert(str1, str2)
	t.Log(ret == hope, "str1 =", str1, "str2 =", str2, "hope =", hope, "ret =", ret)

	str1 = "leetcode"
	str2 = "codeleet"
	hope = false
	ret = canConvert(str1, str2)
	t.Log(ret == hope, "str1 =", str1, "str2 =", str2, "hope =", hope, "ret =", ret)

	str1 = "abcdefghijklmnopqrstuvwxyz"
	str2 = "bcdefghijklmnopqrstuvwxyzq"
	hope = true
	ret = canConvert(str1, str2)
	t.Log(ret == hope, "str1 =", str1, "str2 =", str2, "hope =", hope, "ret =", ret)
}