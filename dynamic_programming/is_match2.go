package dynamic_programming

import (
	// "strings"
	// "fmt"
)

func isMatch(s string, p string) bool {
	sLen := len(s)
	pLen := len(p)

	if pLen == 0 {
		return sLen == 0
	}

	if sLen == 0 {
		for i := 0; i < pLen; i++ {
			if p[i] == '*' {
				continue
			}

			return false
		}

		return true
	}

	dp := make([][]bool, sLen+1)
	for i := 0; i <= sLen; i++ {
		dp[i] = make([]bool, pLen+1)
	}

	dp[0][0] = true
	for i := 0; i < pLen; i++ {
		if p[i] == '*' {
			dp[0][i+1] = true
		} else {
			break
		}
	}

	for i := 0; i < sLen; i++ {
		for j := 0; j < pLen; j++ {
			// fmt.Printf("i=%d j=%d\n", i, j)
			if (dp[i][j] && (s[i] == p[j] || p[j] == '?' || p[j] == '*')) ||
				(dp[i][j+1] && p[j] == '*') ||
				(dp[i+1][j] && p[j] == '*') {
				dp[i+1][j+1] = true
			}
				
		}
	}

	// fmt.Printf("dp=%v\n", dp)

	return dp[sLen][pLen]
}

// func isMatch(s string, p string) bool {
// 	if len(p) > 1 {
// 		sb := strings.Builder{}
// 		sb.WriteByte(p[0])
// 		for i := 1; i < len(p); i ++ {
// 			if p[i] != '*' || p[i - 1] != '*' {
// 				sb.WriteByte(p[i])
// 			}
// 		}
// 		p = sb.String()
// 	}
// 	rpt := make(map[string]bool)
// 	return match2(s, p, &rpt)
// }

// func match2(s, p string, rpt *map[string]bool) bool {
// 	// fmt.Printf("s = %s, p = %s.\n", s, p)
// 	if len(s) == 0 {
// 		if len(p) == 0 {
// 			return true
// 		} else {
// 			for _, pch := range p {
// 				if pch != '*' {
// 					return false
// 				}
// 			}
// 			return true
// 		}
		
// 	}
// 	if len(p) == 0 {
// 		return false
// 	}
// 	key := s + "_" + p
// 	if _, ok := (*rpt)[key]; ok {
// 		return (*rpt)[key]
// 	}
// 	var s0, p0 string
// 	if len(s) == 1 {
// 		s0 = ""
// 	} else {
// 		s0 = s[1:]
// 	}
// 	if len(p) == 1 {
// 		p0 = ""
// 	} else {
// 		p0 = p[1:]
// 	}
// 	if s[0] == p[0] || p[0] == '?' {
// 		(*rpt)[key] = match2(s0, p0, rpt)
// 		return (*rpt)[key]
// 	}
// 	if p[0] == '*' {
// 		(*rpt)[key] = match2(s0, p0, rpt) || match2(s, p0, rpt) || match2(s0, p, rpt)
// 		return (*rpt)[key]
// 	}
// 	return false
// }

// func isMatch2(s string, p string) bool {
// 	rptMap := make(map[string]bool)
// 	return match(s, p, &rptMap)
// }

// func match(s, p string, rptMap *map[string]bool) bool {
// 	key := s + "_" + p
// 	if val, ok := (*rptMap)[key]; ok {
// 		fmt.Printf("rpt: key = %s, val = %t.\n", key, val)
// 		return val
// 	}
// 	var ret bool
// 	if len(s) == 0 {
// 		if len(p) == 0 {
// 			return true
// 		} else if len(p) == 1 {
// 			return p[0] == '*'
// 		} else if p[0] == '*' {
// 			return match(s, p[1:], rptMap)
// 		} else {
// 			return false
// 		}
// 	}
// 	if len(p) == 0 {
// 		return false
// 	}
// 	var s0, p0 string
// 	if len(s) == 1 {
// 		s0 = ""
// 	} else {
// 		s0 = s[1:]
// 	}
// 	if len(p) == 1 {
// 		p0 = ""
// 	} else {
// 		p0 = p[1:]
// 	}
// 	if s[0] == p[0] || p[0] == '?' {
// 		ret = match(s0, p0, rptMap)
// 		(*rptMap)[key] = ret
// 		return ret
// 	} else if p[0] == '*' {
// 		ret = match(s0, p0, rptMap) || match(s, p0, rptMap) || match(s0, p, rptMap)
// 		(*rptMap)[key] = ret
// 		return ret
// 	} else {
// 		ret = false
// 		(*rptMap)[key] = ret
// 		return ret
// 	}
// }
