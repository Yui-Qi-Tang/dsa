package b75

import (
	"fmt"
)

/*

72. Edit Distance

Given two strings word1 and word2,
return the minimum number of operations required to convert word1 to word2.

You have the following three operations permitted on a word:

Insert a character
Delete a character
Replace a character


Example 1:

Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation:
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')
Example 2:

Input: word1 = "intention", word2 = "execution"
Output: 5
Explanation:
intention -> inention (remove 't')
inention -> enention (replace 'i' with 'e')
enention -> exention (replace 'n' with 'x')
exention -> exection (replace 'n' with 'c')
exection -> execution (insert 'u')


Constraints:

0 <= word1.length, word2.length <= 500
word1 and word2 consist of lowercase English letters.

*/

func editDistancev35(word1, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j]))
			}
		}
	}

	return dp[m][n]
}

func editDistancev34(word1, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j]))
			}
		}
	}

	return dp[m][n]
}

func editDistancev33(word1, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))
			}

		}
	}
	return dp[m][n]
}

func editDistancev32(word1, word2 string) int {
	m, n := len(word1), len(word2)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))
			}
		}
	}
	return dp[m][n]
}

func editDistancev31(word1, word2 string) int {
	m, n := len(word1), len(word2)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))
			}
		}
	}

	return dp[m][n]
}

func editDistancev30(word1, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))
			}
		}
	}
	return dp[m][n]
}

func editDistancev29(word1, word2 string) int {
	m, n := len(word1), len(word2)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j]))
			}
		}
	}

	return dp[m][n]
}

func editDistancev28(word1, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))
			}
		}
	}

	return dp[m][n]
}

func editDistancev27(word1, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))
			}
		}
	}

	return dp[m][n]
}

func editDistancev26(word1, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))
			}
		}
	}

	return dp[m][n]
}

func editDistancev25(word1, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j]))
			}
		}
	}

	return dp[m][n]
}

func editDistancev24(word1, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j]))
			}
		}
	}
	return dp[m][n]
}

func editDistancev23(word1, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))
			}
		}
	}

	return dp[m][n]
}

func editDistancev22(word1, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j]))
			}
		}
	}
	return dp[m][n]
}

func editDistancev21(word1, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))
			}
		}
	}

	return dp[m][n]
}

func editDistancev20(word1, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j]))
			}
		}
	}

	return dp[m][n]
}

func editDistancev19(word1, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))
			}
		}
	}

	return dp[m][n]
}

func editDistancev18(word1, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))
			}
		}
	}
	return dp[m][n]
}

func editDistancev17(word1, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1) // +1 for ""
	for i := range dp {
		dp[i] = make([]int, n+1) // +1 for ""
		dp[i][0] = i
	}

	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] // replace
			} else {
				// at least 1 operation for edition
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))
			}
		}
	}

	return dp[m][n]
}

func editDistancev16(word1, word2 string) int {
	m, n := len(word1), len(word2)

	if m == 0 || n == 0 {
		return max(m, n)
	}

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] += 1 + min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j]))
			}
		}
	}

	return dp[m][n]
}

func editDistancev15(word1, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j]))
			}
		}
	}

	return dp[m][n]
}

func editDistancev14(word1, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 || n == 0 {
		return max(m, n)
	}

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for i := 0; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))
			}
		}
	}

	return dp[m-1][n-1]
}

func editDistancev13(word1, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 || n == 0 {
		return max(m, n)
	}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for i := 0; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j]))
			}
		}
	}

	return dp[m][n]
}

func editDistancev12(word1, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 || n == 0 {
		return max(m, n)
	}

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))
			}
		}
	}
	return dp[m][n]
}

func editDistancev11(word1, word2 string) int {
	m, n := len(word2), len(word1)
	if m == 0 || n == 0 {
		return max(m, n)
	}

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for i := 0; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word2[i-1] == word1[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min( // 1 => 1 op for string converting
					dp[i-1][j-1],
					min(dp[i][j-1], dp[i-1][j]),
				)

			}
		}
	}

	return dp[m][n]
}

func editDistancev10(word1, word2 string) int {
	m, n := len(word2), len(word1)

	if m == 0 || n == 0 {
		return max(m, n)
	}

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for i := 0; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word2[i-1] == word1[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(
					dp[i-1][j-1],
					min(dp[i][j-1], dp[i-1][j]),
				)
			}
		}
	}

	return dp[m][n]
}

func editDistancev9(word1, word2 string) int {
	m, n := len(word2), len(word1)
	if m == 0 || n == 0 {
		return max(m, n)
	}

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i // word1 -> _ of word2
	}

	for i := 0; i <= n; i++ {
		dp[0][i] = i // word1 _ -> word2
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word2[i-1] == word1[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(
					dp[i-1][j-1],
					min(dp[i-1][j], dp[i][j-1]),
				)
			}
		}
	}

	return dp[m][n]
}

func editDistancev8(word1, word2 string) int {
	m, n := len(word2), len(word1)

	if m == 0 || n == 0 {
		return max(m, n)
	}

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for i := 0; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word2[i-1] == word1[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(
					dp[i-1][j-1],
					min(dp[i][j-1], dp[i-1][j]),
				)
			}

		}
	}

	return dp[m][n]
}

func editDistancev7(word1, word2 string) int {
	m, n := len(word2), len(word1)

	if m == 0 || n == 0 {
		return max(m, n)
	}

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for i := 0; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word2[i-1] == word1[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(
					dp[i-1][j-1],
					min(dp[i][j-1], dp[i-1][j]),
				)
			}
		}
	}

	return dp[m][n]
}

func editDistancev6(word1, word2 string) int {
	m, n := len(word2), len(word1)
	if m == 0 || n == 0 {
		return max(m, n)
	}

	dp := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = 0
	}

	for i := 0; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {

		for j := 1; j <= n; j++ {
			if word2[i-1] == word1[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(
					dp[i-1][j-1],
					min(dp[i-1][j], dp[i][j-1]),
				)
			}

		}

	}

	return dp[m][n]
}

func editDistancev5(word1, word2 string) int {
	if len(word1) == 0 || len(word2) == 0 {
		return max(len(word1), len(word2))
	}

	m, n := len(word2), len(word1)

	dp := make([][]int, m+1) // word2 = [_, c1, c2,...cm]
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1) // word1 = [_, c1, c2,...cn]
		dp[i][0] = i
	}

	for i := 0; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ { // compute from word2 = [c1,c2,...cm]
		for j := 1; j <= n; j++ { // compute from word1 = [c1,c2,cn]

			if word2[i-1] == word1[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 1: delete/update/replace + the min operations from the previus one
				dp[i][j] = 1 + min(
					dp[i-1][j-1],
					min(dp[i-1][j], dp[i][j-1]),
				)
			}

		}
	}

	return dp[m][n]
}

func editDistancev4(word1, word2 string) int {

	m, n := len(word2), len(word1)
	if len(word2) == 0 || len(word1) == 0 {
		return max(m, n)
	}

	dp := make([][]int, m+1)

	// if word1 = _ -> insert each char from word2
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	// if word2 = _ -> delete each char from word1
	for i := 0; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// the min op from insert/delete/replace
			if word2[i-1] == word1[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(
					dp[i-1][j-1],
					min(dp[i][j-1], dp[i-1][j]),
				)
			}
		}
	}

	for _, d := range dp {
		fmt.Println(d)
	}

	return dp[m][n]
}

func editDistancev3(word1, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 || n == 0 {
		return max(m, n)
	}

	dp := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1) //  [_, a, b, c...]
	}

	// if word2 is _ -> delete each char of word1
	for i := 1; i <= m; i++ {
		dp[0][i] = i
	}

	// if word1 is _ -> insert each char into word2
	for i := 1; i <= n; i++ {
		dp[i][0] = i
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word2[i-1] == word1[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(
					dp[i-1][j-1],
					min(dp[i-1][j], dp[i][j-1]),
				)
			}
		}
	}

	return dp[n][m]
}

func editDistancev2(word1, word2 string) int {
	w1l, w2l := len(word1), len(word2)

	if w1l == 0 {
		return w2l
	}
	if w2l == 0 {
		return w1l
	}
	m := make([][]int, w1l+1)

	for i := 0; i <= w1l; i++ {
		m[i] = make([]int, w2l+1)

		m[i][w2l] = w1l - i
	}

	for i := 0; i <= w2l; i++ {
		m[w1l][i] = w2l - i
	}

	for i := w1l - 1; i >= 0; i-- {
		for j := w2l - 1; j >= 0; j-- {
			if word1[i] == word2[j] {
				m[i][j] = m[i+1][j+1]
			} else {
				m[i][j] = 1 + min(m[i+1][j+1], min(m[i][j+1], m[i+1][j]))
			}
		}
	}

	return m[0][0]
}

func editDistancev1(word1, word2 string) int {
	w1l := len(word1)
	w2l := len(word2)

	m := make([][]int, w2l+1)

	for i := 0; i <= w2l; i++ {
		m[i] = make([]int, w1l+1)
		m[i][w1l] = w2l - i
	}

	for i := 0; i < w1l; i++ {
		m[w2l][i] = w1l - i
	}

	for i := w2l - 1; i >= 0; i-- {
		for j := w1l - 1; j >= 0; j-- {
			if word2[i] == word1[j] {
				m[i][j] = m[i+1][j+1]
			} else {
				m[i][j] = 1 + min(m[i+1][j+1], min(m[i][j+1], m[i+1][j]))
			}

		}
	}

	return m[0][0]
}

func EditDistance(word1, word2 string) int {
	w1l := len(word1) // w1l: length of word1
	w2l := len(word2) // w2l: length of word2

	memo := make([][]int, w1l+1) // w1 + empty

	for i := range memo {
		memo[i] = make([]int, w2l+1)

		memo[i][w2l] = w1l - i
	}

	for i := range memo[w1l] {
		memo[w1l][i] = w2l - i
	}

	//	memo[w1l][w2l] = 0

	for i := w1l - 1; i >= 0; i-- {
		for j := w2l - 1; j >= 0; j-- {

			if word1[i] == word2[j] {
				memo[i][j] = memo[i+1][j+1] // memo[i+1][j+1] the same
			} else {
				memo[i][j] = 1 + min(
					memo[i+1][j+1],
					min(memo[i][j+1], memo[i+1][j]),
				)
			}

		}
	}

	return memo[0][0]
}
