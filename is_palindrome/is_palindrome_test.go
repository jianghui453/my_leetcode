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
	if ret == hope {
		t.Logf("success: x=%d; hope = %t; ret = %t", x, hope, ret)
	} else {
		t.Errorf("fail: x=%d; hope = %t; ret = %t", x, hope, ret)
	}

	x = -121
	hope = false
	ret = isPalindrome(x)
	if ret == hope {
		t.Logf("success: x=%d; hope = %t; ret = %t", x, hope, ret)
	} else {
		t.Errorf("fail: x=%d; hope = %t; ret = %t", x, hope, ret)
	}

	x = 10
	hope = false
	ret = isPalindrome(x)
	if ret == hope {
		t.Logf("success: x=%d; hope = %t; ret = %t", x, hope, ret)
	} else {
		t.Errorf("fail: x=%d; hope = %t; ret = %t", x, hope, ret)
	}

	x = 0
	hope = true
	ret = isPalindrome(x)
	if ret == hope {
		t.Logf("success: x=%d; hope = %t; ret = %t", x, hope, ret)
	} else {
		t.Errorf("fail: x=%d; hope = %t; ret = %t", x, hope, ret)
	}

	x = 1
	hope = true
	ret = isPalindrome(x)
	if ret == hope {
		t.Logf("success: x=%d; hope = %t; ret = %t", x, hope, ret)
	} else {
		t.Errorf("fail: x=%d; hope = %t; ret = %t", x, hope, ret)
	}
}