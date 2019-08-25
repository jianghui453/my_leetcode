package my_std

import (
	"testing"
)

func TestMyAtoI(t *testing.T) {
	var str string
	var ret, hope int

	str = "4193 with words"
	hope = 4193
	ret = myAtoi(str)
	if ret == hope {
		t.Logf("success: str=%s; hope=%d; ret=%d", str, hope, ret)
	} else {
		t.Errorf("fail: str=%s; hope=%d; ret=%d", str, hope, ret)
	}

	str = "42"
	hope = 42
	ret = myAtoi(str)
	if ret == hope {
		t.Logf("success: str=%s; hope=%d; ret=%d", str, hope, ret)
	} else {
		t.Errorf("fail: str=%s; hope=%d; ret=%d", str, hope, ret)
	}

	str = " -42"
	hope = -42
	ret = myAtoi(str)
	if ret == hope {
		t.Logf("success: str=%s; hope=%d; ret=%d", str, hope, ret)
	} else {
		t.Errorf("fail: str=%s; hope=%d; ret=%d", str, hope, ret)
	}

	str = "words and 987"
	hope = 0
	ret = myAtoi(str)
	if ret == hope {
		t.Logf("success: str=%s; hope=%d; ret=%d", str, hope, ret)
	} else {
		t.Errorf("fail: str=%s; hope=%d; ret=%d", str, hope, ret)
	}

	str = "-91283472332"
	hope = -2147483648
	ret = myAtoi(str)
	if ret == hope {
		t.Logf("success: str=%s; hope=%d; ret=%d", str, hope, ret)
	} else {
		t.Errorf("fail: str=%s; hope=%d; ret=%d", str, hope, ret)
	}

	str = "+-2"
	hope = 0
	ret = myAtoi(str)
	if ret == hope {
		t.Logf("success: str=%s; hope=%d; ret=%d", str, hope, ret)
	} else {
		t.Errorf("fail: str=%s; hope=%d; ret=%d", str, hope, ret)
	}

	str = "-+2"
	hope = 0
	ret = myAtoi(str)
	if ret == hope {
		t.Logf("success: str=%s; hope=%d; ret=%d", str, hope, ret)
	} else {
		t.Errorf("fail: str=%s; hope=%d; ret=%d", str, hope, ret)
	}

	str = "  +0 123"
	hope = 0
	ret = myAtoi(str)
	if ret == hope {
		t.Logf("success: str=%s; hope=%d; ret=%d", str, hope, ret)
	} else {
		t.Errorf("fail: str=%s; hope=%d; ret=%d", str, hope, ret)
	}

	str = "9223372036854775808"
	hope = 2147483647
	ret = myAtoi(str)
	if ret == hope {
		t.Logf("success: str=%s; hope=%d; ret=%d", str, hope, ret)
	} else {
		t.Errorf("fail: str=%s; hope=%d; ret=%d", str, hope, ret)
	}
}
