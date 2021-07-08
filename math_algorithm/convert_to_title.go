package math_algorithm

import (
	// "fmt"
)

func convertToTitle(columnNumber int) string {
	alphabet := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N',
		'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

	ret := []byte{}
	var reminder int
	for columnNumber > 0 {
		columnNumber--
		columnNumber, reminder = columnNumber / len(alphabet), columnNumber % len(alphabet)
		ret = append(ret, alphabet[reminder])
	}
	for i := 0; i < len(ret)/2; i++ {
		ret[i], ret[len(ret)-1-i] = ret[len(ret)-1-i], ret[i]
	}
	return string(ret)
}
