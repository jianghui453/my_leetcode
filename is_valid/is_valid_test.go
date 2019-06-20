package leet_code

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	var s string
	var hope, ret bool

	s = "()[]{}"
	hope = true
	ret = isValid(s)
	if hope == ret {
		t.Logf("success: s = %s; hope = %t; ret = %t\v", s, hope, ret)
	} else {
		t.Errorf("error: s = %s; hope = %t; ret = %t\v", s, hope, ret)
	}

	s = "{[]}"
	hope = true
	ret = isValid(s)
	if hope == ret {
		t.Logf("success: s = %s; hope = %t; ret = %t\v", s, hope, ret)
	} else {
		t.Errorf("error: s = %s; hope = %t; ret = %t\v", s, hope, ret)
	}

	s = "([)]"
	hope = false
	ret = isValid(s)
	if hope == ret {
		t.Logf("success: s = %s; hope = %t; ret = %t\v", s, hope, ret)
	} else {
		t.Errorf("error: s = %s; hope = %t; ret = %t\v", s, hope, ret)
	}
}