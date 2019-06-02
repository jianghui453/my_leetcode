package leet_code

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	var x, xc, ret int

	x = 123
	xc = 321
	ret = reverse(x)
	if ret == xc {
		t.Log(fmt.Sprintf("success: x=%d; xc=%d; ret=%d\n", x, xc, ret))
	} else {
		t.Error(fmt.Sprintf("fail: x=%d; xc=%d; ret=%d\n", x, xc, ret))
	}

	x = -123
	xc = -321
	ret = reverse(x)
	if ret == xc {
		t.Log(fmt.Sprintf("success: x=%d; xc=%d; ret=%d\n", x, xc, ret))
	} else {
		t.Error(fmt.Sprintf("fail: x=%d; xc=%d; ret=%d\n", x, xc, ret))
	}
}

func TestReverseV2(t *testing.T) {
	var x, xc, ret int

	x = 123
	xc = 321
	ret = reverseV2(x)
	if ret == xc {
		t.Log(fmt.Sprintf("success: x=%d; xc=%d; ret=%d\n", x, xc, ret))
	} else {
		t.Error(fmt.Sprintf("fail: x=%d; xc=%d; ret=%d\n", x, xc, ret))
	}

	x = -123
	xc = -321
	ret = reverseV2(x)
	if ret == xc {
		t.Log(fmt.Sprintf("success: x=%d; xc=%d; ret=%d\n", x, xc, ret))
	} else {
		t.Error(fmt.Sprintf("fail: x=%d; xc=%d; ret=%d\n", x, xc, ret))
	}
}