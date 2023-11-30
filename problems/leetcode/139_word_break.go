package b75

import (
	"math"
	"strings"
)

/*
139. Word Break
Medium

12006

515

Add to List

Share
Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.

Note that the same word in the dictionary may be reused multiple times in the segmentation.



Example 1:

Input: s = "leetcode", wordDict = ["leet","code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".
Example 2:

Input: s = "applepenapple", wordDict = ["apple","pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
Note that you are allowed to reuse a dictionary word.
Example 3:

Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
Output: false


Constraints:

1 <= s.length <= 300
1 <= wordDict.length <= 1000
1 <= wordDict[i].length <= 20
s and wordDict[i] consist of only lowercase English letters.
All the strings of wordDict are unique.
*/

func wordBreakv17(s string, wordDict []string) bool {
	words := make(map[string]bool, len(wordDict))
	for i := range wordDict {
		words[wordDict[i]] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && words[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

func wordBreakv16(s string, woordDict []string) bool {
	words := make(map[string]bool, len(woordDict))
	for i := range woordDict {
		words[woordDict[i]] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && words[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

func wordBreakv15(s string, wordDict []string) bool {
	words := make(map[string]bool, len(wordDict))
	for i := range wordDict {
		words[wordDict[i]] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && words[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

func wordBreakv14(s string, wordDict []string) bool {
	words := make(map[string]bool, len(wordDict))
	for i := range wordDict {
		words[wordDict[i]] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && words[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

func wordBreakv13(s string, wordDict []string) bool {
	words := make(map[string]bool, len(wordDict))
	for i := range wordDict {
		words[wordDict[i]] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && words[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

func wordBreakv12(s string, wordDict []string) bool {
	words := make(map[string]bool, len(wordDict))
	for i := range wordDict {
		words[wordDict[i]] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && words[s[j:i]] {
				dp[i] = true
			}
		}
	}

	return dp[len(s)]
}

func wordBreakv11(s string, wordDict []string) bool {
	words := make(map[string]bool, len(wordDict))
	for i := range wordDict {
		words[wordDict[i]] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && words[s[j:i]] {
				dp[i] = true
			}
		}
	}

	return dp[len(s)]
}

func wordBreakv10(s string, wordDict []string) bool {
	words := make(map[string]bool, len(wordDict))
	for i := range wordDict {
		words[wordDict[i]] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && words[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

func wordBreakv9(s string, wordDict []string) bool {
	words := make(map[string]bool, len(wordDict))
	for i := range wordDict {
		words[wordDict[i]] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && words[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

func wordBreakv8(s string, wordDict []string) bool {
	words := make(map[string]bool, len(wordDict))
	for i := range wordDict {
		words[wordDict[i]] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && words[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

func wordBreakv7(s string, wordDict []string) bool {
	words := make(map[string]bool, len(wordDict))
	for i := range wordDict {
		words[wordDict[i]] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && words[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

func wordBreakv6(s string, wordDict []string) bool {
	words := make(map[string]bool, len(wordDict))
	for i := range wordDict {
		words[wordDict[i]] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && words[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

func wordBreakv5(s string, wordDict []string) bool {
	words := make(map[string]bool)
	for i := range wordDict {
		words[wordDict[i]] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && words[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

func wordBreakv4(s string, wordDict []string) bool {
	m := make(map[string]bool, len(wordDict))
	for i := range wordDict {
		m[wordDict[i]] = true
	}

	isBreak := make([]bool, len(s)+1)
	isBreak[0] = true
	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if isBreak[j] && m[s[j:i]] {
				isBreak[i] = true
				break
			}
		}
	}

	return isBreak[len(s)]
}

func wordBreakv3(s string, wordDict []string) bool {
	wd := make(map[string]bool, len(wordDict))
	for i := range wordDict {
		wd[wordDict[i]] = true
	}

	wordsExist := make([]bool, len(s)+1)
	wordsExist[0] = true
	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if wordsExist[j] && wd[s[j:i]] {
				wordsExist[i] = true
				break
			}
		}
	}

	return wordsExist[len(s)]
}

// bottom-up
func wordBreakv2(s string, wordDict []string) bool {
	wd := make(map[string]bool, len(wordDict))
	for _, v := range wordDict {
		wd[v] = true
	}

	sLen := len(s)
	m := make([]bool, sLen+1)
	m[0] = true

	for i := 1; i <= sLen; i++ {
		for j := i - 1; j >= 0; j-- {
			if m[j] && wd[s[j:i]] {
				m[i] = true
				break
			}
		}
	}
	return m[sLen]
}

func wordBreakv1(s string, wordDict []string) bool {
	for _, wd := range wordDict {
		sLen := len(s)
		m := make([][]int, sLen+1)
		wdLen := len(wd)
		m[0] = make([]int, wdLen+1)
		q := math.MinInt
		for j := 1; j <= sLen; j++ {
			m[j] = make([]int, wdLen+1)
			for k := 1; k <= wdLen; k++ {
				if s[j-1] == wd[k-1] {
					m[j][k] = 1 + m[j-1][k]
				} else {
					m[j][k] += m[j][k-1]
				}
				q = max(q, m[j][k])
			}
		}
		if q == len(wd) {
			s = strings.Replace(s, wd, "", -1)
		} else {
			return false
		}

	}

	return true
}
