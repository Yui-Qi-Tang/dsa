package b75

import (
	"sort"
)

/*
14. Longest Common Prefix

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".


Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.

Constraints:

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lowercase English letters.
*/

func longestCommonPrefixv7(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	var result string

	minLen := len(strs[0])
	for i := 1; i < len(strs); i++ {
		minLen = min(minLen, len(strs[i]))
	}

	for i := 0; i < minLen; i++ {
		common := ""
		for j := 1; j < len(strs); j++ {
			if strs[0][i] == strs[j][i] {
				common = string(strs[0][i])
			} else {
				common = ""
				break
			}
		}

		if common == "" {
			break
		}
		result += common
	}

	return result
}

func longestCommonPrefixv6(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	sort.Strings(strs)

	first, last := strs[0], strs[len(strs)-1]
	minLen := min(len(first), len(last))
	var common string
	for i := 0; i < minLen; i++ {
		if first[i] == last[i] {
			common += string(first[i])
		} else {
			break
		}
	}

	if common == "" {
		return ""
	}
	return common
}

func longestCommonPrefixv5(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	var result string

	minLen := len(strs[0])
	for i := 1; i < len(strs); i++ {
		minLen = min(minLen, len(strs[i]))
	}

	for i := 0; i < minLen; i++ {
		common := ""
		for j := 1; j < len(strs); j++ {
			if strs[0][i] == strs[j][i] {
				common = string(strs[0][i])
			} else {
				common = ""
				break
			}
		}

		if common == "" {
			break
		}
		result += common
	}

	return result
}

func longestCommonPrefixv4(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	var result string

	minLen := len(strs[0])
	for i := 1; i < len(strs); i++ {
		minLen = min(minLen, len(strs[i]))
	}

	for i := 0; i < minLen; i++ {
		common := ""
		for j := 1; j < len(strs); j++ {
			if strs[0][i] == strs[j][i] {
				common = string(strs[0][i])
			} else {
				common = ""
				break
			}
		}
		if common == "" {
			break
		}
		result += common
	}
	return result
}

func longestCommonPrefixv3(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	var result string
	minLen := len(strs[0])
	for i := 1; i < len(strs); i++ {
		minLen = min(minLen, len(strs[i]))
	}

	for i := 0; i < minLen; i++ {
		common := ""
		for j := 1; j < len(strs); j++ {
			if strs[0][i] != strs[j][i] {
				common = ""
				break
			}
			common = string(strs[0][i])
		}

		if common == "" {
			return result
		}
		result += common
	}

	return result
}

func longestCommonPrefixv2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	var result string = ""
	minLen := len(strs[0])
	for i := 1; i < len(strs); i++ {
		minLen = min(minLen, len(strs[i]))
	}

	for i := 0; i < minLen; i++ {
		common := ""
		for j := 1; j < len(strs); j++ {
			if strs[0][i] != strs[j][i] {
				common = ""
				break
			}
			common = string(strs[j][i])
		}

		if common == "" {
			break
		}

		result += common
	}

	return result
}

func longestCommonPrefixv1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}
	var result string
	sort.Strings(strs)
	first := strs[0]
	for i := range first {
		common := ""
		for _, str := range strs[1:] {
			if i < len(str) && first[i] == str[i] {
				common = string(str[i])
			} else {
				common = ""
			}
		}

		if common == "" {
			break
		}
		result += common
	}
	return result
}
