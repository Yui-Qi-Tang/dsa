package b75

/*
104. Maximum Depth of Binary Tree
Easy
10.5K
167
Companies
Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.



Example 1:


Input: root = [3,9,20,null,null,15,7]
Output: 3
Example 2:

Input: root = [1,null,2]
Output: 2


Constraints:

The number of nodes in the tree is in the range [0, 104].
-100 <= Node.val <= 100
Accepted
2.3M
Submissions
3.1M
Acceptance Rate
73.9%
*/

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepthv25(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepthv25(root.Left), maxDepthv25(root.Right))
}

func maxDepthv24(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepthv24(root.Left), maxDepthv24(root.Right))
}

func maxDepthv23(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepthv23(root.Left), maxDepthv23(root.Right))
}

func maxDepthv22(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepthv22(root.Left), maxDepthv22(root.Right))
}

func maxDepthv21(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepthv21(root.Left), maxDepthv21(root.Right))
}

func maxDepthv20(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepthv20(root.Left), maxDepthv20(root.Right))
}

func maxDepthv19(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepthv19(root.Left), maxDepthv19(root.Right))
}

func maxDepthv18(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepthv18(root.Left), maxDepthv18(root.Right))
}

func maxDepthv17(root *TreeNode) int {

	if root == nil {
		return 0
	}

	return 1 + max(maxDepthv17(root.Left), maxDepthv17(root.Right))
}

func maxDepthv16(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepthv16(root.Left), maxDepthv16(root.Right))
}

func maxDepthv15(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepthv15(root.Left), maxDepthv15(root.Right))
}

func maxDepthv14(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepthv14(root.Left), maxDepthv14(root.Right))
}

func maxDepthv13(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepthv13(root.Left), maxDepthv13(root.Right))
}

func maxDepthv12(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepthv12(root.Left), maxDepthv12(root.Right))
}

func maxDepthv11(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := maxDepthv11(root.Left)
	r := maxDepthv11(root.Right)

	return 1 + max(l, r)
}

func maxDepthv10(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepthv10(root.Left)
	r := maxDepthv10(root.Right)
	return 1 + max(l, r)
}

func maxDepthv9(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := maxDepthv9(root.Left)
	r := maxDepthv9(root.Right)
	return 1 + max(l, r)
}

func maxDepthv8(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := maxDepthv8(root.Left)
	r := maxDepthv8(root.Right)
	return 1 + max(l, r)
}

func maxDepthv7(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := maxDepthv7(root.Left)
	r := maxDepthv7(root.Right)
	return 1 + max(l, r)
}

func maxDepthv6(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lmax := maxDepthv6(root.Left)
	rmax := maxDepthv6(root.Right)

	return 1 + max(lmax, rmax)
}

func maxDepthv5(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := maxDepthv5(root.Left)
	r := maxDepthv5(root.Right)

	return 1 + max(l, r)
}

func maxDepthv4(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := maxDepthv4(root.Left)
	r := maxDepthv4(root.Right)
	return 1 + max(l, r)
}

func maxDepthv3(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepthv3(root.Left), maxDepthv3(root.Right))
}

func maxDepthv2(root *TreeNode) int {

	level := 0
	if root == nil {
		return level
	}

	var dfs func(*TreeNode, int)
	dfs = func(n *TreeNode, l int) {
		if n == nil {
			return
		}
		dfs(n.Left, l+1)
		dfs(n.Right, l+1)
		level = max(level, l)
	}

	dfs(root, 0)
	return level + 1
}

func maxDepthv1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	level := 0
	var dfs func(n *TreeNode, level int)
	dfs = func(n *TreeNode, lv int) {
		if n == nil {
			return
		}

		dfs(n.Left, lv+1)
		dfs(n.Right, lv+1)
		level = max(lv, level)
	}
	dfs(root, 0)
	return level + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
