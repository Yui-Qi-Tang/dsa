package b75

/*
97. Interleaving String
Medium
Topics
Companies
Given strings s1, s2, and s3, find whether s3 is formed by an interleaving of s1 and s2.

An interleaving of two strings s and t is a configuration where s and t are divided into n and m
substrings
 respectively, such that:

s = s1 + s2 + ... + sn
t = t1 + t2 + ... + tm
|n - m| <= 1
The interleaving is s1 + t1 + s2 + t2 + s3 + t3 + ... or t1 + s1 + t2 + s2 + t3 + s3 + ...
Note: a + b is the concatenation of strings a and b.



Example 1:


Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
Output: true
Explanation: One way to obtain s3 is:
Split s1 into s1 = "aa" + "bc" + "c", and s2 into s2 = "dbbc" + "a".
Interleaving the two splits, we get "aa" + "dbbc" + "bc" + "a" + "c" = "aadbbcbcac".
Since s3 can be obtained by interleaving s1 and s2, we return true.
Example 2:

Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
Output: false
Explanation: Notice how it is impossible to interleave s2 with any other string to obtain s3.
Example 3:

Input: s1 = "", s2 = "", s3 = ""
Output: true


Constraints:

0 <= s1.length, s2.length <= 100
0 <= s3.length <= 200
s1, s2, and s3 consist of lowercase English letters.


Follow up: Could you solve it using only O(s2.length) additional memory space?
*/

func isInterleavev11(s1, s2, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := make(map[[2]int]bool)
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if i == len(s1) && j == len(s2) {
			return true
		}

		if ans, exist := dp[[2]int{i, j}]; exist {
			return ans
		}

		if i < len(s1) && s1[i] == s3[i+j] && dfs(i+1, j) {
			return true
		}

		if j < len(s2) && s2[j] == s3[i+j] && dfs(i, j+1) {
			return true
		}

		dp[[2]int{i, j}] = false
		return dp[[2]int{i, j}]
	}
	return dfs(0, 0)
}

func isInterleavev10(s1, s2, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := make(map[[2]int]bool)
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if i == len(s1) && j == len(s2) {
			return true
		}

		if ans, exist := dp[[2]int{i, j}]; exist {
			return ans
		}

		if i < len(s1) && s1[i] == s3[i+j] && dfs(i+1, j) {
			return true
		}

		if j < len(s2) && s2[j] == s3[i+j] && dfs(i, j+1) {
			return true
		}

		dp[[2]int{i, j}] = false
		return dp[[2]int{i, j}]
	}
	return dfs(0, 0)
}

func isInterleavev9(s1, s2, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := make(map[[2]int]bool)
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if i == len(s1) && j == len(s2) {
			return true
		}

		if ans, exist := dp[[2]int{i, j}]; exist {
			return ans
		}

		if i < len(s1) && s1[i] == s3[i+j] && dfs(i+1, j) {
			return true
		}

		if j < len(s2) && s2[j] == s3[i+j] && dfs(i, j+1) {
			return true
		}
		dp[[2]int{i, j}] = false
		return dp[[2]int{i, j}]
	}
	return dfs(0, 0)
}

func isInterleavev8(s1, s2, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := make(map[[2]int]bool)
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if i == len(s1) && j == len(s2) {
			return true
		}

		if ans, exist := dp[[2]int{i, j}]; exist {
			return ans
		}

		if i < len(s1) && s1[i] == s3[i+j] && dfs(i+1, j) {
			return true
		}

		if j < len(s2) && s2[j] == s3[i+j] && dfs(i, j+1) {
			return true
		}
		dp[[2]int{i, j}] = false
		return dp[[2]int{i, j}]
	}
	return dfs(0, 0)
}

func isInterleavev7(s1, s2, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := make(map[[2]int]bool)
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if i == len(s1) && j == len(s2) {
			return true
		}

		if ans, exist := dp[[2]int{i, j}]; exist {
			return ans
		}

		if i < len(s1) && s1[i] == s3[i+j] && dfs(i+1, j) {
			return true
		}

		if j < len(s2) && s2[j] == s3[i+j] && dfs(i, j+1) {
			return true
		}

		dp[[2]int{i, j}] = false
		return dp[[2]int{i, j}]
	}
	return dfs(0, 0)
}

func isInterleavev6(s1, s2, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	dp := make(map[[2]int]bool)
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if i == len(s1) && j == len(s2) {
			return true
		}

		if ans, exist := dp[[2]int{i, j}]; exist {
			return ans
		}

		if i < len(s1) && s1[i] == s3[i+j] && dfs(i+1, j) {
			return true
		}

		if j < len(s2) && s2[j] == s3[i+j] && dfs(i, j+1) {
			return true
		}
		dp[[2]int{i, j}] = false
		return false
	}
	return dfs(0, 0)
}

func isInterleavev5(s1, s2, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	dp := make(map[[2]int]bool)
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		// base case:
		// i == len(s1) -> empty
		// j == len(s2) -> empty
		// empty == empty -> true
		if i == len(s1) && j == len(s2) {
			return true
		}

		if ans, exist := dp[[2]int{i, j}]; exist {
			return ans
		}

		if i < len(s1) && s1[i] == s3[i+j] && dfs(i+1, j) {
			return true
		}

		if j < len(s2) && s2[j] == s3[i+j] && dfs(i, j+1) {
			return true
		}
		dp[[2]int{i, j}] = false
		return false
	}
	return dfs(0, 0)
}

func isInterleavev4(s1, s2, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	dp := make(map[[2]int]bool)
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if i == len(s1) && j == len(s2) {
			return true
		}

		if ans, exist := dp[[2]int{i, j}]; exist {
			return ans
		}

		if i < len(s1) && s1[i] == s3[i+j] && dfs(i+1, j) {
			return true
		}

		if j < len(s2) && s2[j] == s3[i+j] && dfs(i, j+1) {
			return true
		}

		dp[[2]int{i, j}] = false
		return false
	}
	return dfs(0, 0)
}

func isInterleavev3(s1, s2, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[m][n] = true

	for i := m; i >= 0; i-- {
		for j := n; j >= 0; j-- {
			if i < m && s1[i] == s3[i+j] && dp[i+1][j] {
				dp[i][j] = true
			}

			if j < n && s2[j] == s3[i+j] && dp[i][j+1] {
				dp[i][j] = true
			}
		}
	}

	return dp[0][0]
}

func isInterleavev2(s1, s2, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	dp := make(map[[2]int]bool)
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if i == len(s1) && j == len(s2) {
			return true
		}

		if ans, exist := dp[[2]int{i, j}]; exist {
			return ans
		}

		if i < len(s1) && s1[i] == s3[i+j] && dfs(i+1, j) {
			return true
		}

		if j < len(s2) && s2[j] == s3[i+j] && dfs(i, j+1) {
			return true
		}
		dp[[2]int{i, j}] = false
		return false
	}
	return dfs(0, 0)
}

func isInterleavev1(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}

	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[m][n] = true

	for i := m; i >= 0; i-- {
		for j := n; j >= 0; j-- {
			if i < len(s1) && s1[i] == s3[i+j] && dp[i+1][j] {
				dp[i][j] = true
			}

			if j < len(s2) && s2[j] == s3[i+j] && dp[i][j+1] {
				dp[i][j] = true
			}
		}
	}

	return dp[0][0]
}
