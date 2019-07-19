package is_match

import "fmt"

func isMatch2(s string, p string) bool {
	rptMap := make(map[string]bool)
	return match(s, p, &rptMap)
}

func match(s, p string, rptMap *map[string]bool) bool {
	key := s + "_" + p
	if val, ok := (*rptMap)[key]; ok {
		fmt.Printf("rpt: key = %s, val = %t.\n", key, val)
		return val
	}
	var ret bool
	if len(s) == 0 {
		if len(p) == 0 {
			return true
		} else if len(p) == 1 {
			return p[0] == '*'
		} else if p[0] == '*' {
			return match(s, p[1:], rptMap)
		} else {
			return false
		}
	}
	if len(p) == 0 {
		return false
	}
	var s0, p0 string
	if len(s) == 1 {
		s0 = ""
	} else {
		s0 = s[1:]
	}
	if len(p) == 1 {
		p0 = ""
	} else {
		p0 = p[1:]
	}
	if s[0] == p[0] || p[0] == '?' {
		ret = match(s0, p0, rptMap)
		(*rptMap)[key] = ret
		return ret
	} else if p[0] == '*' {
		ret = match(s0, p0, rptMap) || match(s, p0, rptMap) || match(s0, p, rptMap)
		(*rptMap)[key] = ret
		return ret
	} else {
		ret = false
		(*rptMap)[key] = ret
		return ret
	}
}