package leet_code

import (
	"testing"
)

func TestIntToRoman(t *testing.T) {
	var num int
	var hope, ret string

	num = 3
	hope = "III"
	ret = intToRoman(num)
	if ret == hope {
		t.Logf("success: num = %d; hope = %s; ret = %s\n", num, hope, ret)
	} else {
		t.Errorf("fail: num = %d; hope = %s; ret = %s\n", num, hope, ret)
	}

	num = 4
	hope = "IV"
	ret = intToRoman(num)
	if ret == hope {
		t.Logf("success: num = %d; hope = %s; ret = %s\n", num, hope, ret)
	} else {
		t.Errorf("fail: num = %d; hope = %s; ret = %s\n", num, hope, ret)
	}

	num = 58
	hope = "LVIII"
	ret = intToRoman(num)
	if ret == hope {
		t.Logf("success: num = %d; hope = %s; ret = %s\n", num, hope, ret)
	} else {
		t.Errorf("fail: num = %d; hope = %s; ret = %s\n", num, hope, ret)
	}

	num = 1994
	hope = "MCMXCIV"
	ret = intToRoman(num)
	if ret == hope {
		t.Logf("success: num = %d; hope = %s; ret = %s\n", num, hope, ret)
	} else {
		t.Errorf("fail: num = %d; hope = %s; ret = %s\n", num, hope, ret)
	}
}