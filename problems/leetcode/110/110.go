package b75

/*
110. Balanced Binary Tree
Easy
8.7K
490
Companies
Given a binary tree, determine if it is
height-balanced
.



Example 1:


Input: root = [3,9,20,null,null,15,7]
Output: true
Example 2:


Input: root = [1,2,2,3,3,null,null,4,4]
Output: false
Example 3:

Input: root = []
Output: true


Constraints:

The number of nodes in the tree is in the range [0, 5000].
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

func isBalancedv24(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var md func(n *TreeNode) int
	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		return 1 + max(md(n.Left), md(n.Right))
	}

	abs := func(x int) int {
		if x < 0 {
			return -x
		}

		return x
	}

	if abs(md(root.Left)-md(root.Right)) > 1 {
		return false
	}

	return isBalancedv24(root.Left) && isBalancedv24(root.Right)
}

func isBalancedv23(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var md func(n *TreeNode) int
	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		return 1 + max(md(n.Left), md(n.Right))
	}

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	if abs(md(root.Left)-md(root.Right)) > 1 {
		return false
	}

	return isBalancedv23(root.Left) && isBalancedv23(root.Right)
}

func isBalancedv22(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var md func(n *TreeNode) int
	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		return 1 + max(md(n.Right), md(n.Left))
	}

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	if abs(md(root.Left)-md(root.Right)) > 1 {
		return false
	}

	return isBalancedv22(root.Left) && isBalancedv22(root.Right)
}

func isBalancedv21(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var md func(n *TreeNode) int
	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		return 1 + max(md(n.Left), md(n.Right))
	}

	abs := func(x int) int {
		if x < 0 {
			return -x
		}

		return x
	}

	if abs(md(root.Left)-md(root.Right)) > 1 {
		return false
	}

	return isBalancedv21(root.Left) && isBalancedv21(root.Right)
}

func isBalancedv20(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var md func(n *TreeNode) int
	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		return 1 + max(md(n.Left), md(n.Right))
	}

	abs := func(x int) int {
		if x < 0 {
			return -x
		}

		return x
	}

	if abs(md(root.Left)-md(root.Right)) > 1 {
		return false
	}

	return isBalancedv20(root.Left) && isBalancedv20(root.Right)
}

func isBalancedv19(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var md func(n *TreeNode) int

	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		return 1 + max(md(n.Left), md(n.Right))
	}

	abs := func(x int) int {
		if x < 0 {
			return -x
		}

		return x
	}

	if abs(md(root.Left)-md(root.Right)) > 1 {
		return false
	}

	return isBalancedv19(root.Left) && isBalancedv19(root.Right)
}

func isBalancedv18(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var md func(n *TreeNode) int
	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		return 1 + max(md(n.Left), md(n.Right))
	}

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	if abs(md(root.Left)-md(root.Right)) > 1 {
		return false
	}

	return isBalancedv18(root.Left) && isBalancedv18(root.Right)
}

func isBalancedv17(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var md func(n *TreeNode) int

	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		return 1 + max(md(n.Left), md(n.Right))
	}

	abs := func(x int) int {
		if x < 0 {
			return -x
		}

		return x
	}

	if abs(md(root.Left)-md(root.Right)) > 1 {
		return false
	}

	return isBalancedv17(root.Left) && isBalancedv17(root.Right)
}

func isBalancedv16(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var md func(n *TreeNode) int
	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		return 1 + max(md(n.Left), md(n.Right))
	}

	abs := func(x int) int {
		if x < 0 {
			return -x
		}

		return x
	}

	if abs(md(root.Left)-md(root.Right)) > 1 {
		return false
	}

	return isBalancedv16(root.Left) && isBalancedv16(root.Right)
}

func isBalancedv15(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var md func(n *TreeNode) int

	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		return 1 + max(md(n.Left), md(n.Right))
	}

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	if abs(md(root.Left)-md(root.Right)) > 1 {
		return false
	}

	return isBalancedv15(root.Left) && isBalancedv15(root.Right)
}

func isBalancedv14(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var md func(n *TreeNode) int
	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		return 1 + max(md(n.Left), md(n.Right))
	}

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	if abs(md(root.Left)-md(root.Right)) > 1 {
		return false
	}

	return isBalancedv14(root.Left) && isBalancedv14(root.Right)
}

func isBalancedv13(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var md func(n *TreeNode) int
	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		l := md(n.Left)
		r := md(n.Right)
		return 1 + max(l, r)
	}

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	if abs(md(root.Left)-md(root.Right)) > 1 {
		return false
	}

	return isBalancedv13(root.Left) && isBalancedv13(root.Right)
}

func isBalancedv12(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var md func(n *TreeNode) int
	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		return 1 + max(md(n.Left), md(n.Right))
	}

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	if abs(md(root.Left)-md(root.Right)) > 1 {
		return false
	}

	return isBalancedv12(root.Left) && isBalancedv12(root.Right)
}

func isBalancedv11(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var md func(n *TreeNode) int
	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		return 1 + max(md(n.Left), md(n.Right))
	}

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	if abs(md(root.Left)-md(root.Right)) > 1 {
		return false
	}

	return isBalancedv11(root.Left) && isBalancedv11(root.Right)
}

func isBalancedv10(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var md func(n *TreeNode) int
	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		return 1 + max(md(n.Left), md(n.Right))
	}

	l := md(root.Left)
	r := md(root.Right)

	abs := func(a int) int {
		if a < 0 {
			return -a
		}

		return a
	}

	if abs(l-r) > 1 {
		return false
	}

	return isBalancedv10(root.Left) && isBalancedv10(root.Right)
}

func isBalancedv9(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var md func(n *TreeNode) int

	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		l := md(n.Left)
		r := md(n.Right)
		return 1 + max(l, r)
	}

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	l := md(root.Left)
	r := md(root.Right)

	if abs(l-r) > 1 {
		return false
	}

	return isBalancedv9(root.Left) && isBalancedv9(root.Right)
}

func isBalancedv8(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var md func(n *TreeNode) int
	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		l := md(n.Left)
		r := md(n.Right)
		return 1 + max(l, r)
	}

	l, r := md(root.Left), md(root.Right)

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	if abs(l-r) > 1 {
		return false
	}

	return isBalancedv8(root.Left) && isBalancedv8(root.Right)
}

func isBalancedv7(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var md func(n *TreeNode) int // max depth
	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		l := md(n.Left)
		r := md(n.Right)

		return 1 + max(l, r)
	}

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	l := md(root.Left)
	r := md(root.Right)

	if abs(l-r) > 1 {
		return false
	}

	return isBalancedv7(root.Left) && isBalancedv7(root.Right)
}

func isBalancedv6(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var md func(n *TreeNode) int
	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		l := md(n.Left)
		r := md(n.Right)
		return 1 + max(l, r)
	}

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	lnum := md(root.Left)
	rnum := md(root.Right)
	if abs(lnum-rnum) > 1 {
		return false
	}

	return isBalancedv6(root.Left) && isBalancedv6(root.Right)
}
func isBalancedv5(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// md: maxDepth
	var md func(n *TreeNode) int
	md = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		l := md(n.Left)
		r := md(n.Right)
		return 1 + max(l, r)
	}

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	l := md(root.Left)
	r := md(root.Right)
	if abs(l-r) > 1 {
		return false
	}
	return isBalancedv5(root.Left) && isBalancedv5(root.Right)
}

func isBalancedv4(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var maxdepth func(n *TreeNode) int
	maxdepth = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		l := maxdepth(n.Left)
		r := maxdepth(n.Right)
		return 1 + max(l, r)
	}

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	l := maxdepth(root.Left)
	r := maxdepth(root.Right)
	if abs(l-r) > 1 {
		return false
	}
	return isBalancedv4(root.Left) && isBalancedv4(root.Right)
}

func isBalancedv3(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var md func(*TreeNode) int // md: maxDepth
	md = func(tn *TreeNode) int {
		if tn == nil {
			return 0
		}

		l := md(tn.Left)
		r := md(tn.Right)
		return 1 + max(l, r)
	}
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	l := md(root.Left)
	r := md(root.Right)
	if abs(l-r) > 1 {
		return false
	}

	return isBalancedv3(root.Left) && isBalancedv3(root.Right)
}

func isBalancedv2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var depth func(*TreeNode) int
	depth = func(tn *TreeNode) int {
		if tn == nil {
			return 0
		}
		l := depth(tn.Left)
		r := depth(tn.Right)
		return 1 + max(l, r)
	}

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	if abs(depth(root.Left)-depth(root.Right)) > 1 {
		return false
	}

	return isBalancedv2(root.Left) && isBalancedv2(root.Right)
}

func isBalancedv1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	if abs(l-r) > 1 {
		return false
	}

	return isBalancedv1(root.Left) && isBalancedv1(root.Right)
}

func maxDepth(n *TreeNode) int {
	if n == nil {
		return 0
	}

	l := maxDepth(n.Left)
	r := maxDepth(n.Right)

	return 1 + max(l, r)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
