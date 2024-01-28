package b75

import (
	"fmt"
)

/*
64. Minimum Path Sum
Medium
Given a m x n grid filled with non-negative numbers,
find a path from top left to bottom right,
which minimizes the sum of all numbers along its path.

Note: You can only move either down or right at any point in time.



Example 1:


1,3,1
1,5,1
4,2,1

Input: grid = [[1,3,1],[1,5,1],[4,2,1]]
Output: 7
Explanation: Because the path 1 → 3 → 1 → 1 → 1 minimizes the sum.
Example 2:

Input: grid = [[1,2,3],[4,5,6]]
Output: 12


Constraints:

m == grid.length
n == grid[i].length
1 <= m, n <= 200
0 <= grid[i][j] <= 100
*/

func minPathSumv26(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
			}
		}
	}

	return grid[m-1][n-1]
}

func minPathSumv25(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += min(grid[i][j-1], grid[i-1][j])
			}
		}
	}

	return grid[m-1][n-1]
}

func minPathSumv24(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
			}

		}
	}
	return grid[m-1][n-1]
}

func minPathSumv23(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += min(grid[i][j-1], grid[i-1][j])
			}
		}
	}

	return grid[m-1][n-1]
}

func minPathSumv22(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
			}
		}
	}
	return grid[m-1][n-1]
}

func minPathSumv21(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += min(grid[i][j-1], grid[i-1][j])
			}
		}
	}

	return grid[m-1][n-1]
}

func minPathSumv20(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += min(grid[i][j-1], grid[i-1][j])
			}
		}
	}

	return grid[m-1][n-1]
}

func minPathSumv19(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += min(grid[i][j-1], grid[i-1][j])
			}
		}
	}
	return grid[m-1][n-1]
}

func minPathSumv18(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
			}
		}
	}

	return grid[m-1][n-1]
}

func minPathSumv17(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
			}
		}
	}

	return grid[m-1][n-1]
}

func minPathSumv16(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
			}
		}
	}

	return grid[m-1][n-1]
}

func minPathSumv15(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += min(grid[i][j-1], grid[i-1][j])
			}
		}
	}

	return grid[m-1][n-1]
}

func minPathSumv14(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += min(grid[i][j-1], grid[i-1][j])
			}
		}
	}

	return grid[m-1][n-1]
}

func minPathSumv13(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 {
				dp[i][j] = grid[i][j] + dp[i][j-1]
			} else if j == 0 {
				dp[i][j] = grid[i][j] + dp[i-1][j]
			} else {
				dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m-1][n-1]
}

func minPathSumv12(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 {
				dp[i][j] += grid[i][j] + dp[i][j-1]
			} else if j == 0 {
				dp[i][j] += grid[i][j] + dp[i-1][j]
			} else {
				dp[i][j] += min(grid[i][j]+dp[i-1][j], grid[i][j]+dp[i][j-1])
			}
		}
	}

	return dp[m-1][n-1]
}

func minPathSumv11(grid [][]int) int {
	if grid == nil {
		return 0
	}

	if grid[0] == nil {
		return 0
	}

	m, n := len(grid), len(grid[0])

	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if j == 0 && i == 0 {
				dp[i][j] = grid[i][j]
			} else if j == 0 {
				dp[i][j] = grid[i][j] + dp[i-1][j]
			} else if i == 0 {
				dp[i][j] = grid[i][j] + dp[i][j-1]
			} else {
				dp[i][j] = min(grid[i][j]+dp[i][j-1], grid[i][j]+dp[i-1][j])
			}
		}
	}

	return dp[m-1][n-1]
}

func minPathSumv10(grid [][]int) int {
	if grid == nil {
		return 0
	}

	if grid[0] == nil {
		return 0
	}

	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {

			if j == 0 && i == 0 {
				dp[i][j] = grid[i][j]
			} else if j == 0 {
				dp[i][j] = grid[i][j] + dp[i-1][j]
			} else if i == 0 {
				dp[i][j] = grid[i][j] + dp[i][j-1]
			} else {
				dp[i][j] = min(grid[i][j]+dp[i-1][j], grid[i][j]+dp[i][j-1])
			}

		}
	}
	return dp[m-1][n-1]
}

func minPathSumv9(grid [][]int) int {
	if grid == nil {
		return 0
	}

	if grid[0] == nil {
		return 0
	}

	m, n := len(grid), len(grid[0])

	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {

			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 {
				dp[i][j] = grid[i][j] + dp[i][j-1]
			} else if j == 0 {
				dp[i][j] = grid[i][j] + dp[i-1][j]
			} else {
				dp[i][j] = min(grid[i][j]+dp[i][j-1], grid[i][j]+dp[i-1][j])
			}

		}
	}

	return dp[m-1][n-1]
}

func minPathSumv8(grid [][]int) int {
	if grid == nil {
		return 0
	}

	if grid[0] == nil {
		return 0
	}

	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if j == 0 && i == 0 {
				dp[i][j] = grid[i][j]
			} else if j == 0 {
				dp[i][j] = grid[i][j] + dp[i-1][j]
			} else if i == 0 {
				dp[i][j] = grid[i][j] + dp[i][j-1]
			} else {
				dp[i][j] = min(grid[i][j]+dp[i-1][j], grid[i][j]+dp[i][j-1])
			}
		}

	}

	return dp[m-1][n-1]
}

func minPathSumv7(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}

	n := len(grid[0])
	if n == 0 {
		return 0
	}

	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)

		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 {
				dp[i][j] = grid[i][j] + dp[i][j-1]
			} else if j == 0 {
				dp[i][j] = grid[i][j] + dp[i-1][j]
			} else {
				dp[i][j] = min(grid[i][j]+dp[i][j-1], grid[i][j]+dp[i-1][j])
			}
		}

	}

	fmt.Println(dp)

	return dp[m-1][n-1]
}

func minPathSumv6(grid [][]int) int {

	m, n := len(grid), len(grid[0])

	if m == 0 {
		return 0
	}

	if n == 0 {
		return 0
	}

	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[0][0]
			} else if i == 0 {
				dp[i][j] = grid[i][j] + dp[i][j-1]
			} else if j == 0 {
				dp[i][j] = grid[i][j] + dp[i-1][j]
			} else {
				dp[i][j] = min(grid[i][j]+dp[i][j-1], grid[i][j]+dp[i-1][j])
			}

		}
	}

	return dp[m-1][n-1]

}

func minPathSumv5(grid [][]int) int {
	if grid == nil {
		return 0
	}

	if grid[0] == nil {
		return 0
	}

	m, n := len(grid), len(grid[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {

			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 && j > 0 {
				dp[i][j] = grid[i][j] + dp[i][j-1]
			} else if j == 0 && i > 0 {
				dp[i][j] = grid[i][j] + dp[i-1][j]
			} else {
				dp[i][j] += min(grid[i][j]+dp[i-1][j], grid[i][j]+dp[i][j-1])
			}

		}
	}

	return dp[m-1][n-1]
}

// bottom-up
func minPathSumv4(grid [][]int) int {

	m := len(grid)
	if m == 0 {
		return 0
	}

	n := len(grid[0])

	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i-1 < 0 && j-1 < 0 { // stand alone
				dp[i][j] += grid[i][j]
			} else if i-1 < 0 { // only from left
				dp[i][j] += grid[i][j] + dp[i][j-1]
			} else if j-1 < 0 { // only from top
				dp[i][j] += grid[i][j] + dp[i-1][j]
			} else { // left or top
				dp[i][j] += grid[i][j] + min(dp[i-1][j], dp[i][j-1])
			}

		}
	}

	return dp[m-1][n-1]

}

func minPathSumv3(grid [][]int) int {

	m, n := len(grid), len(grid[0])

	dp := make([][]int, m+1)

	for i := 1; i <= m; i++ {
		dp[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {

			if i-1 == 0 && j-1 == 0 {
				dp[i][j] = grid[i-1][j-1]
			} else if i-1 == 0 {
				dp[i][j] = grid[i-1][j-1] + dp[i][j-1]
			} else if j-1 == 0 {
				dp[i][j] = grid[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = min(grid[i-1][j-1]+dp[i-1][j], grid[i-1][j-1]+dp[i][j-1])
			}
		}
	}

	return dp[m][n]

}

func minPathSumv2(grid [][]int) int {

	m, n := len(grid), len(grid[0])

	dp := make(map[int]map[int]int) // i, j -> min

	var backtrack func(i, j int) int
	backtrack = func(i, j int) int {

		if v, exist := dp[i][j]; exist {
			return v
		}

		dp[i] = make(map[int]int)
		if i == m-1 && j == n-1 { // bottom
			dp[i][j] = grid[i][j]
		} else if i == m-1 { // only right
			dp[i][j] = grid[i][j] + backtrack(i, j+1)
		} else if j == n-1 { // only down
			dp[i][j] = grid[i][j] + backtrack(i+1, j)
		} else { // right or down
			dp[i][j] = min(grid[i][j]+backtrack(i, j+1), grid[i][j]+backtrack(i+1, j))
		}
		return dp[i][j]

	}

	return backtrack(0, 0)

}

func minPathSumv1(grid [][]int) int {

	dp := make(map[int]map[int]int) // i, j -> min

	var backtrack func(i, j int) int
	backtrack = func(i, j int) int {

		if v, exist := dp[i][j]; exist {
			return v
		}
		dp[i] = make(map[int]int)

		if i == len(grid)-1 && j == len(grid[0])-1 {
			dp[i][j] = grid[i][j]
			return dp[i][j]
		}

		if i == len(grid)-1 {
			dp[i][j] = grid[i][j] + backtrack(i, j+1)
			return dp[i][j]
		}

		if j == len(grid[0])-1 {
			dp[i][j] = grid[i][j] + backtrack(i+1, j)
			return dp[i][j]
		}

		dp[i][j] = min(grid[i][j]+backtrack(i, j+1), grid[i][j]+backtrack(i+1, j))

		return dp[i][j]

	}

	return backtrack(0, 0)

}
