package leet_code

import (
	"testing"
)

func TestIsMatch(t *testing.T) {
	var s, p string
	var ret, hope bool

	s = "aa"
	p = "a"
	hope = false
	ret = isMatch(s, p)
	if ret == hope {
		t.Logf("success: s=%s; p = %s; hope = %t; ret = %t", s, p, hope, ret)
	} else {
		t.Errorf("fail: s=%s; p = %s; hope = %t; ret = %t", s, p, hope, ret)
	}

	s = "aa"
	p = "a*"
	hope = true
	ret = isMatch(s, p)
	if ret == hope {
		t.Logf("success: s=%s; p = %s; hope = %t; ret = %t", s, p, hope, ret)
	} else {
		t.Errorf("fail: s=%s; p = %s; hope = %t; ret = %t", s, p, hope, ret)
	}

	s = "ab"
	p = ".*"
	hope = true
	ret = isMatch(s, p)
	if ret == hope {
		t.Logf("success: s=%s; p = %s; hope = %t; ret = %t", s, p, hope, ret)
	} else {
		t.Errorf("fail: s=%s; p = %s; hope = %t; ret = %t", s, p, hope, ret)
	}

	s = "aab"
	p = "c*a*b"
	hope = true
	ret = isMatch(s, p)
	if ret == hope {
		t.Logf("success: s=%s; p = %s; hope = %t; ret = %t", s, p, hope, ret)
	} else {
		t.Errorf("fail: s=%s; p = %s; hope = %t; ret = %t", s, p, hope, ret)
	}

	s = "mississippi"
	p = "mis*is*p*."
	hope = false
	ret = isMatch(s, p)
	if ret == hope {
		t.Logf("success: s=%s; p = %s; hope = %t; ret = %t", s, p, hope, ret)
	} else {
		t.Errorf("fail: s=%s; p = %s; hope = %t; ret = %t", s, p, hope, ret)
	}

	s = "aaa"
	p = "aaaa"
	hope = false
	ret = isMatch(s, p)
	if ret == hope {
		t.Logf("success: s=%s; p = %s; hope = %t; ret = %t", s, p, hope, ret)
	} else {
		t.Errorf("fail: s=%s; p = %s; hope = %t; ret = %t", s, p, hope, ret)
	}

	s = "ab"
	p = ".*c"
	hope = false
	ret = isMatch(s, p)
	if ret == hope {
		t.Logf("success: s=%s; p = %s; hope = %t; ret = %t", s, p, hope, ret)
	} else {
		t.Errorf("fail: s=%s; p = %s; hope = %t; ret = %t", s, p, hope, ret)
	}

	s = "aaa"
	p = "a*a"
	hope = true
	ret = isMatch(s, p)
	if ret == hope {
		t.Logf("success: s=%s; p = %s; hope = %t; ret = %t", s, p, hope, ret)
	} else {
		t.Errorf("fail: s=%s; p = %s; hope = %t; ret = %t", s, p, hope, ret)
	}

	s = "a"
	p = "ab*"
	hope = true
	ret = isMatch(s, p)
	if ret == hope {
		t.Logf("success: s=%s; p = %s; hope = %t; ret = %t", s, p, hope, ret)
	} else {
		t.Errorf("fail: s=%s; p = %s; hope = %t; ret = %t", s, p, hope, ret)
	}

	s = "a"
	p = "ab*b*"
	hope = true
	ret = isMatch(s, p)
	if ret == hope {
		t.Logf("success: s=%s; p = %s; hope = %t; ret = %t", s, p, hope, ret)
	} else {
		t.Errorf("fail: s=%s; p = %s; hope = %t; ret = %t", s, p, hope, ret)
	}

	s = "a"
	p = ".*..a*"
	hope = false
	ret = isMatch(s, p)
	if ret == hope {
		t.Logf("success: s=%s; p = %s; hope = %t; ret = %t", s, p, hope, ret)
	} else {
		t.Errorf("fail: s=%s; p = %s; hope = %t; ret = %t", s, p, hope, ret)
	}
	t.Logf("memo = %v", Memo)
}
