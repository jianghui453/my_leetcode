package leet_code

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	letterMap := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}
	strs := []string{""}
	var i int
	for ; i < len(digits); i++ {
		strs_new := []string{}
		letters := letterMap[string(digits[i])]
		for _, letter := range letters {
			for _, str := range strs {
				str += letter
				strs_new = append(strs_new, str)
			}
		}
		strs = strs_new
	}
	return strs
}
