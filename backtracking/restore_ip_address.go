//给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
//
//示例:
//
//输入: "25525511135"
//输出: ["255.255.11.135", "255.255.111.35"]

package backtracking

func restoreIpAddresses(s string) []string {
	r := make([]string, 0)
	var f func(string, string, int, int)
	f = func(s, ip string, n, l int) {
	    if l == 0 {
	        return
        }
        if n == 3 {
            if l == 1 ||
                (s[0] != '0' &&
                    (l == 2 ||
                        (l == 3 &&
                            s[0] != '0' && int(s[0]-'0')*100 + int(s[1]-'0')*10 + int(s[2]-'0') < 256))) {
                r = append(r, ip+s)
            }
            return
        }
	    if l > 1 {
            f(s[1: ], ip+ s[: 1] + ".", n+1, l-1)
            if s[0] != '0' {
                if l > 2 {
                    f(s[2: ], ip+ s[0: 2] + ".", n+1, l-2)
                }
                if l > 3 && int(s[0]-'0')*100 + int(s[1]-'0')*10 + int(s[2]-'0') < 256 {
                    f(s[3: ], ip+ string(s[: 3]) + ".", n+1, l-3)
                }
            }
        }
	}
	f(s, "", 0, len(s))
	return r
}
