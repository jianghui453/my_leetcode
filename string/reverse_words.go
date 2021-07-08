package string

import "strings"

func reverseWords151(s string) string {
	s = strings.TrimSpace(s)
	strs := strings.Split(s, " ")
	rWords := []string{}
	for _, str := range strs {
		if str != "" {
			rWords = append(rWords, str)
		}
	}
	for i := 0; i < len(rWords) / 2; i++ {
		rWords[i], rWords[len(rWords)-1-i] = rWords[len(rWords)-1-i], rWords[i]
	}
	return strings.Join(rWords, " ")
}
