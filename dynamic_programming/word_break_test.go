package dynamic_programming

import "testing"

func TestWordBreak(t *testing.T) {
	var s string
	var wordDict []string
	//var h, r bool

	//s = "leetcode"
	//wordDict = []string{"leet", "code"}
	//h = true
	//r = wordBreak(s, wordDict)
	//t.Logf("s=%s wordDict=%v h=%t r=%t", s, wordDict, h, r)
	//
	//s = "applepenapple"
	//wordDict = []string{"apple", "pen"}
	//h = true
	//r = wordBreak(s, wordDict)
	//t.Logf("s=%s wordDict=%v h=%t r=%t", s, wordDict, h, r)
	//
	//s = "catsandog"
	//wordDict = []string{"cats", "dog", "sand", "and", "cat"}
	//h = false
	//r = wordBreak(s, wordDict)
	//t.Logf("s=%s wordDict=%v h=%t r=%t", s, wordDict, h, r)

	//var h, r []string
    //
	//s = "catsanddog"
	//wordDict = []string{"cat", "cats", "and", "sand", "dog"}
	//h = []string{"cats and dog", "cat sand dog"}
	//r = wordBreak(s, wordDict)
	//t.Logf("s=%s wordDict=%v \nh=%s \nr=%s", s, wordDict, h, r)

	//s = "pineapplepenapple"
	//wordDict = []string{"apple", "pen", "applepen", "pine", "pineapple"}
	//h = []string{"pine apple pen apple", "pineapple pen apple", "pine applepen apple"}
	//r = wordBreak(s, wordDict)
	//t.Logf("s=%s wordDict=%v \nh=%s \nr=%s", s, wordDict, h, r)
	//
	//s = "catsandog"
	//wordDict = []string{"cats", "dog", "sand", "and", "cat"}
	//h = []string{}
	//r = wordBreak(s, wordDict)
	//t.Logf("s=%s wordDict=%v \nh=%s \nr=%s", s, wordDict, h, r)
	//
	//s = "ab"
	//wordDict = []string{"a", "b"}
	//h = []string{"a", "b"}
	//r = wordBreak(s, wordDict)
	//t.Logf("s=%s wordDict=%v \nh=%s \nr=%s", s, wordDict, h, r)

	//s = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	//s = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
    s = "aaaaaaaaaaaaaaaaaaaaaaaaaaa"
	//s = "aaaaaaaaaaaaaaaaaa"
    //
	wordDict = []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	//h = []string{"a", "b"}
    wordBreak(s, wordDict)
	//t.Logf("s=%s wordDict=%v \nh=%s \nr=%s", s, wordDict, h, r)
}
