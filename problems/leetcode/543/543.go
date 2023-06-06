package b75

/*
543. Diameter of Binary Tree
Easy
11.1K
698
Companies
Given the root of a binary tree, return the length of the diameter of the tree.

The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.

The length of a path between two nodes is represented by the number of edges between them.



Example 1:


Input: root = [1,2,3,4,5]
Output: 3
Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].
Example 2:

Input: root = [1,2]
Output: 1


Constraints:

The number of nodes in the tree is in the range [1, 104].
-100 <= Node.val <= 100
*/

/**
* Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTreev24(root *TreeNode) int {

	result := 0
	var md func(n *TreeNode) int
	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		l := md(n.Left)
		r := md(n.Right)
		result = max(result, l+r)
		return 1 + max(l, r)
	}
	md(root)
	return result
}

func diameterOfBinaryTreev23(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := 0
	var md func(n *TreeNode) int
	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		l := md(n.Left)
		r := md(n.Right)
		result = max(result, l+r)
		return 1 + max(l, r)
	}

	md(root)

	return result
}

func diameterOfBinaryTreev22(root *TreeNode) int {
	result := 0

	var md func(n *TreeNode) int
	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		lmax := md(n.Left)
		rmax := md(n.Right)
		result = max(result, lmax+rmax)
		return 1 + max(lmax, rmax)
	}

	md(root)
	return result
}

func diameterOfBinaryTreev21(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := 0
	var md func(n *TreeNode) int

	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		ld := md(n.Left)
		rd := md(n.Right)
		result = max(result, ld+rd)
		return 1 + max(ld, rd)
	}

	md(root)
	return result
}

func diameterOfBinaryTreev20(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := 0

	var md func(n *TreeNode) int
	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		lmax := md(n.Left)
		rmax := md(n.Right)
		result = max(result, lmax+rmax)
		return 1 + max(lmax, rmax)
	}

	md(root)

	return result
}

func diameterOfBinaryTreev19(root *TreeNode) int {
	result := 0
	var md func(n *TreeNode) int
	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		ld := md(n.Left)
		rd := md(n.Right)
		result = max(result, ld+rd)
		return 1 + max(ld, rd)
	}

	md(root)
	return result
}

func diameterOfBinaryTreev18(root *TreeNode) int {
	result := 0
	if root == nil {
		return result
	}

	var md func(n *TreeNode) int
	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		lmax := md(n.Left)
		rmax := md(n.Right)

		result = max(result, lmax+rmax)
		return 1 + max(lmax, rmax)
	}
	md(root)
	return result
}

func diameterOfBinaryTreev17(root *TreeNode) int {

	result := 0

	var md func(n *TreeNode) int
	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		l := md(n.Left)
		r := md(n.Right)
		result = max(result, l+r)
		return 1 + max(l, r)
	}

	md(root)

	return result
}

func diameterOfBinaryTreev16(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := 0

	var md func(n *TreeNode) int

	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		l := md(n.Left)
		r := md(n.Right)
		result = max(result, l+r)
		return 1 + max(l, r)
	}

	md(root)

	return result
}

func diameterOfBinaryTreev15(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := 0
	var md func(n *TreeNode) int
	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		l := md(n.Left)
		r := md(n.Right)
		result = max(result, l+r)
		return 1 + max(l, r)
	}
	md(root)

	return result
}

func diameterOfBinaryTreev14(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result := 0

	var md func(n *TreeNode) int
	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		l := md(n.Left)
		r := md(n.Right)
		result = max(result, l+r)
		return 1 + max(l, r)
	}
	md(root)

	return result
}

func diameterOfBinaryTreev13(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var md func(n *TreeNode) int

	result := 0
	md = func(n *TreeNode) int {

		if n == nil {
			return 0
		}

		l := md(n.Left)
		r := md(n.Right)
		result = max(result, l+r)
		return 1 + max(l, r)
	}

	md(root)
	return result
}

func diameterOfBinaryTreev12(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := 0

	var md func(n *TreeNode) int
	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		l := md(n.Left)
		r := md(n.Right)
		result = max(result, l+r)
		return 1 + max(l, r)
	}

	md(root)

	return result
}

func diameterOfBinaryTreev11(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := 0

	var md func(n *TreeNode) int

	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		l := md(n.Left)
		r := md(n.Right)
		result = max(result, l+r)
		return 1 + max(l, r)
	}

	md(root)

	return result
}

func diameterOfBinaryTreev10(root *TreeNode) int {
	result := 0

	var md func(n *TreeNode) int

	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		l := md(n.Left)
		r := md(n.Right)
		result = max(result, l+r)
		return 1 + max(l, r)
	}

	md(root)
	return result
}

func diameterOfBinaryTreev9(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := 0

	var md func(n *TreeNode) int
	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		l := md(n.Left)
		r := md(n.Right)
		result = max(result, l+r)
		return 1 + max(l, r)
	}

	md(root)

	return result
}

func diameterOfBinaryTreev8(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := 0

	var md func(n *TreeNode) int

	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		l := md(n.Left)
		r := md(n.Right)
		result = max(result, l+r)
		return 1 + max(l, r)
	}
	md(root)

	return result
}

func diameterOfBinaryTreev7(root *TreeNode) int {
	result := 0

	var md func(n *TreeNode) int
	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		l := md(n.Left)
		r := md(n.Right)
		result = max(result, l+r)
		return 1 + max(l, r)
	}
	md(root)
	return result
}

func diameterOfBinaryTreev6(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var md func(n *TreeNode) int
	result := 0
	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		l := md(n.Left)
		r := md(n.Right)
		result = max(result, l+r)
		return 1 + max(l, r)
	}

	md(root)

	return result
}

func diameterOfBinaryTreev5(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := 0

	var dfs func(n *TreeNode) int
	dfs = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		l := dfs(n.Left)
		r := dfs(n.Right)
		result = max(result, l+r) // combine the depth of each child
		return 1 + max(l, r)
	}
	dfs(root)

	return result
}

func diameterOfBinaryTreev4(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := 0

	var dfs func(n *TreeNode) int
	dfs = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		l := dfs(n.Left)
		r := dfs(n.Right)
		result = max(result, l+r)
		return 1 + max(l, r)
	}
	dfs(root)

	return result
}

func diameterOfBinaryTreev3(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := 0
	var dfs func(tn *TreeNode) int

	dfs = func(tn *TreeNode) int {
		if tn == nil {
			return 0
		}

		l := dfs(tn.Left)
		r := dfs(tn.Right)
		result = max(result, l+r)
		return 1 + max(l, r)
	}

	dfs(root)

	return result
}

func diameterOfBinaryTreev2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := 0
	var dfs func(n *TreeNode) int
	dfs = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		l := dfs(n.Left)
		r := dfs(n.Right)
		result = max(result, l+r)
		return max(l, r) + 1
	}

	dfs(root)

	return result
}

func diameterOfBinaryTreev1(root *TreeNode) int {

	if root == nil {
		return 0
	}

	var dfs func(*TreeNode) int
	result := 0
	dfs = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		l := dfs(n.Left)
		r := dfs(n.Right)
		result = max(result, l+r)

		return max(l, r) + 1
	}
	dfs(root)

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
