package hash_table

// 187. 重复的DNA序列
func findRepeatedDnaSequences(s string) []string {
	if len(s) <= 10 {
		return []string{}
	}

	strs := make([]string, 0)
	hash := make(map[string]int)
	for i := 0; i < len(s) - 9; i++ {
		if hash[s[i: i+10]] == 0 {
			hash[s[i: i+10]] = 1
		} else if hash[s[i: i+10]] == 1 {
			hash[s[i: i+10]] = 2
			strs = append(strs, s[i: i + 10])
		}
	}

	return strs
}
