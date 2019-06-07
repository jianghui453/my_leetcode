package leet_code

import (
    "fmt"
)

var memo = make(map[string]map[string]bool)

var Memo = &memo

func isMatch(s string, p string) bool {
    if ret, ok := memo[s][p]; ok {
        return ret
    }
    fmt.Printf("s = %s; p = %s; memo = %v\n", s, p, memo)
    memo[s] = make(map[string]bool)
    if len(p) == 0 {
        memo[s][p] = len(s) == 0
        return memo[s][p]
    }
    if len(s) == 0 {
        memo[s][p] = (len(p) == 2 && p[1] == '*') || (len(p) > 2 && p[1] == '*' && isMatch("", p[2:]))
        return memo[s][p]
    }
    // len(p) > 0 && len(s) > 0
    firstMatch := p[0] == s[0] || p[0] == '.'
    var s1, p1, p2 string
    if len(s) == 1 {
        s1 = ""
    } else {
        s1 = s[1:]
    }
    if len(p) == 1 {
        p1 = ""
    } else {
        p1 = p[1:]
    }
    if len(p) >= 2 && p[1] == '*' {
        if len(p) == 2 {
            p2 = ""
        } else {
            p2 = p[2:]
        }
        memo[s][p] = isMatch(s, p2) || (firstMatch && isMatch(s1, p2)) || (firstMatch && isMatch(s1, p))
        return  memo[s][p]
    }
    memo[s][p] = firstMatch && isMatch(s1, p1)
    return memo[s][p]
}

// func isMatchV2(s, p string) bool {

// }
