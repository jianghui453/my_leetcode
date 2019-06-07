package leet_code

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	var s string
	var hope, ret int

	s = "III"
	hope = 3
	ret = romanToInt(s);
	if hope == ret {
		t.Logf("success: s = %v; hope = %d; ret = %d\n", s, hope, ret)
	} else {
		t.Errorf("fail: s = %v; hope = %d; ret = %d\n", s, hope, ret)
	}

	s = "IV"
	hope = 4
	ret = romanToInt(s);
	if hope == ret {
		t.Logf("success: s = %v; hope = %d; ret = %d\n", s, hope, ret)
	} else {
		t.Errorf("fail: s = %v; hope = %d; ret = %d\n", s, hope, ret)
	}

	s = "IX"
	hope = 9
	ret = romanToInt(s);
	if hope == ret {
		t.Logf("success: s = %v; hope = %d; ret = %d\n", s, hope, ret)
	} else {
		t.Errorf("fail: s = %v; hope = %d; ret = %d\n", s, hope, ret)
	}

	s = "LVIII"
	hope = 58
	ret = romanToInt(s);
	if hope == ret {
		t.Logf("success: s = %v; hope = %d; ret = %d\n", s, hope, ret)
	} else {
		t.Errorf("fail: s = %v; hope = %d; ret = %d\n", s, hope, ret)
	}

	s = "MCMXCIV"
	hope = 1994
	ret = romanToInt(s);
	if hope == ret {
		t.Logf("success: s = %v; hope = %d; ret = %d\n", s, hope, ret)
	} else {
		t.Errorf("fail: s = %v; hope = %d; ret = %d\n", s, hope, ret)
	}
}