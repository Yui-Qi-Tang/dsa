package b75

/*
36. Valid Sudoku
Medium
7.9K
874
Companies
Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
Note:

A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.

Example 1:

Input: board =
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: true
Example 2:

Input: board =
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: false
Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.

Constraints:

board.length == 9
board[i].length == 9
board[i][j] is a digit 1-9 or '.'.
*/

func isValidSudokuv30(board [][]byte) bool {
	n := len(board)
	rows := make([]map[byte]bool, n)
	cols := make([]map[byte]bool, n)
	squares := make([]map[byte]bool, n)
	for i := range rows {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		squares[i] = make(map[byte]bool)
	}

	for i := range rows {
		for j := range cols {
			if board[i][j] == '.' {
				continue
			}

			sn := 3*(i/3) + (j / 3)

			if rows[i][board[i][j]] || cols[j][board[i][j]] || squares[sn][board[i][j]] {
				return false
			}

			rows[i][board[i][j]] = true
			cols[j][board[i][j]] = true
			squares[sn][board[i][j]] = true
		}
	}

	return true
}

func isValidSudokuv29(board [][]byte) bool {
	n := len(board)
	rows := make([]map[byte]bool, n)
	cols := make([]map[byte]bool, n)
	squares := make([]map[byte]bool, n)

	for i := range rows {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		squares[i] = make(map[byte]bool)
	}

	for i := range rows {
		for j := range cols {
			if board[i][j] == '.' {
				continue
			}
			sn := 3*(i/3) + (j / 3)
			if rows[i][board[i][j]] || cols[j][board[i][j]] || squares[sn][board[i][j]] {
				return false
			}

			rows[i][board[i][j]] = true
			cols[j][board[i][j]] = true
			squares[sn][board[i][j]] = true
		}
	}

	return true
}

func isValidSudokuv28(board [][]byte) bool {
	n := len(board)
	rows := make([]map[byte]bool, n)
	cols := make([]map[byte]bool, n)
	square := make([]map[byte]bool, n)

	for i := range rows {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		square[i] = make(map[byte]bool)
	}

	for i := range rows {
		for j := range cols {
			if board[i][j] == '.' {
				continue
			}

			sn := 3*(i/3) + (j / 3)
			if rows[i][board[i][j]] || cols[j][board[i][j]] || square[sn][board[i][j]] {
				return false
			}

			rows[i][board[i][j]] = true
			cols[j][board[i][j]] = true
			square[sn][board[i][j]] = true
		}
	}

	return true
}

func isValidSudokuv27(board [][]byte) bool {
	n := len(board)

	rows := make([]map[byte]bool, n)
	cols := make([]map[byte]bool, n)
	squares := make([]map[byte]bool, n)

	for i := range rows {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		squares[i] = make(map[byte]bool)
	}

	for i := range rows {
		for j := range cols {
			if board[i][j] == '.' {
				continue
			}

			sn := 3*(i/3) + (j / 3)
			if rows[i][board[i][j]] || cols[j][board[i][j]] || squares[sn][board[i][j]] {
				return false
			}

			rows[i][board[i][j]] = true
			cols[j][board[i][j]] = true
			squares[sn][board[i][j]] = true
		}
	}

	return true
}

func isValidSudokuv26(board [][]byte) bool {
	n := len(board)
	rows := make([]map[byte]bool, n)
	cols := make([]map[byte]bool, n)
	squares := make([]map[byte]bool, n)
	for i := range rows {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		squares[i] = make(map[byte]bool)
	}

	for i := range rows {
		for j := range cols {
			if board[i][j] == '.' {
				continue
			}

			sn := 3*(i/3) + (j / 3)
			if rows[i][board[i][j]] || cols[j][board[i][j]] || squares[sn][board[i][j]] {
				return false
			}

			rows[i][board[i][j]] = true
			cols[j][board[i][j]] = true
			squares[sn][board[i][j]] = true
		}
	}

	return true
}

func isValidSudokuv25(board [][]byte) bool {
	n := len(board)

	rows := make([]map[byte]bool, n)
	cols := make([]map[byte]bool, n)
	squares := make([]map[byte]bool, n)

	for i := range rows {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		squares[i] = map[byte]bool{}
	}

	for i := range rows {
		for j := range cols {
			if board[i][j] == '.' {
				continue
			}

			sn := 3*(i/3) + (j / 3)
			if rows[i][board[i][j]] || cols[j][board[i][j]] || squares[sn][board[i][j]] {
				return false
			}

			rows[i][board[i][j]] = true
			cols[j][board[i][j]] = true
			squares[sn][board[i][j]] = true
		}
	}

	return true
}

func isValidSudokuv24(board [][]byte) bool {
	n := len(board)

	rows := make([]map[byte]bool, n)
	cols := make([]map[byte]bool, n)
	squares := make([]map[byte]bool, n)

	for i := range rows {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		squares[i] = make(map[byte]bool)
	}

	for i := range rows {
		for j := range cols {
			if board[i][j] == '.' {
				continue
			}

			sn := 3*(i/3) + (j / 3)
			if rows[i][board[i][j]] || cols[j][board[i][j]] || squares[sn][board[i][j]] {
				return false
			}

			rows[i][board[i][j]] = true
			cols[j][board[i][j]] = true
			squares[sn][board[i][j]] = true
		}
	}

	return true
}

func isValidSudokuv23(board [][]byte) bool {
	n := len(board)
	rows := make([]map[byte]bool, n)
	cols := make([]map[byte]bool, n)
	squares := make([]map[byte]bool, n)

	for i := range rows {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		squares[i] = make(map[byte]bool)
	}

	for i := range rows {
		for j := range rows {
			if board[i][j] == '.' {
				continue
			}

			sn := 3*(i/3) + (j / 3)
			if rows[i][board[i][j]] || cols[j][board[i][j]] || squares[sn][board[i][j]] {
				return false
			}

			rows[i][board[i][j]] = true
			cols[j][board[i][j]] = true
			squares[sn][board[i][j]] = true
		}
	}

	return true
}

func isValidSudokuv22(board [][]byte) bool {
	n := len(board)
	rows := make([]map[byte]bool, n)
	cols := make([]map[byte]bool, n)
	squares := make([]map[byte]bool, n)

	for i := range rows {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		squares[i] = make(map[byte]bool)
	}

	for i := range rows {
		for j := range rows {
			if board[i][j] == '.' {
				continue
			}

			sn := 3*(i/3) + (j / 3)
			if rows[i][board[i][j]] || cols[j][board[i][j]] || squares[sn][board[i][j]] {
				return false
			}

			rows[i][board[i][j]] = true
			cols[j][board[i][j]] = true
			squares[sn][board[i][j]] = true
		}
	}

	return true
}

func isValidSudokuv21(board [][]byte) bool {
	n := len(board)
	rows := make([]map[byte]bool, n)
	cols := make([]map[byte]bool, n)
	squares := make([]map[byte]bool, n)

	for i := range rows {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		squares[i] = make(map[byte]bool)
	}

	for i := range rows {
		for j := range rows {
			if board[i][j] == '.' {
				continue
			}

			sn := 3*(i/3) + (j / 3)
			if rows[i][board[i][j]] || cols[j][board[i][j]] || squares[sn][board[i][j]] {
				return false
			}

			rows[i][board[i][j]] = true
			cols[j][board[i][j]] = true
			squares[sn][board[i][j]] = true
		}
	}

	return true
}

func isValidSudokuv20(board [][]byte) bool {
	n := len(board)
	rows := make([]map[byte]bool, n)
	cols := make([]map[byte]bool, n)
	squares := make([]map[byte]bool, n)

	for i := range rows {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		squares[i] = make(map[byte]bool)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '.' {
				continue
			}

			sn := 3*(i/3) + (j / 3)
			if rows[i][board[i][j]] || cols[j][board[i][j]] || squares[sn][board[i][j]] {
				return false
			}

			rows[i][board[i][j]] = true
			cols[j][board[i][j]] = true
			squares[sn][board[i][j]] = true
		}
	}

	return true
}

func isValidSudokuv19(board [][]byte) bool {
	n := len(board)
	if n != 9 {
		return false
	}

	rows := make([]map[byte]bool, n)
	cols := make([]map[byte]bool, n)
	squares := make([]map[byte]bool, n)

	for i := range rows {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		squares[i] = make(map[byte]bool)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '.' {
				continue
			}

			sn := 3*(i/3) + (j / 3)
			if rows[i][board[i][j]] || cols[j][board[i][j]] || squares[sn][board[i][j]] {
				return false
			}

			rows[i][board[i][j]] = true
			cols[j][board[i][j]] = true
			squares[sn][board[i][j]] = true
		}
	}

	return true
}

func isValidSudokuv18(board [][]byte) bool {
	n := len(board)
	if n < 9 {
		return false
	}

	rows := make([]map[byte]bool, n)
	cols := make([]map[byte]bool, n)
	squares := make([]map[byte]bool, n)

	for i := 0; i < n; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		squares[i] = make(map[byte]bool)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '.' {
				continue
			}

			sn := 3*(i/3) + (j / 3)
			if rows[i][board[i][j]] || cols[j][board[i][j]] || squares[sn][board[i][j]] {
				return false
			}

			rows[i][board[i][j]] = true
			cols[j][board[i][j]] = true
			squares[sn][board[i][j]] = true
		}
	}

	return true
}

func isValidSudokuv17(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	squares := make([]map[byte]bool, 9)

	for i := range rows {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		squares[i] = make(map[byte]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			sn := 3*(i/3) + (j / 3)
			if rows[i][board[i][j]] || cols[j][board[i][j]] || squares[sn][board[i][j]] {
				return false
			}

			rows[i][board[i][j]] = true
			cols[j][board[i][j]] = true
			squares[sn][board[i][j]] = true
		}
	}

	return true
}

func isValidSudokuv16(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	squares := make([]map[byte]bool, 9)

	for i := range rows {
		rows[i] = make(map[byte]bool, 9)
		cols[i] = make(map[byte]bool, 9)
		squares[i] = make(map[byte]bool, 9)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			sn := 3*(i/3) + (j / 3)

			if rows[i][board[i][j]] || cols[j][board[i][j]] || squares[sn][board[i][j]] {
				return false
			}

			rows[i][board[i][j]] = true
			cols[j][board[i][j]] = true
			squares[sn][board[i][j]] = true
		}
	}

	return true
}

func isValidSudokuv15(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	squares := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		squares[i] = make(map[byte]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			sn := 3*(i/3) + (j / 3)
			if rows[i][board[i][j]] || cols[j][board[i][j]] || squares[sn][board[i][j]] {
				return false
			}

			rows[i][board[i][j]] = true
			cols[j][board[i][j]] = true
			squares[sn][board[i][j]] = true
		}
	}

	return true
}

func isValidSudokuv14(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	squares := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool, 9)
		cols[i] = make(map[byte]bool, 9)
		squares[i] = make(map[byte]bool, 9)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			sn := 3*(i/3) + (j / 3)
			if rows[i][board[i][j]] || cols[j][board[i][j]] || squares[sn][board[i][j]] {
				return false
			}

			rows[i][board[i][j]] = true
			cols[j][board[i][j]] = true
			squares[sn][board[i][j]] = true
		}
	}

	return true
}

func isValidSudokuv13(board [][]byte) bool {
	rows := make([]map[byte]bool, len(board))
	cols := make([]map[byte]bool, len(board))
	squares := make([]map[byte]bool, len(board))

	for i := 0; i < len(board); i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		squares[i] = make(map[byte]bool)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] == '.' {
				continue
			}

			sn := 3*(i/3) + (j / 3)
			if rows[i][board[i][j]] || cols[j][board[i][j]] || squares[sn][board[i][j]] {
				return false
			}

			rows[i][board[i][j]] = true
			cols[j][board[i][j]] = true
			squares[sn][board[i][j]] = true
		}
	}

	return true
}

func isValidSudokuv12(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	squares := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool, 9)
		cols[i] = make(map[byte]bool, 9)
		squares[i] = make(map[byte]bool, 9)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			sn := 3*(i/3) + (j / 3)
			if rows[i][board[i][j]] || cols[j][board[i][j]] || squares[sn][board[i][j]] {
				return false
			}

			rows[i][board[i][j]] = true
			cols[j][board[i][j]] = true
			squares[sn][board[i][j]] = true
		}
	}

	return true
}

func isValidSudokuv11(board [][]byte) bool {
	rows := make([]map[byte]bool, len(board))
	cols := make([]map[byte]bool, len(board))
	square := make([]map[byte]bool, len(board))

	for i := 0; i < len(board); i++ {
		rows[i] = make(map[byte]bool, 9)
		cols[i] = make(map[byte]bool, 9)
		square[i] = make(map[byte]bool, 9)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			sn := 3*(i/3) + (j / 3)

			if rows[i][board[i][j]] || cols[j][board[i][j]] || square[sn][board[i][j]] {
				return false
			}
			rows[i][board[i][j]] = true
			cols[j][board[i][j]] = true
			square[sn][board[i][j]] = true
		}
	}

	return true
}

func isValidSudokuv10(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	squares := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool, 9)
		cols[i] = make(map[byte]bool, 9)
		squares[i] = make(map[byte]bool, 9)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			sn := 3*(i/3) + (j / 3)
			if rows[i][board[i][j]] || cols[j][board[i][j]] || squares[sn][board[i][j]] {
				return false
			}
			rows[i][board[i][j]] = true
			cols[j][board[i][j]] = true
			squares[sn][board[i][j]] = true
		}
	}

	return true
}

func isValidSudokuv9(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	squares := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		squares[i] = make(map[byte]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			sn := 3*(i/3) + (j / 3)
			if rows[i][board[i][j]] || cols[j][board[i][j]] || squares[sn][board[i][j]] {
				return false
			}
			rows[i][board[i][j]] = true
			cols[j][board[i][j]] = true
			squares[sn][board[i][j]] = true

		}
	}

	return true
}

func isValidSudokuv8(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	squares := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool, 9)
		cols[i] = make(map[byte]bool, 9)
		squares[i] = make(map[byte]bool, 9)

	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			sn := 3*(i/3) + (j / 3)
			if rows[i][board[i][j]] || cols[j][board[i][j]] || squares[sn][board[i][j]] {
				return false
			}

			rows[i][board[i][j]] = true
			cols[j][board[i][j]] = true
			squares[sn][board[i][j]] = true

		}
	}

	return true
}

func isValidSudokuv7(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	squares := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool, 9)
		cols[i] = make(map[byte]bool, 9)
		squares[i] = make(map[byte]bool, 9)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			if board[i][j] == '.' {
				continue
			}

			sn := 3*(i/3) + (j / 3)
			if rows[i][board[i][j]] || cols[j][board[i][j]] || squares[sn][board[i][j]] {
				return false
			}

			rows[i][board[i][j]] = true
			cols[j][board[i][j]] = true
			squares[sn][board[i][j]] = true

		}
	}

	return true
}

func isValidSudokuv6(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	squares := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool, 9)
		cols[i] = make(map[byte]bool, 9)
		squares[i] = make(map[byte]bool, 9)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			sn := 3*(i/3) + (j / 3)
			if rows[i][board[i][j]] || cols[j][board[i][j]] || squares[sn][board[i][j]] {
				return false
			}

			rows[i][board[i][j]] = true
			cols[j][board[i][j]] = true
			squares[sn][board[i][j]] = true
		}
	}

	return true
}

func isValidSudokuv5(board [][]byte) bool {

	// 9: according to contrains
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	squares := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool, 9)
		cols[i] = make(map[byte]bool, 9)
		squares[i] = make(map[byte]bool, 9)

	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			sn := 3*(i/3) + (j / 3)
			if rows[i][board[i][j]] || cols[j][board[i][j]] || squares[sn][board[i][j]] {
				return false
			}

			rows[i][board[i][j]] = true
			cols[j][board[i][j]] = true
			squares[sn][board[i][j]] = true
		}
	}

	return true
}

func isValidSudokuv4(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	square := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool, 9)
		cols[i] = make(map[byte]bool, 9)
		square[i] = make(map[byte]bool, 9)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			if board[i][j] == '.' {
				continue
			}

			// check duplicated
			sn := 3*(i/3) + (j / 3) // square no.
			if rows[i][board[i][j]] || cols[j][board[i][j]] || square[sn][board[i][j]] {
				return false
			}

			rows[i][board[i][j]] = true
			cols[j][board[i][j]] = true
			square[sn][board[i][j]] = true
		}
	}

	return true

}

func isValidSudokuv3(board [][]byte) bool {
	// according to constaints
	m, n := 9, 9

	rows := make([]map[byte]bool, m)   // rows[m][number]= exist
	cols := make([]map[byte]bool, n)   // cols[n][number]= exist
	square := make([]map[byte]bool, m) // square[m][number] = exist

	for i := 0; i < m; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		square[i] = make(map[byte]bool)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if board[i][j] == '.' {
				//skip
				continue
			}

			sn := 3*(i/3) + (j / 3) // square number
			if rows[i][board[i][j]] || cols[j][board[i][j]] || square[sn][board[i][j]] {
				return false
			}

			// save the number that is visited
			rows[i][board[i][j]] = true
			cols[j][board[i][j]] = true
			square[sn][board[i][j]] = true
		}
	}

	return true
}

func isValidSudokuv2(board [][]byte) bool {
	m := len(board)

	// check the size of board is valid
	for i := 0; i < m; i++ {
		if len(board[i]) != m {
			return false
		}
	}

	// rows[i] has byte {1,...9}; if there exists one element in the row map then it's invalid
	rows := make([]map[byte]bool, m)
	// cols[j] has byte ...
	cols := make([]map[byte]bool, m)
	// 9 squares; square[i] has type ...
	square := make([]map[byte]bool, m)

	for i := 0; i < m; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		square[i] = make(map[byte]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			// skip if (i,j) == '.'
			if board[i][j] == '.' {
				continue
			}

			// check duplicated
			if rows[i][board[i][j]] || cols[j][board[i][j]] || square[3*(i/3)+j/3][board[i][j]] {
				return false
			}

			rows[i][board[i][j]] = true
			cols[j][board[i][j]] = true
			square[3*(i/3)+j/3][board[i][j]] = true

		}
	}

	return true
}

func isValidSudokuv1(board [][]byte) bool {

	m, n := len(board), len(board[0])

	rows := make([]map[byte]bool, m)
	cols := make([]map[byte]bool, n)
	squares := make([]map[byte]bool, m)

	for i := 0; i < m; i++ {
		squares[i] = make(map[byte]bool)
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '.' {
				continue
			}

			if rows[i][board[i][j]] || cols[j][board[i][j]] || squares[3*(i/3)+j/3][board[i][j]] {
				return false
			}

			rows[i][board[i][j]] = true
			cols[j][board[i][j]] = true
			squares[3*(i/3)+j/3][board[i][j]] = true
		}
	}

	return true
}
