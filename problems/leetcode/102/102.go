package b75

/*
102. Binary Tree Level Order Traversal
Medium
12.9K
257
Companies
Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).



Example 1:


Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]
Example 2:

Input: root = [1]
Output: [[1]]
Example 3:

Input: root = []
Output: []


Constraints:

The number of nodes in the tree is in the range [0, 2000].
-1000 <= Node.val <= 1000
Accepted
1.8M
Submissions
2.7M

*/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderv20(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		qLen := len(queue)
		lvs := []int{}
		for qLen > 0 {
			qLen--
			n := queue[0]
			queue = queue[1:]
			lvs = append(lvs, n.Val)
			if n.Left != nil {
				queue = append(queue, n.Left)
			}

			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}

		result = append(result, lvs)
	}

	return result
}

func levelOrderv19(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		qLen := len(queue)
		lvs := []int{}

		for qLen > 0 {
			qLen--
			n := queue[0]
			queue = queue[1:]
			lvs = append(lvs, n.Val)
			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}

		result = append(result, lvs)
	}

	return result
}

func levelOrderv18(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		lvs := []int{}
		qLen := len(queue)
		for qLen > 0 {
			n := queue[0]
			queue = queue[1:]
			lvs = append(lvs, n.Val)
			qLen--
			if n.Left != nil {
				queue = append(queue, n.Left)
			}

			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}

		result = append(result, lvs)
	}

	return result
}

func levelOrderv17(root *TreeNode) [][]int {
	result := [][]int{}

	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		lvs := []int{}
		qLen := len(queue)
		for qLen > 0 {
			qLen--
			n := queue[0]
			lvs = append(lvs, n.Val)
			queue = queue[1:]
			if n.Left != nil {
				queue = append(queue, n.Left)
			}

			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}

		result = append(result, lvs)
	}

	return result
}

func levelOrderv16(root *TreeNode) [][]int {
	result := [][]int{}

	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		qLen := len(queue)
		lvs := []int{}
		for qLen > 0 {
			n := queue[0]
			queue = queue[1:]
			qLen--

			lvs = append(lvs, n.Val)
			if n.Left != nil {
				queue = append(queue, n.Left)
			}

			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}

		result = append(result, lvs)
	}

	return result
}

func levelOrderv15(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		lvs := []int{}

		qLen := len(queue)
		for qLen > 0 {
			n := queue[0]
			queue = queue[1:]
			qLen--
			lvs = append(lvs, n.Val)
			if n.Left != nil {
				queue = append(queue, n.Left)
			}

			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}

		result = append(result, lvs)
	}

	return result
}

func levelOrderv14(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		lvs := []int{}
		qLen := len(queue)

		for qLen > 0 {

			n := queue[0]
			queue = queue[1:]
			lvs = append(lvs, n.Val)

			if n.Left != nil {
				queue = append(queue, n.Left)
			}

			if n.Right != nil {
				queue = append(queue, n.Right)
			}

			qLen--
		}

		result = append(result, lvs)

	}

	return result
}

func levelOrderv13(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		lvs := []int{}
		qLen := len(queue)
		for qLen > 0 {
			qLen--

			n := queue[0]
			queue = queue[1:]
			lvs = append(lvs, n.Val)
			if n.Left != nil {
				queue = append(queue, n.Left)
			}

			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}

		result = append(result, lvs)
	}

	return result
}

func levelOrderv12(root *TreeNode) [][]int {
	result := [][]int{}

	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		qLen := len(queue)
		lvs := make([]int, 0, qLen)
		for qLen > 0 {
			qLen--

			n := queue[0]
			queue = queue[1:]
			lvs = append(lvs, n.Val)

			if n.Left != nil {
				queue = append(queue, n.Left)
			}

			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}
		result = append(result, lvs)
	}

	return result
}

func levelOrderv11(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		qLen := len(queue)
		lvs := []int{}
		for qLen > 0 {
			qLen--

			n := queue[0]
			lvs = append(lvs, n.Val)
			queue = queue[1:]

			if n.Left != nil {
				queue = append(queue, n.Left)
			}

			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}
		result = append(result, lvs)
	}

	return result
}

func levelOrderv10(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	q := []*TreeNode{root}
	result := [][]int{}
	for len(q) > 0 {
		qLen := len(q)
		lvs := []int{}
		for qLen > 0 {

			n := q[0]
			q = q[1:]

			lvs = append(lvs, n.Val)
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}

			qLen--
		}

		result = append(result, lvs)
	}

	return result
}

func levelOrderv9(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}
	result := [][]int{}

	for len(queue) > 0 {
		qlen := len(queue)
		lvs := []int{}

		for qlen > 0 {
			n := queue[0]
			queue = queue[1:]

			lvs = append(lvs, n.Val)
			if n.Left != nil {
				queue = append(queue, n.Left)
			}

			if n.Right != nil {
				queue = append(queue, n.Right)
			}

			qlen--
		}

		result = append(result, lvs)
	}

	return result
}

func levelOrderv8(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		qLen := len(queue)
		levelValues := []int{}
		for qLen > 0 {
			n := queue[0]
			queue = queue[1:]
			qLen--
			levelValues = append(levelValues, n.Val)
			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}

		result = append(result, levelValues)
	}

	return result
}

func levelOrderv7(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := make([][]int, 0)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		qLen := len(queue)
		lvs := make([]int, 0) // lvs: level values
		for qLen > 0 {
			n := queue[0]
			queue = queue[1:]
			lvs = append(lvs, n.Val)
			if n.Left != nil {
				queue = append(queue, n.Left)
			}

			if n.Right != nil {
				queue = append(queue, n.Right)
			}
			qLen--
		}

		result = append(result, lvs)
	}

	return result
}

func levelOrderv6(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		levelValues := make([]int, 0)
		qLen := len(queue)
		for qLen > 0 {
			n := queue[0]
			queue = queue[1:]
			levelValues = append(levelValues, n.Val)

			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}

			qLen--
		}

		result = append(result, levelValues)
	}

	return result
}

func levelOrderv5(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		qLen := len(queue)
		levelValues := make([]int, 0, qLen)
		for qLen > 0 {
			n := queue[0]
			queue = queue[1:]
			levelValues = append(levelValues, n.Val)
			qLen--

			if n.Left != nil {
				queue = append(queue, n.Left)
			}

			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}
		result = append(result, levelValues)
	}

	return result
}

func levelOrderv4(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		levelValues := make([]int, 0)
		qLen := len(queue)
		for qLen > 0 {
			n := queue[0]
			queue = queue[1:]
			levelValues = append(levelValues, n.Val)
			if n.Left != nil {
				queue = append(queue, n.Left)
			}

			if n.Right != nil {
				queue = append(queue, n.Right)
			}
			qLen--
		}

		result = append(result, levelValues)
	}

	return result
}

func levelOrderv3(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		qLen := len(queue)
		levelValues := make([]int, 0, qLen)
		for qLen > 0 {
			n := queue[0]
			queue = queue[1:]
			levelValues = append(levelValues, n.Val)
			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
			qLen--
		}
		result = append(result, levelValues)
	}

	return result
}

func levelOrderv2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		qLen := len(queue)
		levelValues := make([]int, 0, qLen)

		for qLen > 0 {
			n := queue[0]
			queue = queue[1:]
			levelValues = append(levelValues, n.Val)
			if n.Left != nil {
				queue = append(queue, n.Left)
			}

			if n.Right != nil {
				queue = append(queue, n.Right)
			}
			qLen--
		}
		result = append(result, levelValues)
	}

	return result
}

func levelOrderv1(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		qLen := len(queue)
		levelValues := make([]int, 0, qLen)
		for qLen > 0 {
			n := queue[0]
			queue = queue[1:]
			levelValues = append(levelValues, n.Val)
			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
			qLen--
		}
		result = append(result, levelValues)
	}
	return result
}

func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		levelValues := make([]int, 0, len(queue))
		qLen := len(queue)
		for qLen > 0 {
			n := queue[0]
			queue = queue[1:]
			levelValues = append(levelValues, n.Val)
			if n.Left != nil {
				queue = append(queue, n.Left)
			}

			if n.Right != nil {
				queue = append(queue, n.Right)
			}
			qLen--
		}
		result = append(result, levelValues)
	}

	return result
}
