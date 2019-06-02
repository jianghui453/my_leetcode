package leet_code

import (
    "fmt"
)

func isMatch(s string, p string) (ret bool) {
    defer fmt.Printf("s = %s; p = %s; ret = %v\n", s, p, ret)
    if len(p) == 0 {
        ret = len(s) == 0
        return
    }
    p_zero := len(p) == 0 || (len(p) == 2 && p[1] == '*')
    if len(s) == 0 {
        ret = p_zero
        return
    }
    first_match := len(s) >=1 && (p[0] == s[0] || p[0] == '.')

    if len(p) >= 2 && p[1] == '*' {
        var s1 string
        if len(s) >= 1 {
            s1 = s[1:]
        } else {
            s1 = ""
        }
        return isMatch(s, p) || (first_match && isMatch(s1, p))
    }
    if len(s) > 1 {
        s = s[1:]   
    } else {
        s = ""
    }
    if len(p) > 1 {
        p = p[1:]
    } else {
        p = ""
    }
    return first_match && isMatch(s, p)
}
