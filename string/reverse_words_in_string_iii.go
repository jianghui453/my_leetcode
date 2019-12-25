// 给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。

// 示例 1:

// 输入: "Let's take LeetCode contest"
// 输出: "s'teL ekat edoCteeL tsetnoc" 
// 注意：在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。

package string

func reverseWords(s string) string {
    var (
		ret []byte
		word []byte
		l int = len(s)
	)

	for i := 0; i < l; i++ {
		if s[i] == ' ' {
			if len(word) > 0 {
				for i := len(word)-1; i >= 0; i-- {
					ret = append(ret, word[i])
				}
				word = word[: 0]
			}
			ret = append(ret, ' ')
		} else {
			word = append(word, s[i])
		}
	}

	for i := len(word)-1; i >= 0; i-- {
		ret = append(ret, word[i])
	}

	return string(ret)
}