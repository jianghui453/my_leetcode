package backtracking

import "testing"

func TestMinCut(t *testing.T) {
	var s string
	var h, r int

	s = "aab"
	h = 1
	r = minCut(s)
	t.Logf("%t s=%s h=%d r=%d", r == h, s, h, r)

	s = "abc"
	h = 2
	r = minCut(s)
	t.Logf("%t s=%s h=%d r=%d", r == h, s, h, r)

	s = "aaa"
	h = 0
	r = minCut(s)
	t.Logf("%t s=%s h=%d r=%d", r == h, s, h, r)

	s = "a"
	h = 0
	r = minCut(s)
	t.Logf("%t s=%s h=%d r=%d", r == h, s, h, r)

	s = "abbab"
	h = 1
	r = minCut(s)
	t.Logf("%t s=%s h=%d r=%d", r == h, s, h, r)

	s = "aaabaa"
	h = 1
	r = minCut(s)
	t.Logf("%t s=%s h=%d r=%d", r == h, s, h, r)

	s = "fifgbeajcacehiicccfecbfhhgfiiecdcjjffbghdidbhbdbfbfjccgbbdcjheccfbhafehieabbdfeigbiaggchaeghaijfbjhi"
	h = 75
	r = minCut(s)
	t.Logf("%t s=%s h=%d r=%d", r == h, s, h, r)

	s = "apjesgpsxoeiokmqmfgvjslcjukbqxpsobyhjpbgdfruqdkeiszrlmtwgfxyfostpqczidfljwfbbrflkgdvtytbgqalguewnhvvmcgxboycffopmtmhtfizxkmeftcucxpobxmelmjtuzigsxnncxpaibgpuijwhankxbplpyejxmrrjgeoevqozwdtgospohznkoyzocjlracchjqnggbfeebmuvbicbvmpuleywrpzwsihivnrwtxcukwplgtobhgxukwrdlszfaiqxwjvrgxnsveedxseeyeykarqnjrtlaliyudpacctzizcftjlunlgnfwcqqxcqikocqffsjyurzwysfjmswvhbrmshjuzsgpwyubtfbnwajuvrfhlccvfwhxfqthkcwhatktymgxostjlztwdxritygbrbibdgkezvzajizxasjnrcjwzdfvdnwwqeyumkamhzoqhnqjfzwzbixclcxqrtniznemxeahfozp"
	h = 452
	r = minCut(s)
	t.Logf("%t s=%s h=%d r=%d", r == h, s, h, r)
}
