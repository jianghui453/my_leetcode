package leet_code

import (
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	var digits string
	var hope, ret []string

	digits = "23"
	hope = []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	ret = letterCombinations(digits)
	t.Logf("digits = %v; hope = %v; ret = %v", digits, hope, ret)
}
