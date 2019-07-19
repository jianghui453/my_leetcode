package group_anagrams

func groupAnaprams(strings []string) [][]string {
	bytess := [][]byte{}
	rets := [][]string{}
	for _, str := range strings {
		sLen := len(str)
		for idx, bytes := range bytess {
			if len(bytes) != sLen {
				continue
			}
			for _, ch := range bytes {
				
			}
		}
		for i := 0; i < sLen; i ++ {

		}
	}
}