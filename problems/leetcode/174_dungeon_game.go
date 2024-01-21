package b75

/*
174. Dungeon Game
Hard
The demons had captured the princess and imprisoned her in the bottom-right corner of a dungeon. The dungeon consists of m x n rooms laid out in a 2D grid. Our valiant knight was initially positioned in the top-left room and must fight his way through dungeon to rescue the princess.

The knight has an initial health point represented by a positive integer. If at any point his health point drops to 0 or below, he dies immediately.

Some of the rooms are guarded by demons (represented by negative integers), so the knight loses health upon entering these rooms; other rooms are either empty (represented as 0) or contain magic orbs that increase the knight's health (represented by positive integers).

To reach the princess as quickly as possible, the knight decides to move only rightward or downward in each step.

Return the knight's minimum initial health so that he can rescue the princess.

Note that any room can contain threats or power-ups, even the first room the knight enters and the bottom-right room where the princess is imprisoned.

Example 1:

Input: dungeon = [[-2,-3,3],[-5,-10,1],[10,30,-5]]
Output: 7
Explanation: The initial health of the knight must be at least 7 if he follows the optimal path: RIGHT-> RIGHT -> DOWN -> DOWN.

Example 2:

Input: dungeon = [[0]]
Output: 1
*/

func calculateMinimumHPv37(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(1, min(dungeon[i][j+1]-dungeon[i][j], dungeon[i+1][j]-dungeon[i][j]))
			}
		}
	}

	return dungeon[0][0]
}

func calculateMinimumHPv36(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(1, min(dungeon[i][j+1]-dungeon[i][j], dungeon[i+1][j]-dungeon[i][j]))
			}
		}
	}

	return dungeon[0][0]
}

func calculateMinimumHPv35(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(1, min(dungeon[i][j+1]-dungeon[i][j], dungeon[i+1][j]-dungeon[i][j]))
			}
		}
	}
	return dungeon[0][0]
}

func calculateMinimumHPv34(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(1, min(dungeon[i][j+1]-dungeon[i][j], dungeon[i+1][j]-dungeon[i][j]))
			}
		}
	}

	return dungeon[0][0]
}

func calculateMinimumHPv33(dungeon [][]int) int {

	m, n := len(dungeon), len(dungeon[0])
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(1, min(dungeon[i][j+1]-dungeon[i][j], dungeon[i+1][j]-dungeon[i][j]))
			}
		}
	}

	return dungeon[0][0]
}

func calculateMinimumHPv32(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(1, min(dungeon[i][j+1]-dungeon[i][j], dungeon[i+1][j]-dungeon[i][j]))
			}
		}
	}
	return dungeon[0][0]
}

func calculateMinimumHPv31(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(1, min(dungeon[i][j+1]-dungeon[i][j], dungeon[i+1][j]-dungeon[i][j]))
			}
		}
	}
	return dungeon[0][0]
}

func calculateMinimumHPv30(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(1, min(dungeon[i][j+1]-dungeon[i][j], dungeon[i+1][j]-dungeon[i][j]))
			}
		}
	}

	return dungeon[0][0]
}

func calculateMinimumHPv29(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(1, min(dungeon[i+1][j]-dungeon[i][j], dungeon[i][j+1]-dungeon[i][j]))
			}
		}
	}
	return dungeon[0][0]
}

func calculateMinimumHPv28(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(1, min(dungeon[i][j+1]-dungeon[i][j], dungeon[i+1][j]-dungeon[i][j]))
			}
		}
	}
	return dungeon[0][0]
}

func calculateMinimumHPv27(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(1, min(dungeon[i][j+1]-dungeon[i][j], dungeon[i+1][j]-dungeon[i][j]))
			}
		}
	}
	return dungeon[0][0]
}

func calculateMinimumHPv26(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(1, min(dungeon[i][j+1]-dungeon[i][j], dungeon[i+1][j]-dungeon[i][j]))
			}
		}
	}

	return dungeon[0][0]
}

func calculateMinimumHPv25(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(1, min(dungeon[i][j+1]-dungeon[i][j], dungeon[i+1][j]-dungeon[i][j]))
			}
		}
	}

	return dungeon[0][0]
}

func calculateMinimumHPv24(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(1, min(dungeon[i+1][j]-dungeon[i][j], dungeon[i][j+1]-dungeon[i][j]))
			}
		}
	}

	return dungeon[0][0]
}

func calculateMinimumHPv23(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(1, min(dungeon[i][j+1]-dungeon[i][j], dungeon[i+1][j]-dungeon[i][j]))
			}
		}
	}

	return dungeon[0][0]
}

func calculateMinimumHPv22(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(1, min(dungeon[i+1][j]-dungeon[i][j], dungeon[i][j+1]-dungeon[i][j]))
			}
		}
	}

	return dungeon[0][0]
}

func calculateMinimumHPv21(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(1, min(dungeon[i][j+1]-dungeon[i][j], dungeon[i+1][j]-dungeon[i][j]))
			}
		}
	}

	return dungeon[0][0]
}

func calculateMinimumHPv20(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(1, min(dungeon[i+1][j]-dungeon[i][j], dungeon[i][j+1]-dungeon[i][j]))
			}
		}
	}

	return dungeon[0][0]
}

func calculateMinimumHPv19(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(1, min(dungeon[i+1][j]-dungeon[i][j], dungeon[i][j+1]-dungeon[i][j]))
			}

		}
	}

	return dungeon[0][0]
}

func calculateMinimumHPv18(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(1, min(dungeon[i+1][j]-dungeon[i][j], dungeon[i][j+1]-dungeon[i][j]))
			}
		}
	}
	return dungeon[0][0]
}

func calculateMinimumHPv17(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(1, min(dungeon[i][j+1]-dungeon[i][j], dungeon[i+1][j]-dungeon[i][j]))
			}
		}
	}

	return dungeon[0][0]
}

func calculateMinimumHPv16(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(1, min(dungeon[i+1][j]-dungeon[i][j], dungeon[i][j+1]-dungeon[i][j]))
			}
		}
	}

	return dungeon[0][0]
}

func calculateMinimumHPv15(dungeon [][]int) int {

	m, n := len(dungeon), len(dungeon[0])

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(1, min(dungeon[i+1][j]-dungeon[i][j], dungeon[i][j+1]-dungeon[i][j]))
			}
		}
	}

	return dungeon[0][0]
}

func calculateMinimumHPv14(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(1,
					min(dungeon[i][j+1]-dungeon[i][j], dungeon[i+1][j]-dungeon[i][j]),
				)
			}
		}
	}

	return dungeon[0][0]
}

func calculateMinimumHPv13(dungeon [][]int) int {

	m, n := len(dungeon), len(dungeon[0])

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][i])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(1, min(
					dungeon[i][j+1]-dungeon[i][j], dungeon[i+1][j]-dungeon[i][j],
				))
			}
		}
	}

	return dungeon[0][0]
}

func calculateMinimumHPv12(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == m-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(1, min(dungeon[i][j+1]-dungeon[i][j], dungeon[i+1][j]-dungeon[i][j]))
			}
		}
	}

	return dungeon[0][0]
}

func calculateMinimumHPv11(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(1, min(dungeon[i+1][j]-dungeon[i][j], dungeon[i][j+1]-dungeon[i][j]))
			}
		}
	}

	return dungeon[0][0]
}

func calculateMinimumHPv10(dungeon [][]int) int {

	m, n := len(dungeon), len(dungeon[0])

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(1,
					min(dungeon[i+1][j]-dungeon[i][j], dungeon[i][j+1]-dungeon[i][j]))

			}
		}
	}

	return dungeon[0][0]

}

func calculateMinimumHPv9(dungeon [][]int) int {

	m, n := len(dungeon), len(dungeon[0])

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(1,
					min(dungeon[i][j+1]-dungeon[i][j], dungeon[i+1][j]-dungeon[i][j]))

			}
		}
	}

	return dungeon[0][0]
}

func calculateMinimumHPv8(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(1,
					min(dungeon[i][j+1]-dungeon[i][j], dungeon[i+1][j]-dungeon[i][j]),
				)
			}
		}
	}

	return dungeon[0][0]
}

func calculateMinimumHPv7(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(1,
					min(dungeon[i+1][j]-dungeon[i][j], dungeon[i][j+1]-dungeon[i][j]),
				)
			}
		}
	}

	return dungeon[0][0]
}

func calculateMinimumHPv6(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 { // start with hp=1
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 { // only one way from j+1 (left)
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 { // only one way from i+1 (up)
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else { // two ways fom left and up and pick the mininum one
				dungeon[i][j] = max(1,
					min(dungeon[i][j+1]-dungeon[i][j], dungeon[i+1][j]-dungeon[i][j]),
				)
			}
		}
	}

	return dungeon[0][0]
}

func calculateMinimumHPv5(dungeon [][]int) int {

	m, n := len(dungeon), len(dungeon[0])

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][j]) // base case: 1-dungeon[m-1][n-1]
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(1,
					min(dungeon[i+1][j]-dungeon[i][j], dungeon[i][j+1]-dungeon[i][j]),
				)
			}
		}
	}

	return dungeon[0][0]

}

func calculateMinimumHPv4(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(
					1,
					min(dungeon[i][j+1]-dungeon[i][j], dungeon[i+1][j]-dungeon[i][j]),
				)
			}
		}
	}

	return dungeon[0][0]
}

func calculateMinimumHPv3(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][i])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(
					1,
					min(dungeon[i][j+1]-dungeon[i][j], dungeon[i+1][j]-dungeon[i][j]),
				)
			}
		}

	}

	return dungeon[0][0]
}

func calculateMinimumHPv2(dungeon [][]int) int {

	m, n := len(dungeon), len(dungeon[0])

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			} else {
				dungeon[i][j] = max(1,
					min(dungeon[i][j+1]-dungeon[i][j], dungeon[i+1][j]-dungeon[i][j]),
				)
			}
		}
	}

	return dungeon[0][0]
}

/* bad implementation
func calculateMinimumHPv1(dungeon [][]int) int {
	if dungeon == nil {
		return 0
	}
	m := len(dungeon)
	n := len(dungeon[0])

	dp := make([][]struct{ need, value int }, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]struct{ need, value int }, n)
	}

	if dungeon[0][0] <= 0 {
		dp[0][0].need = (0 - dungeon[0][0]) + 1
		dp[0][0].value = dungeon[0][0] + dp[0][0].need
	} else {
		dp[0][0].need = 0
		dp[0][0].value = dungeon[0][0]
	}

	// cols
	for i := 1; i < n; i++ {
		if dp[0][i-1].value+dungeon[0][i] <= 0 {
			dp[0][i].need = (0 - (dp[0][i-1].value + dungeon[0][i])) + 1 + dp[0][i-1].need
			dp[0][i].value = (dp[0][i-1].value + dungeon[0][i]) + (0 - (dp[0][i-1].value + dungeon[0][i])) + 1
		} else {
			dp[0][i].need = dp[0][i-1].need
			dp[0][i].value = dungeon[0][i] + dp[0][i-1].value
		}
	}

	// rows
	for i := 1; i < m; i++ {
		if dp[i-1][0].value+dungeon[i][0] <= 0 {
			dp[i][0].need = (0 - (dp[i-1][0].value + dungeon[i][0])) + 1 + dp[i-1][0].need
			dp[i][0].value = (dp[i-1][0].value + dungeon[i][0]) + (0 - (dp[i-1][0].value + dungeon[i][0])) + 1
		} else {
			dp[i][0].need = dp[i-1][0].need
			dp[i][0].value = dungeon[i][0] + dp[i-1][0].value
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			var k struct{ need, value int }
			// choice a best one
			if dp[i][j-1].need < dp[i-1][j].need {
				k.need = dp[i][j-1].need
				k.value = dp[i][j-1].value
			} else {
				k.need = dp[i-1][j].need
				k.value = dp[i-1][j].value
			}
			if k.value+dungeon[i][j] <= 0 {
				dp[i][j].need = (0 - (k.value + dungeon[i][j])) + 1 + k.need
				dp[i][j].value = (k.value + dungeon[i][0]) + (0 - (k.value + dungeon[i][0])) + 1
			} else {
				dp[i][j].need = k.need
				dp[i][j].value = dungeon[i][j] + k.value
			}
		}
	}

	if n-2 < 0 {
		return dp[m-1][n-1].need
	}

	fromLeftSideMin := 0 - (dungeon[m-1][n-1] + dp[m-1][n-2].value) + 1

	if fromLeftSideMin <= 0 {
		fromLeftSideMin = -fromLeftSideMin
	}

	// return dp[m-1][n-1].need
	return min(dp[m-1][n-1].need, fromLeftSideMin)
}
*/
