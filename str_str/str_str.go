package str_str

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	if needle == "" {
		return 0
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i] != needle[0] {
			continue
		}
		if i+len(needle) == len(haystack) {
			fmt.Printf("14: str = %s\n", haystack[i:])
			if haystack[i:] == needle {
				return i
			}
			return -1
		}
		if haystack[i:i+len(needle)] == needle {
			fmt.Printf("21: str = %s\n", haystack[i:i+len(needle)])
			return i
		}
	}
	return -1
}
