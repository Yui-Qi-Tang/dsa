package b75

import (
	"math"
)

/*
221. Maximal Square
Given an m x n binary matrix filled with 0's and 1's,
find the largest square containing only 1's and return its area.



Example 1:
Input: matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
Output: 4

Example 2:
Input: matrix = [["0","1"],["1","0"]]
Output: 1

Example 3:
Input: matrix = [["0"]]
Output: 0


Constraints:

m == matrix.length
n == matrix[i].length
1 <= m, n <= 300
matrix[i][j] is '0' or '1'.
*/

func maximalSquarev23(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	result := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))
				result = max(result, dp[i][j])
			}
		}
	}

	return result * result
}

func maximalSquarev22(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	result := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))
				result = max(result, dp[i][j])
			}
		}
	}
	return result * result
}

func maximalSquarev21(matrix [][]byte) int {
	result := 0

	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1])
				result = max(result, dp[i][j])
			}
		}
	}

	return result * result
}

func maximalSquarev20(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	result := 0

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = 1 + min(dp[i][j-1], dp[i-1][j])
				result = max(result, dp[i][j])
			}
		}
	}

	return result * result
}

func maximalSquarev19(matirx [][]byte) int {
	m, n := len(matirx), len(matirx[0])

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	var result int

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matirx[i-1][j-1] == '1' {
				dp[i][j] = 1 + min(dp[i][j-1], dp[i-1][j])
				result = max(result, dp[i][j])
			}
		}
	}

	return result * result
}

func maximalSquarev18(matrix [][]byte) int {

	m, n := len(matrix), len(matrix[0])

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	result := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1])
				result = max(result, dp[i][j])
			}
		}
	}

	return result * result
}

func maximalSquarev17(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	result := 0

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1])
				result = max(result, dp[i][j])
			}
		}
	}

	return result * result
}

func maximalSquarev16(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	result := 0

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1])
				result = max(result, dp[i][j])
			}
		}
	}

	return result * result
}

func maximalSquarev15(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	result := 0

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = 1 + min(dp[i][j-1], dp[i-1][j])
				result = max(result, dp[i][j])
			}
		}
	}

	return result * result
}

func maximalSquarev14(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])

	dp := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	result := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1])
				result = max(result, dp[i][j])
			}
		}
	}

	return result * result
}

func maximalSquarev13(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])

	dp := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	result := 0

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1])
				result = max(result, dp[i][j])
			}
		}
	}

	return result * result
}

func maximalSquarev12(matrix [][]byte) int {

	m, n := len(matrix), len(matrix[0])

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	result := 0

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {

			if matrix[i-1][j-1] == '1' {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1])

				result = max(result, dp[i][j])
			}

		}
	}

	return result * result
}

func maximalSquarev11(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	result := 0

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1])
				result = max(result, dp[i][j])
			}
		}
	}

	return result * result
}

func maximalSquarev10(matrix [][]byte) int {

	if matrix == nil {
		return 0
	}

	if matrix[0] == nil {
		return 0
	}

	m, n := len(matrix), len(matrix[0])

	dp := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	result := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1])
			}
			result = max(result, dp[i][j])
		}
	}

	return result * result

}

func maximalSquarev9(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])

	dp := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	result := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= m; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = 1 + min(dp[i][j-1], dp[i-1][j])
				result = max(result, dp[i][j])
			}
		}
	}

	return result * result
}

func maximalSquarev8(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)

	}

	result := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = 1 + min(dp[i][j-1], dp[i-1][j])
				result = max(result, dp[i][j])
			}
		}
	}

	return result * result
}

func maximalSquarev7(matrix [][]byte) int {

	m, n := len(matrix), len(matrix[0])

	dp := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	result := 0

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1])
				result = max(result, dp[i][j])
			}
		}
	}

	return result * result

}

func maximalSquarev6(matirx [][]byte) int {
	m, n := len(matirx), len(matirx[0])

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	result := 0

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matirx[i-1][j-1] == '1' {
				dp[i][j] = 1 + min(
					dp[i-1][j], dp[i][j-1],
				)
				result = max(result, dp[i][j])
			}

		}
	}

	return result * result
}

func maximalSquarev5(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	result := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1])
				result = max(result, dp[i][j])
			}
		}
	}

	return result * result
}

func maximalSquarev4(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	result := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))

				result = max(result, dp[i][j])
			}
		}
	}

	return result * result

}

func maximalSquarev3(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	result := math.MinInt
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = 1 + min(
					dp[i-1][j-1],
					min(dp[i-1][j], dp[i][j-1]),
				)
			}
			result = max(result, dp[i][j])
		}
	}

	return result * result
}

func maximalSquarev2(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	result := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = 1 + min(
					dp[i-1][j-1],
					min(dp[i-1][j], dp[i][j-1]),
				)
				result = max(result, dp[i][j])
			}
		}
	}

	return result * result
}

// top-down recursive
func maximalSquarev1(matrix [][]byte) int {

	m, n := len(matrix), len(matrix[0])

	dp := make(map[int]map[int]int, m)

	// dp[i][j] -> max square length

	var helper func(i, j int) int

	result := 0
	helper = func(i, j int) int {

		if i >= m || j >= n {
			return 0
		}

		if v, exist := dp[i][j]; exist {
			return v
		}

		dp[i] = make(map[int]int)
		down := helper(i+1, j)
		right := helper(i, j+1)
		diag := helper(i+1, j+1)

		if matrix[i][j] == '1' {
			dp[i][j] = 1 + min(down, min(right, diag))
			result = max(result, dp[i][j])
		}

		return dp[i][j]

	}

	helper(0, 0)

	return result * result
}
