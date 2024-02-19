package b75

/*
Given an m x n integers matrix, return the length of the longest increasing path in matrix.

From each cell, you can either move in four directions: left, right, up, or down. You may not move diagonally or move outside the boundary (i.e., wrap-around is not allowed).



Example 1:


Input: matrix = [[9,9,4],[6,6,8],[2,1,1]]
Output: 4
Explanation: The longest increasing path is [1, 2, 6, 9].
Example 2:


Input: matrix = [[3,4,5],[3,2,6],[2,2,1]]
Output: 4
Explanation: The longest increasing path is [3, 4, 5, 6]. Moving diagonally is not allowed.
Example 3:

Input: matrix = [[1]]
Output: 1


Constraints:

m == matrix.length
n == matrix[i].length
1 <= m, n <= 200
0 <= matrix[i][j] <= 231 - 1
*/

func longestIncreasingPathv1(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make(map[[2]int]int)
	var dfs func(i, j, prevVal int) int
	dfs = func(i, j, prevVal int) int {
		if i < 0 || i == m || j < 0 || j == n || matrix[i][j] <= prevVal {
			return 0
		}

		if ans, exist := dp[[2]int{i, j}]; exist {
			return ans
		}

		result := 1
		result = max(result, 1+dfs(i+1, j, matrix[i][j])) // right
		result = max(result, 1+dfs(i-1, j, matrix[i][j])) // left
		result = max(result, 1+dfs(i, j+1, matrix[i][j])) // down
		result = max(result, 1+dfs(i, j-1, matrix[i][j])) // up
		dp[[2]int{i, j}] = result
		return result
	}
	result := 1
	for i := range matrix {
		for j := range matrix[i] {
			result = max(result, dfs(i, j, -1))
		}
	}
	return result
}
