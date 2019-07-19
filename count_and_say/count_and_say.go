package count_and_say

import "strings"
import "fmt"

func countAndSay(n int) string {
    if n < 1 {
		return ""
	}
	str := "1"
	for ; n > 1; n -- {
		var i int
		var count int
		sb := strings.Builder{}
		for _, ch := range str {
			if i == 0 {
				i = int(ch - '0')
				count ++
				continue
			}
			if i == int(ch - '0') {
				count ++
				continue
			}
			fmt.Printf("1count = %d, i = %d, str = %s.\n", count, i, str)
			sb.Write([]byte{byte(count + '0')})
			sb.Write([]byte{byte(i + '0')})
			i = int(ch - '0')
			count = 1
		}
		if count > 0 {
			fmt.Printf("2count = %d, i = %d, str = %s.\n", count, i, str)
			sb.Write([]byte{byte(count + '0')})
			sb.Write([]byte{byte(i + '0')})
		}
		str = sb.String()
	}
	return str
}