package length_of_last_word

func lengthOfLastWord(s string) int {
	ret := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if ret == 0 {
				continue
			} else {
				break
			}
		}
		ret++
	}
	return ret
}
