package leet_code

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T)  {
	var s, ret, sRet string
	var numRows int

	s = "LEETCODEISHIRING"
	numRows = 3
	ret = "LCIRETOESIIGEDHN"
	sRet = convert(s, numRows)
	if ret == sRet {
		t.Log(fmt.Sprintf("success: s=%s; numRows=%d; hope %s get %s", s, numRows, ret, sRet))
	} else {
		t.Error(fmt.Sprintf("fail: s=%s; numRows=%d; hope %s get %s", s, numRows, ret, sRet))
	}

	s = "LEETCODEISHIRING"
	numRows = 4
	ret = "LDREOEIIECIHNTSG"
	sRet = convert(s, numRows)
	if ret == sRet {
		t.Log(fmt.Sprintf("success: s=%s; numRows=%d; hope %s get %s", s, numRows, ret, sRet))
	} else {
		t.Error(fmt.Sprintf("fail: s=%s; numRows=%d; hope %s get %s", s, numRows, ret, sRet))
	}
}

func TestConvertV2(t *testing.T) {
	var s, ret, sRet string
	var numRows int

	s = "LEETCODEISHIRING"
	numRows = 3
	ret = "LCIRETOESIIGEDHN"
	sRet = convertV2(s, numRows)
	if ret == sRet {
		t.Log(fmt.Sprintf("success: s=%s; numRows=%d; hope %s get %s", s, numRows, ret, sRet))
	} else {
		t.Error(fmt.Sprintf("fail: s=%s; numRows=%d; hope %s get %s", s, numRows, ret, sRet))
	}

	s = "LEETCODEISHIRING"
	numRows = 4
	ret = "LDREOEIIECIHNTSG"
	sRet = convertV2(s, numRows)
	if ret == sRet {
		t.Log(fmt.Sprintf("success: s=%s; numRows=%d; hope %s get %s", s, numRows, ret, sRet))
	} else {
		t.Error(fmt.Sprintf("fail: s=%s; numRows=%d; hope %s get %s", s, numRows, ret, sRet))
	}
}