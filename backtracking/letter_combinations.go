// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
// 示例:
// 输入："23"
// 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].

package backtracking

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	letterMap := [][]string{
		{"a", "b", "c"},
		{"d", "e", "f"},
		{"g", "h", "i"},
		{"j", "k", "l"},
		{"m", "n", "o"},
		{"p", "q", "r", "s"},
		{"t", "u", "v"},
		{"w", "x", "y", "z"},
	}
	ret := []string{}
	dfs(digits, "", &ret, &letterMap)
	return ret
}

func dfs(digits, str string, ret *[]string, letterMap *[][]string) {
	if len(digits) == 0 {
		*ret = append(*ret, str)
	} else {
		idx := digits[0] - '2'
		for i := range (*letterMap)[idx] {
			newStr := str + (*letterMap)[idx][i]
			if len(digits) == 1 {
				dfs("", newStr, ret, letterMap)
			} else {
				dfs(digits[1:], newStr, ret, letterMap)
			}
		}
	}
}

// func letterCombinations(digits string) []string {
// 	if len(digits) == 0 {
// 		return []string{}
// 	}
// 	letterMap := map[string][]string{
// 		"2": []string{"a", "b", "c"},
// 		"3": []string{"d", "e", "f"},
// 		"4": []string{"g", "h", "i"},
// 		"5": []string{"j", "k", "l"},
// 		"6": []string{"m", "n", "o"},
// 		"7": []string{"p", "q", "r", "s"},
// 		"8": []string{"t", "u", "v"},
// 		"9": []string{"w", "x", "y", "z"},
// 	}
// 	strs := []string{""}
// 	var i int
// 	for ; i < len(digits); i++ {
// 		strs_new := []string{}
// 		letters := letterMap[string(digits[i])]
// 		for _, letter := range letters {
// 			for _, str := range strs {
// 				str += letter
// 				strs_new = append(strs_new, str)
// 			}
// 		}
// 		strs = strs_new
// 	}
// 	return strs
// }
