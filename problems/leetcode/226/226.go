package b75

/*
226. Invert Binary Tree
Easy
12K
169
Companies
Given the root of a binary tree, invert the tree, and return its root.



Example 1:


Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]
Example 2:


Input: root = [2,1,3]
Output: [2,3,1]
Example 3:

Input: root = []
Output: []


Constraints:

The number of nodes in the tree is in the range [0, 100].
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

func invertTreev30(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTreev30(root.Left)
	invertTreev30(root.Right)
	return root
}

func invertTreev29(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTreev29(root.Left)
	invertTreev29(root.Right)
	return root
}

func invertTreev28(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = root.Right, root.Left
	invertTreev28(root.Left)
	invertTreev28(root.Right)

	return root
}

func invertTreev27(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left

	invertTreev27(root.Left)
	invertTreev27(root.Right)

	return root
}

func invertTreev26(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left

	invertTreev26(root.Left)
	invertTreev26(root.Right)
	return root
}

func invertTreev25(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTreev25(root.Left)
	invertTreev25(root.Right)

	return root
}

func invertTreev24(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTreev24(root.Left)
	invertTreev24(root.Right)

	return root
}
func invertTreev23(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTreev23(root.Left)
	invertTreev23(root.Right)

	return root
}

func invertTreev22(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTreev22(root.Left)
	invertTreev22(root.Right)

	return root
}

func invertTreev21(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left

	invertTreev21(root.Left)
	invertTreev21(root.Right)

	return root
}

func invertTreev20(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTreev20(root.Left)
	invertTreev20(root.Right)
	return root
}

func invertTreev19(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTreev19(root.Left)
	invertTreev19(root.Right)
	return root
}

func invertTreev18(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTreev18(root.Left)
	invertTreev18(root.Right)

	return root
}

func invertTreev17(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTreev17(root.Left)
	invertTreev17(root.Right)
	return root
}

func invertTreev16(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTreev16(root.Left)
	invertTreev16(root.Right)
	return root
}

func invertTreev15(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTreev15(root.Left)
	invertTreev15(root.Right)
	return root
}

func invertTreev14(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTreev14(root.Left)
	invertTreev14(root.Right)
	return root
}

func invertTreev13(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var invert func(n *TreeNode)
	invert = func(n *TreeNode) {
		if n == nil {
			return
		}

		n.Left, n.Right = n.Right, n.Left
		invert(n.Left)
		invert(n.Right)
	}

	invert(root)

	return root
}

func invertTreev12(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTreev12(root.Left)
	invertTreev12(root.Right)
	return root
}

func invertTreev11(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTreev11(root.Left)
	invertTreev11(root.Right)

	return root
}

func invertTreev10(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTreev10(root.Left)
	invertTreev10(root.Right)
	return root
}

func invertTreev9(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTreev9(root.Left)
	invertTreev9(root.Right)
	return root
}

func invertTreev8(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTreev8(root.Left)
	invertTreev8(root.Right)

	return root
}

func invertTreev7(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTreev7(root.Left)
	invertTreev7(root.Right)
	return root
}

func invertTreev6(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTreev6(root.Left)
	invertTreev6(root.Right)
	return root
}

func invertTreev5(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTreev5(root.Left)
	invertTreev5(root.Right)
	return root
}

func invertTreev4(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left

	invertTreev4(root.Left)
	invertTreev4(root.Right)
	return root
}

func invertTreev3(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	left := invertTreev3(root.Left)
	root.Left = invertTreev3(root.Right)
	root.Right = left

	return root
}

func invertTreev2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	left := invertTreev2(root.Left)
	root.Left = invertTreev2(root.Right)
	root.Right = left

	return root
}

func invertTreev1(root *TreeNode) *TreeNode {

	if root == nil {
		return root
	}

	left := invertTreev1(root.Left)
	root.Left = invertTreev1(root.Right)
	root.Right = left

	return root
}
