package b75

/*
572. Subtree of Another Tree
Easy
7.1K
404
Companies
Given the roots of two binary trees root and subRoot, return true if there is a subtree of root with the same structure and node values of subRoot and false otherwise.

A subtree of a binary tree tree is a tree that consists of a node in tree and all of this node's descendants. The tree tree could also be considered as a subtree of itself.



Example 1:


Input: root = [3,4,5,1,2], subRoot = [4,1,2]
Output: true
Example 2:


Input: root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
Output: false


Constraints:

The number of nodes in the root tree is in the range [1, 2000].
The number of nodes in the subRoot tree is in the range [1, 1000].
-104 <= root.val <= 104
-104 <= subRoot.val <= 104
*/

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtreev28(root, sub *TreeNode) bool {
	var same func(p, q *TreeNode) bool
	same = func(p, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}

		if p.Val != q.Val {
			return false
		}

		return same(p.Left, q.Left) && same(p.Right, q.Right)
	}

	return same(root, sub) || isSubtreev28(root.Left, sub) || isSubtreev28(root.Right, sub)
}

func isSubtreev27(root, sub *TreeNode) bool {
	var same func(p, q *TreeNode) bool

	same = func(p, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}

		if p.Val != q.Val {
			return false
		}

		return same(p.Left, q.Left) && same(p.Right, q.Right)
	}

	return same(root, sub) || isSubtreev27(root.Left, sub) || isSubtreev27(root.Right, sub)
}

func isSubtreev26(root, sub *TreeNode) bool {
	var same func(p, q *TreeNode) bool

	same = func(p, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}

		if p.Val != q.Val {
			return false
		}

		return same(p.Left, q.Left) && same(p.Right, q.Right)
	}

	return same(root, sub) || isSubtreev26(root.Left, sub) || isSubtreev26(root.Right, sub)
}

func isSubtreev25(root, sub *TreeNode) bool {
	var same func(p, q *TreeNode) bool
	same = func(p, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}

		if p.Val != q.Val {
			return false
		}

		return same(p.Left, q.Left) && same(p.Right, q.Right)
	}

	return same(root, sub) || isSubtreev25(root.Left, sub) || isSubtreev25(root.Right, sub)
}

func isSubtreev24(root, sub *TreeNode) bool {
	var same func(p, q *TreeNode) bool
	same = func(p, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}

		if p.Val != q.Val {
			return false
		}

		return same(p.Left, q.Left) && same(p.Right, q.Right)
	}

	return same(root, sub) || isSubtreev24(root.Left, sub) || isSubtreev24(root.Right, sub)
}

func isSubtreev23(root, sub *TreeNode) bool {
	var same func(p, q *TreeNode) bool
	same = func(p, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}

		if p.Val != q.Val {
			return false
		}

		return same(p.Left, q.Left) && same(p.Right, q.Right)
	}

	return same(root, sub) || isSubtreev23(root.Left, sub) || isSubtreev23(root.Right, sub)
}

func isSubtreev22(root, sub *TreeNode) bool {
	var same func(p, q *TreeNode) bool

	same = func(p, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}

		if p.Val != q.Val {
			return false
		}

		return same(p.Left, q.Left) && same(p.Right, q.Right)
	}

	return same(root, sub) || isSubtreev22(root.Left, sub) || isSubtreev22(root.Right, sub)
}

func isSubtreev21(root, sub *TreeNode) bool {
	var same func(p, q *TreeNode) bool

	same = func(p, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}

		if p.Val != q.Val {
			return false
		}

		return same(p.Left, q.Left) && same(p.Right, q.Right)
	}

	return same(root, sub) || isSubtreev21(root.Left, sub) || isSubtreev21(root.Right, sub)
}

func isSubtreev20(root, sub *TreeNode) bool {
	var same func(p, q *TreeNode) bool
	same = func(p, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}

		if p.Val != q.Val {
			return false
		}

		return same(p.Left, q.Left) && same(p.Right, q.Right)
	}

	return same(root, sub) || isSubtreev20(root.Left, sub) || isSubtreev20(root.Right, sub)
}

func isSubtreev19(root, sub *TreeNode) bool {
	var same func(p, q *TreeNode) bool

	same = func(p, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}

		if p.Val != q.Val {
			return false
		}

		return same(p.Left, q.Left) && same(p.Right, q.Right)
	}

	return same(root, sub) || isSubtreev19(root.Left, sub) || isSubtreev19(root.Right, sub)
}

func isSubtreev18(root, sub *TreeNode) bool {
	var same func(p, q *TreeNode) bool
	same = func(p *TreeNode, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}

		if p.Val != q.Val {
			return false
		}

		return same(p.Left, q.Left) && same(p.Right, q.Right)
	}

	return same(root, sub) || isSubtreev18(root.Right, sub) || isSubtreev18(root.Left, sub)
}

func isSubtreev17(root, sub *TreeNode) bool {
	var same func(p, q *TreeNode) bool
	same = func(p, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}

		if p.Val != q.Val {
			return false
		}

		return same(p.Left, q.Left) && same(p.Right, q.Right)
	}

	return same(root, sub) || isSubtreev17(root.Left, sub) || isSubtreev17(root.Right, sub)
}

func isSubtreev16(root, sub *TreeNode) bool {
	if root == nil || sub == nil {
		return root == sub
	}

	var same func(p, q *TreeNode) bool

	same = func(p, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}

		if p.Val != q.Val {
			return false
		}

		return same(p.Left, q.Left) && same(p.Right, q.Right)
	}

	return same(root, sub) || isSubtreev16(root.Left, sub) || isSubtreev16(root.Right, sub)
}

func isSubtreev15(root, sub *TreeNode) bool {
	var same func(p, q *TreeNode) bool

	same = func(p, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}

		if p.Val != q.Val {
			return false
		}

		return same(p.Left, q.Left) && same(p.Right, q.Right)
	}

	return same(root, sub) || isSubtreev15(root.Left, sub) || isSubtreev15(root.Right, sub)
}

func isSubtreev14(root, sub *TreeNode) bool {
	var same func(p, q *TreeNode) bool

	same = func(p, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}

		if p.Val != q.Val {
			return false
		}

		return same(p.Left, q.Left) && same(p.Right, q.Right)
	}

	return same(root, sub) || isSubtreev14(root.Left, sub) || isSubtreev14(root.Right, sub)
}

func isSubtreev13(root, sub *TreeNode) bool {

	var same func(p, q *TreeNode) bool
	same = func(p, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}

		if p.Val != q.Val {
			return false
		}

		return same(p.Left, q.Left) && same(p.Right, q.Right)
	}

	return same(root, sub) || isSubtreev13(root.Left, sub) || isSubtreev13(root.Right, sub)
}

func isSubtreev12(root, sub *TreeNode) bool {
	if root == nil || sub == nil {
		return root == sub
	}

	var same func(p, q *TreeNode) bool

	same = func(p, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}

		if p.Val != q.Val {
			return false
		}

		return same(p.Left, q.Left) && same(p.Right, q.Right)
	}

	return same(root, sub) || isSubtreev12(root.Left, sub) || isSubtreev12(root.Right, sub)
}

func isSubtreev11(root, sub *TreeNode) bool {
	if root == nil || sub == nil {
		return root == sub
	}

	var same func(p, q *TreeNode) bool
	same = func(p, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}

		if p.Val != q.Val {
			return false
		}

		return same(p.Left, q.Left) && same(p.Right, q.Right)
	}

	return same(root, sub) || isSubtreev11(root.Left, sub) || isSubtreev11(root.Right, sub)
}

func isSubtreev10(root, sub *TreeNode) bool {
	if root == nil || sub == nil {
		return root == sub
	}

	var same func(p, q *TreeNode) bool
	same = func(p, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}

		if p.Val != q.Val {
			return false
		}

		return same(p.Left, q.Left) && same(p.Right, q.Right)
	}

	return same(root, sub) || isSubtreev10(root.Left, sub) || isSubtreev10(root.Right, sub)
}

func isSubtreev9(root, sub *TreeNode) bool {
	if root == nil || sub == nil {
		return root == sub
	}
	var same func(p, q *TreeNode) bool

	same = func(p, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}

		if p.Val != q.Val {
			return false
		}

		return same(p.Left, q.Left) && same(p.Right, q.Right)
	}

	return same(root, sub) || isSubtreev9(root.Left, sub) || isSubtreev9(root.Right, sub)
}

func isSubtreev8(root, sub *TreeNode) bool {
	if root == nil || sub == nil {
		return root == sub
	}

	var same func(p, q *TreeNode) bool
	same = func(p *TreeNode, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}

		if p.Val != q.Val {
			return false
		}

		return same(p.Left, q.Left) && same(p.Right, q.Right)
	}

	return same(root, sub) || isSubtreev8(root.Left, sub) || isSubtreev8(root.Right, sub)
}

func isSubtreev7(root, sub *TreeNode) bool {
	var same func(p, q *TreeNode) bool

	same = func(p, q *TreeNode) bool {
		if p == q {
			return p == q
		}
		if p.Val != q.Val {
			return false
		}

		return same(p.Left, q.Left) && same(p.Right, q.Right)
	}

	return same(root, sub) || isSubtreev7(root.Left, sub) || isSubtreev7(root.Right, sub)
}

func isSubtreev6(root, sub *TreeNode) bool {

	var same func(p, q *TreeNode) bool
	same = func(p, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}

		if p.Val != q.Val {
			return false
		}
		return same(p.Left, q.Left) && same(p.Right, q.Right)
	}

	return same(root, sub) || isSubtreev6(root.Left, sub) || isSubtreev6(root.Right, sub)
}

func isSubtreev5(root, sub *TreeNode) bool {
	if root == nil || sub == nil {
		return root == sub
	}

	var same func(p, q *TreeNode) bool
	same = func(p, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}

		if p.Val != q.Val {
			return false
		}

		return same(p.Left, q.Left) && same(p.Right, q.Right)
	}

	return same(root, sub) || same(root.Left, sub) || same(root.Right, sub)
}

func isSubtreev4(root, subRoot *TreeNode) bool {

	if root == nil || subRoot == nil {
		return root == subRoot
	}

	var same func(p, q *TreeNode) bool
	same = func(p, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}
		if p.Val != q.Val {
			return false
		}
		return same(p.Left, q.Left) && same(p.Right, q.Right)
	}

	return same(root, subRoot) || same(root.Left, subRoot) || same(root.Right, subRoot)
}

func isSubtreev3(root, subRoot *TreeNode) bool {

	var same func(p, q *TreeNode) bool
	same = func(p, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}

		if p.Val != q.Val {
			return false
		}
		return same(p.Left, q.Left) && same(p.Right, q.Right)
	}

	if same(root, subRoot) {
		return true
	}

	return isSubtreev3(root.Left, subRoot) || isSubtreev3(root.Right, subRoot)
}

func isSubtreev2(root, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return root == subRoot
	}

	var same func(p, q *TreeNode) bool
	same = func(p, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}

		if p.Val != q.Val {
			return false
		}

		return same(p.Left, q.Left) && same(p.Right, q.Right)
	}

	return same(root, subRoot) || isSubtreev2(root.Left, subRoot) || isSubtreev2(root.Right, subRoot)
}

func isSubtreev1(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return root == subRoot
	}

	var same func(p, q *TreeNode) bool
	same = func(p, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}

		if p.Val != q.Val {
			return false
		}
		return same(p.Left, q.Left) && same(p.Right, q.Right)
	}

	if same(root, subRoot) {
		return true
	}

	return isSubtreev1(root.Left, subRoot) || isSubtreev1(root.Right, subRoot)
}
