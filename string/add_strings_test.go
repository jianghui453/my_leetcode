package string

import (
	"testing"
)

func Test_addStrings(t *testing.T) {
	var num1, num2, hope, ret string

	num1, num2 = "1", ""
	hope = "1"
	ret = addStrings(num1, num2)
	t.Log(ret == hope, num1, num2, hope, ret)

	num1, num2 = "9", "9"
	hope = "18"
	ret = addStrings(num1, num2)
	t.Log(ret == hope, num1, num2, hope, ret)

	num1, num2 = "9", "1"
	hope = "10"
	ret = addStrings(num1, num2)
	t.Log(ret == hope, num1, num2, hope, ret)

	num1, num2 = "99", "1"
	hope = "100"
	ret = addStrings(num1, num2)
	t.Log(ret == hope, num1, num2, hope, ret)

	num1, num2 = "119", "1"
	hope = "120"
	ret = addStrings(num1, num2)
	t.Log(ret == hope, num1, num2, hope, ret)

	num1, num2 = "8", "1"
	hope = "9"
	ret = addStrings(num1, num2)
	t.Log(ret == hope, num1, num2, hope, ret)
}