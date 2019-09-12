package backtracking

import "testing"

func TestRestoreIPAddress(t *testing.T) {
	var s string
	var h, r []string

	s = "25525511135"
	h = []string{"255.255.11.135", "255.255.111.35"}
	r = restoreIpAddresses(s)
	t.Logf("s=%s \nh=%v \nr=%v", s, h, r)

    s = "0000"
    h = []string{"0.0.0.0"}
    r = restoreIpAddresses(s)
    t.Logf("s=%s \nh=%v \nr=%v", s, h, r)

    s = "000"
    h = []string{}
    r = restoreIpAddresses(s)
    t.Logf("s=%s \nh=%v \nr=%v", s, h, r)
}
