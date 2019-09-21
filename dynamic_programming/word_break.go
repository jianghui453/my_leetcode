//给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
//
//说明：
//
//拆分时可以重复使用字典中的单词。
//你可以假设字典中没有重复的单词。
//示例 1：
//
//输入: s = "leetcode", wordDict = ["leet", "code"]
//输出: true
//解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
//示例 2：
//
//输入: s = "applepenapple", wordDict = ["apple", "pen"]
//输出: true
//解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
//     注意你可以重复使用字典中的单词。
//示例 3：
//
//输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
//输出: false

package dynamic_programming

//func wordBreak(s string, wordDict []string) bool {
//    lenS := len(s)
//    dp := make([]bool, lenS+1)
//    dp[0] = true
//    loop:
//    for i := 0; i < lenS; i ++ {
//        for j := 0; j <= i; j ++ {
//            if dp[j] {
//                for _, word := range wordDict {
//                    if (i < lenS-1 && word == s[j: i+1]) || word == s[j: ] {
//                        dp[i+1] = true
//                        continue loop
//                    }
//                }
//            }
//        }
//    }
//    return dp[lenS]
//}

//给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，在字符串中增加空格来构建一个句子，使得句子中所有的单词都在词典中。返回所有这些可能的句子。
//
//说明：
//
//分隔时可以重复使用字典中的单词。
//你可以假设字典中没有重复的单词。
//示例 1：
//
//输入:
//s = "catsanddog"
//wordDict = ["cat", "cats", "and", "sand", "dog"]
//输出:
//[
//  "cats and dog",
//  "cat sand dog"
//]
//示例 2：
//
//输入:
//s = "pineapplepenapple"
//wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
//输出:
//[
//  "pine apple pen apple",
//  "pineapple pen apple",
//  "pine applepen apple"
//]
//解释: 注意你可以重复使用字典中的单词。
//示例 3：
//
//输入:
//s = "catsandog"
//wordDict = ["cats", "dog", "sand", "and", "cat"]
//输出:
//[]

//func wordBreak(s string, wordDict []string) []string {
//    lenS := len(s)
//    if lenS == 0 {
//        return []string{}
//    }
//
//    dp := make([][]bool, lenS)
//    for i := 0; i < lenS; i ++ {
//        dp[i] = make([]bool, lenS)
//    }
//
//    wordMap := make(map[string]bool)
//    wordsLen := make([]int, 0)
//    loop:
//    for _, word := range wordDict {
//        wordMap[word] = true
//        wordLen := len(word)
//        for _, wsLen := range wordsLen {
//            if wordLen == wsLen {
//                continue loop
//            }
//        }
//        wordsLen = append(wordsLen, wordLen)
//    }
//
//    for i := 0; i < lenS; i ++ {
//        if i == lenS-1 {
//            if _, ok := wordMap[s[i: ]]; ok {
//                dp[i][i] = true
//            }
//        } else {
//            for j := i; j < lenS; j ++ {
//                if _, ok := wordMap[s[i: j]]; ok {
//                    dp[i][j-1] = true
//                }
//            }
//            if _, ok := wordMap[s[i: ]]; ok {
//                dp[i][lenS-1] = true
//            }
//        }
//    }
//
//    ret := make([]string, 0)
//    dfs(dp, s, []string{}, &ret, lenS, 0, wordsLen)
//    return ret
//}
//
//func dfs(dp [][]bool, s string, strs []string, ret *[]string, lenS, idx int, wordsLen []int) {
//   //fmt.Printf("strs=%v idx=%d\n", strs, idx)
//   for _, wsLen := range wordsLen {
//       if wsLen+idx < lenS && dp[idx][idx+wsLen-1] {
//           newStrs := make([]string, len(strs))
//           copy(newStrs, strs)
//           newStrs = append(newStrs, s[idx: idx+wsLen])
//           dfs(dp, s, newStrs, ret, lenS, idx+wsLen, wordsLen)
//       } else if wsLen+idx == lenS && dp[idx][lenS-1] {
//           newStrs := make([]string, len(strs))
//           copy(newStrs, strs)
//           newStrs = append(newStrs, s[idx: ])
//           *ret = append(*ret, strings.Join(newStrs, " "))
//       }
//   }
//}

//def wordBreak(self, s, wordDict):
//    memo = {len(s): ['']}
//    def sentences(i):
//        if i not in memo:
//            memo[i] = [s[i:j] + (tail and ' ' + tail)
//            for j in range(i+1, len(s)+1)
//              if s[i:j] in wordDict
//                for tail in sentences(j)]
//        return memo[i]
//    return sentences(0)

//vector<string> wordBreak(string s, unordered_set<string>& wordDict) {
//    unordered_map<int, vector<string>> memo {{s.size(), {""}}};
//    function<vector<string>(int)> sentences = [&](int i) {
//        if (!memo.count(i))
//            for (int j=i+1; j<=s.size(); j++)
//                if (wordDict.count(s.substr(i, j-i)))
//                    for (string tail : sentences(j))
//                        memo[i].push_back(s.substr(i, j-i) + (tail=="" ? "" : ' ' + tail));
//        return memo[i];
//    };
//    return sentences(0);
//}

func wordBreak(s string, wordDict []string) []string {
    memo := make([][]string, len(s)+1)
    memo[len(s)] = []string{""}
    var f func(j int) []string
    f = func(j int) []string {
        if len(memo[j]) == 0 {
            for i := j+1; i <= len(s); i ++ {
                var st string
                if i < len(s) {
                    st = s[j: i]
                } else {
                    st = s[j: ]
                }

                for _, word := range wordDict {
                    if st == word {
                        for _, str := range f(i) {
                            if str == "" {
                                memo[j] = append(memo[j], s[j: i])
                            } else {
                                memo[j] = append(memo[j], s[j: i]+" "+str)
                            }
                        }
                    }
                }
            }
        }
        return memo[j]
    }
    r := f(0)
    return r
}