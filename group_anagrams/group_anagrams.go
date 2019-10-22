package group_anagrams

import "fmt"
import "sort"
import "strings"

func groupAnagrams(strs []string) [][]string {
	groups := []string{}
	ret := [][]string{}
Loop:
	for _, str := range strs {
		_str := strings.Split(str, "")
		sort.Strings(_str)
		__str := strings.Join(_str, "")
		for idx, group := range groups {
			if group == __str {
				ret[idx] = append(ret[idx], str)
				continue Loop
			}
		}
		groups = append(groups, __str)
		ret = append(ret, []string{str})
	}
	return ret
}

func groupAnagramsv1(strs []string) [][]string {
	groups := [][]byte{}
	ret := [][]string{}
Loop:
	for _, str := range strs {
		for idx, group := range groups {
			if equal(group, str) {
				ret[idx] = append(ret[idx], str)
				continue Loop
			}
		}
		groups = append(groups, []byte(str))
		ret = append(ret, []string{str})
	}
	return ret
}

func equal(chs []byte, str string) bool {
	fmt.Printf("chs=%s str=%s\n", string(chs), str)
	sLen := len(str)
	if sLen != len(chs) {
		return false
	}
	if sLen == 0 {
		return true
	}
	if sLen == 1 {
		return str[0] == chs[0]
	}
	ch := str[0]
	_str := str[1:]
	for idx, c := range chs {
		if ch != c {
			continue
		}
		var _chs []byte
		if idx == sLen-1 {
			_chs = make([]byte, sLen-1)
			copy(_chs, chs[:sLen-1])
		} else {
			_chs = make([]byte, idx)
			copy(_chs, chs[:idx])
			_chs = append(_chs, chs[idx+1:]...)
		}
		fmt.Printf("_chs=%s _str=%s\n", string(_chs), _str)
		if equal(_chs, _str) {
			return true
		}
	}
	return false
}
