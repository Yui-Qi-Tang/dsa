package b75

/*
100. Same Tree
Easy
9.2K
186
Companies
Given the roots of two binary trees p and q, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.



Example 1:


Input: p = [1,2,3], q = [1,2,3]
Output: true
Example 2:


Input: p = [1,2], q = [1,null,2]
Output: false
Example 3:


Input: p = [1,2,1], q = [1,1,2]
Output: false


Constraints:

The number of nodes in both trees is in the range [0, 100].
-104 <= Node.val <= 104
*/

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTreev21(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTreev21(p.Left, q.Left) && isSameTreev21(p.Right, q.Right)
}

func isSameTreev20(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTreev20(p.Left, q.Left) && isSameTreev20(p.Right, q.Right)
}

func isSameTreev19(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTreev19(p.Left, q.Left) && isSameTreev19(p.Right, q.Right)
}

func isSameTreev18(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTreev18(p.Left, q.Left) && isSameTreev18(q.Right, p.Right)
}

func isSameTreev17(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTreev17(p.Left, q.Left) && isSameTreev17(p.Right, q.Right)
}

func isSameTreev16(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTreev16(p.Left, q.Left) && isSameTreev16(p.Right, q.Right)
}

func isSameTreev15(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTreev15(p.Left, q.Left) && isSameTreev15(p.Right, q.Right)
}

func isSameTreev14(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTreev14(p.Left, q.Left) && isSameTreev14(p.Right, q.Right)
}

func isSameTreev13(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTreev13(p.Left, q.Left) && isSameTreev13(p.Right, q.Right)
}

func isSameTreev12(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTreev12(p.Left, q.Left) && isSameTreev12(p.Right, q.Right)
}

func isSameTreev11(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTreev11(p.Left, q.Left) && isSameTreev11(p.Right, q.Right)
}

func isSameTreev10(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTreev10(p.Left, q.Left) && isSameTreev10(p.Right, q.Right)
}

func isSameTreev9(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTreev9(p.Left, q.Left) && isSameTreev9(p.Right, q.Right)
}

func isSameTreev8(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTreev8(p.Left, q.Left) && isSameTreev8(p.Right, q.Right)
}

func isSameTreev7(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTreev7(p.Left, q.Left) && isSameTreev7(p.Right, q.Right)
}

func isSameTreev6(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTreev6(p.Left, q.Left) && isSameTreev6(p.Right, q.Right)
}

func isSameTreev5(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTreev5(p.Left, q.Left) && isSameTreev5(p.Right, q.Right)
}

func isSameTreev4(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTreev4(p.Left, q.Left) && isSameTreev4(p.Right, q.Right)
}

func isSameTreev3(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTreev3(p.Left, q.Left) && isSameTreev3(p.Right, q.Right)
}

func isSameTreev2(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p != nil && q == nil {
		return false
	}

	if p == nil && q != nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTreev2(p.Left, q.Left) && isSameTreev2(p.Right, q.Right)
}

func isSameTreev1(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil && q != nil {
		return false
	}

	if p != nil && q == nil {
		return false
	}

	if p != nil && q != nil && (p.Val != q.Val) {
		return false
	}
	return isSameTreev1(p.Left, q.Left) && isSameTreev1(p.Right, q.Right)
}
