package greedy

import (
	"testing"
)

func TestIntToRoman(t *testing.T) {
	var num int
	var h, r string

	num = 3
	h = "III"
	r = intToRoman(num)
	t.Logf("%t num=%d h=%s r=%s", r == h, num, h, r)

	num = 4
	h = "IV"
	r = intToRoman(num)
	t.Logf("%t num=%d h=%s r=%s", r == h, num, h, r)

	num = 58
	h = "LVIII"
	r = intToRoman(num)
	t.Logf("%t num=%d h=%s r=%s", r == h, num, h, r)

	num = 1994
	h = "MCMXCIV"
	r = intToRoman(num)
	t.Logf("%t num=%d h=%s r=%s", r == h, num, h, r)
}
