package b75

import (
	"math"
)

/*
124. Binary Tree Maximum Path Sum
Hard
14.1K
646
Companies
A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them. A node can only appear in the sequence at most once. Note that the path does not need to pass through the root.

The path sum of a path is the sum of the node's values in the path.

Given the root of a binary tree, return the maximum path sum of any non-empty path.



Example 1:


Input: root = [1,2,3]
Output: 6
Explanation: The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 = 6.
Example 2:


Input: root = [-10,9,20,null,null,15,7]
Output: 42
Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 + 7 = 42.


Constraints:

The number of nodes in the tree is in the range [1, 3 * 104].
-1000 <= Node.val <= 1000
Accepted
970.2K
Submissions
2.5M
Acceptance Rate
39.3%
*/

/*
*
Definition for a binary tree node.
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSumv10(root *TreeNode) int {

	result := 0
	var maxsum func(n *TreeNode) int
	maxsum = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		lmax := max(0, maxsum(n.Left))
		rmax := max(0, maxsum(n.Right))
		result = max(result, n.Val+lmax+rmax)
		return n.Val + max(lmax, rmax)
	}
	maxsum(root)
	return result
}

func maxPathSumv9(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := 0

	var maxsum func(n *TreeNode) int
	maxsum = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		lmax := max(0, maxsum(n.Left))
		rmax := max(0, maxsum(n.Right))
		result = max(result, n.Val+lmax+rmax)
		return n.Val + max(lmax, rmax)
	}

	maxsum(root)

	return result
}

func maxPathSumv8(root *TreeNode) int {
	result := 0

	var maxsum func(n *TreeNode) int

	maxsum = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		lmax := max(0, maxsum(n.Left))
		rmax := max(0, maxsum(n.Right))

		result = max(result, n.Val+lmax+rmax)

		return n.Val + max(lmax, rmax)
	}

	maxsum(root)

	return result
}

func maxPathSumv7(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := 0
	var maxsum func(n *TreeNode) int

	maxsum = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		lmax := max(0, maxsum(n.Left))
		rmax := max(0, maxsum(n.Right))
		result = max(result, n.Val+lmax+rmax)
		return n.Val + max(lmax, rmax)
	}

	maxsum(root)
	return result
}

func maxPathSumv6(root *TreeNode) int {
	result := 0

	var maxsum func(n *TreeNode) int
	maxsum = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		lmax := max(0, maxsum(n.Left))
		rmax := max(0, maxsum(n.Right))
		result = max(result, n.Val+lmax+rmax)
		return n.Val + max(lmax, rmax)
	}
	maxsum(root)

	return result
}

func maxPathSumv5(root *TreeNode) int {

	result := 0
	var maxsum func(n *TreeNode) int
	maxsum = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		lmax := max(0, maxsum(n.Left))
		rmax := max(0, maxsum(n.Right))
		result = max(result, n.Val+lmax+rmax)
		return n.Val + max(lmax, rmax)
	}

	maxsum(root)

	return result
}

func maxPathSumv4(root *TreeNode) int {

	result := 0

	var maxSum func(n *TreeNode) int
	maxSum = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		lmax := max(0, maxSum(n.Left))
		rmax := max(0, maxSum(n.Right))
		result = max(result, n.Val+lmax+rmax)
		return n.Val + max(lmax, rmax)
	}

	maxSum(root)

	return result
}

func maxPathSumv3(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := 0
	var maxSum func(n *TreeNode) int

	maxSum = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		// 0 if maxsum(node) is nagetive number
		lmax := max(0, maxSum(n.Left))
		rmax := max(0, maxSum(n.Right))
		result = max(result, n.Val+lmax+rmax)
		return n.Val + max(lmax, rmax)
	}

	maxSum(root)

	return result
}

func maxPathSumv2(root *TreeNode) int {

	result := -1001
	var dfs func(n *TreeNode) int
	dfs = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		lmax := max(0, dfs(n.Left))
		rmax := max(0, dfs(n.Right))
		result = max(result, n.Val+lmax+rmax)
		return n.Val + max(lmax, rmax)
	}

	dfs(root)

	return result
}

func maxPathSumv1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := math.MinInt

	var helper func(n *TreeNode) int

	helper = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		lmax := max(0, helper(n.Left))
		rmax := max(0, helper(n.Right))
		result = max(result, n.Val+lmax+rmax)
		return n.Val + max(lmax, rmax)
	}
	helper(root)

	return result
}

func maxPathSum(root *TreeNode) int {

	var result int = math.MinInt32
	p := &result

	var helper func(n *TreeNode, maxSum *int) int
	helper = func(n *TreeNode, maxSum *int) int {
		if n == nil {
			return 0
		}
		lmax := max(0, helper(n.Left, maxSum))
		rmax := max(0, helper(n.Right, maxSum))
		*maxSum = max(*maxSum, n.Val+lmax+rmax)
		return n.Val + max(lmax, rmax)
	}

	helper(root, p)

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
