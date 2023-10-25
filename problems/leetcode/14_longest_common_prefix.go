package b75

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

aaa
bbb
ccc
ddd
aab


Constraints:

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lowercase English letters.
*/

func longestCommonPrefixv2(strs []string) string {
	var result string

	for i := 0; i < len(strs)-1; i++ {
		j := i + 1
		commonIdx := 0
		minLen := min(len(strs[i]), len(strs[j]))
		for k := 0; k < minLen; k++ {
			if strs[i][k] == strs[j][k] {
				commonIdx++
			}
		}

		if commonIdx > 0 && strs[i][:commonIdx] == strs[j][:commonIdx] {
			result = strs[i][:commonIdx]
		} else {
			return ""
		}
	}

	return result
}

func longestCommonPrefixv1(strs []string) string {

	var result string
	commonIdx := 0
	for i := 0; i < len(strs); i++ {
		j := i + 1

		if j < len(strs) {
			mLen := min(len(strs[i]), len(strs[j]))

			for k := 0; k < mLen; k++ {
				if strs[i][k] == strs[j][k] {
					commonIdx = k
				}
			}

			if commonIdx > 0 && strs[i][:commonIdx] == strs[j][:commonIdx] {
				result = strs[i][:commonIdx+1]
			} else {
				result = ""
			}
		}
	}

	return result
}
