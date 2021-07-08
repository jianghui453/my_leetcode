package backtracking

import (
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	var digits string
	var h, r []string

	digits = "23"
	h = []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	r = letterCombinations(digits)
	t.Logf("digits=%v h=%v r=%v", digits, h, r)
}
