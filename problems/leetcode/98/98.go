package b75

/*
98. Validate Binary Search Tree
Medium
14.5K
1.2K
Companies
Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

The left
subtree
 of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.


Example 1:


Input: root = [2,1,3]
Output: true
Example 2:


Input: root = [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.


Constraints:

The number of nodes in the tree is in the range [1, 104].
-231 <= Node.val <= 231 - 1
Accepted
1.9M
Submissions
5.9M
Acceptance Rate
32.0%
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

func isValidBSTv28(root *TreeNode) bool {
	var valid func(n, min, max *TreeNode) bool

	valid = func(n, min, max *TreeNode) bool {
		if n == nil {
			return true
		}

		if min != nil && min.Val >= n.Val {
			return false
		}

		if max != nil && max.Val <= n.Val {
			return false
		}

		return valid(n.Left, min, n) && valid(n.Right, n, max)
	}

	return valid(root, nil, nil)
}

func isValidBSTv27(root *TreeNode) bool {
	var valid func(n, min, max *TreeNode) bool

	valid = func(n, min, max *TreeNode) bool {
		if n == nil {
			return true
		}

		if min != nil && min.Val >= n.Val {
			return false
		}

		if max != nil && max.Val <= n.Val {
			return false
		}

		return valid(n.Left, min, n) && valid(n.Right, n, max)
	}

	return valid(root, nil, nil)
}

func isValidBSTv26(root *TreeNode) bool {
	var valid func(n, min, max *TreeNode) bool

	valid = func(n, min, max *TreeNode) bool {
		if n == nil {
			return true
		}

		if min != nil && min.Val >= n.Val {
			return false
		}

		if max != nil && max.Val <= n.Val {
			return false
		}

		return valid(n.Left, min, n) && valid(n.Right, n, max)
	}

	return valid(root, nil, nil)
}

func isValidBSTv25(root *TreeNode) bool {
	var valid func(n, min, max *TreeNode) bool

	valid = func(n, min, max *TreeNode) bool {
		if n == nil {
			return true
		}

		if min != nil && min.Val >= n.Val {
			return false
		}

		if max != nil && max.Val <= n.Val {
			return false
		}

		return valid(n.Left, min, n) && valid(n.Right, n, max)
	}

	return valid(root, nil, nil)
}

func isValidBSTv24(root *TreeNode) bool {
	var valid func(n, min, max *TreeNode) bool

	valid = func(n, min, max *TreeNode) bool {
		if n == nil {
			return true
		}

		if min != nil && min.Val >= n.Val {
			return false
		}

		if max != nil && max.Val <= n.Val {
			return false
		}

		return valid(n.Left, min, n) && valid(n.Right, n, max)
	}

	return valid(root, nil, nil)
}

func isValidBSTv23(root *TreeNode) bool {
	var valid func(n, min, max *TreeNode) bool

	valid = func(n, min, max *TreeNode) bool {
		if n == nil {
			return true
		}

		if min != nil && min.Val >= n.Val {
			return false
		}

		if max != nil && max.Val <= n.Val {
			return false
		}

		return valid(n.Left, min, n) && valid(n.Right, n, max)
	}

	return valid(root, nil, nil)
}

func isValidBSTv22(root *TreeNode) bool {
	var valid func(n, min, max *TreeNode) bool

	valid = func(n, min, max *TreeNode) bool {
		if n == nil {
			return true
		}

		if min != nil && min.Val >= n.Val {
			return false
		}

		if max != nil && max.Val <= n.Val {
			return false
		}

		return valid(n.Left, min, n) && valid(n.Right, n, max)
	}

	return valid(root, nil, nil)
}

func isValidBSTv21(root *TreeNode) bool {
	var valid func(n, min, max *TreeNode) bool

	valid = func(n, min, max *TreeNode) bool {
		if n == nil {
			return true
		}

		if min != nil && min.Val >= n.Val {
			return false
		}

		if max != nil && max.Val <= n.Val {
			return false
		}

		return valid(n.Left, min, n) && valid(n.Right, n, max)
	}

	return valid(root, nil, nil)
}

func isValidBSTv20(root *TreeNode) bool {
	var valid func(n, min, max *TreeNode) bool

	valid = func(n, min, max *TreeNode) bool {
		if n == nil {
			return true
		}

		if min != nil && min.Val >= n.Val {
			return false
		}

		if max != nil && max.Val <= n.Val {
			return false
		}

		return valid(n.Left, min, n) && valid(n.Right, n, max)
	}

	return valid(root, nil, nil)
}

func isValidBSTv19(root *TreeNode) bool {

	var valid func(n, min, max *TreeNode) bool

	valid = func(n, min, max *TreeNode) bool {
		if n == nil {
			return true
		}

		if min != nil && min.Val >= n.Val {
			return false
		}

		if max != nil && max.Val <= n.Val {
			return false
		}

		return valid(n.Left, min, n) && valid(n.Right, n, max)
	}
	return valid(root, nil, nil)
}

func isValidBSTv18(root *TreeNode) bool {
	var valid func(n, min, max *TreeNode) bool

	valid = func(n, min, max *TreeNode) bool {
		if n == nil {
			return true
		}

		if min != nil && min.Val >= n.Val {
			return false
		}

		if max != nil && max.Val <= n.Val {
			return false
		}

		return valid(n.Left, min, n) && valid(n.Right, n, max)
	}

	return valid(root, nil, nil)
}

func isValidBSTv17(root *TreeNode) bool {

	var valid func(n, min, max *TreeNode) bool
	valid = func(n, min, max *TreeNode) bool {
		if n == nil {
			return true
		}

		if min != nil && min.Val >= n.Val {
			return false
		}

		if max != nil && max.Val <= n.Val {
			return false
		}

		return valid(n.Left, min, n) && valid(n.Right, n, max)
	}

	return valid(root, nil, nil)
}

func isValidBSTv16(root *TreeNode) bool {

	var valid func(n, min, max *TreeNode) bool

	valid = func(n, min, max *TreeNode) bool {
		if n == nil {
			return true
		}

		if min != nil && min.Val >= n.Val {
			return false
		}

		if max != nil && max.Val <= n.Val {
			return false
		}

		return valid(n.Left, min, n) && valid(n.Right, n, max)
	}

	return valid(root, nil, nil)
}

func isValidBSTv15(root *TreeNode) bool {

	var valid func(n, min, max *TreeNode) bool

	valid = func(n, min, max *TreeNode) bool {
		if n == nil {
			return true
		}

		if min != nil && min.Val >= n.Val {
			return false
		}

		if max != nil && max.Val <= n.Val {
			return false
		}

		return valid(n.Left, min, n) && valid(n.Right, n, max)
	}

	return valid(root, nil, nil)
}

func isValidBSTv14(root *TreeNode) bool {

	var valid func(n, max, min *TreeNode) bool

	valid = func(n, max, min *TreeNode) bool {
		if n == nil {
			return true
		}

		if min != nil && min.Val >= n.Val {
			return false
		}

		if max != nil && max.Val <= n.Val {
			return false
		}

		return valid(n.Left, n, min) && valid(n.Right, max, n)
	}

	return valid(root, nil, nil)
}

func isValidBSTv13(root *TreeNode) bool {
	var valid func(n, min, max *TreeNode) bool

	valid = func(n, min, max *TreeNode) bool {
		if n == nil {
			return true
		}

		if min != nil && min.Val >= n.Val {
			return false
		}

		if max != nil && max.Val <= n.Val {
			return false
		}

		return valid(n.Left, min, n) && valid(n.Right, n, max)
	}

	return valid(root, nil, nil)
}

func isValidBSTv12(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var valid func(n, min, max *TreeNode) bool
	valid = func(n, min, max *TreeNode) bool {
		if n == nil {
			return true
		}

		if min != nil && min.Val >= n.Val {
			return false
		}

		if max != nil && max.Val <= n.Val {
			return false
		}

		return valid(n.Left, min, n) && valid(n.Right, n, max)
	}

	return valid(root, nil, nil)
}

func isValidBSTv11(root *TreeNode) bool {

	var valid func(n, min, max *TreeNode) bool

	valid = func(n, min, max *TreeNode) bool {
		if n == nil {
			return true
		}

		if min != nil && min.Val >= n.Val {
			return false
		}

		if max != nil && max.Val <= n.Val {
			return false
		}

		return valid(n.Left, min, n) && valid(n.Right, n, max)
	}

	return valid(root, nil, nil)
}

func isValidBSTv10(root *TreeNode) bool {

	var valid func(n, min, max *TreeNode) bool

	valid = func(n, min, max *TreeNode) bool {
		if n == nil {
			return true
		}

		if min != nil && min.Val >= n.Val {
			return false
		}

		if max != nil && max.Val <= n.Val {
			return false
		}

		return valid(n.Left, min, n) && valid(n.Right, n, max)
	}

	return valid(root, nil, nil)
}

func isValidBSTv9(root *TreeNode) bool {
	var valid func(n, min, max *TreeNode) bool

	valid = func(n, min, max *TreeNode) bool {
		if n == nil {
			return true
		}

		if min != nil && min.Val >= n.Val {
			return false
		}

		if max != nil && max.Val <= n.Val {
			return false
		}

		return valid(n.Left, min, n) && valid(n.Right, n, max)
	}

	return valid(root, nil, nil)
}

func isValidBSTv8(root *TreeNode) bool {
	var valid func(n, min, max *TreeNode) bool

	valid = func(n, min, max *TreeNode) bool {
		if n == nil {
			return true
		}

		if min != nil && min.Val >= n.Val {
			return false
		}

		if max != nil && max.Val <= n.Val {
			return false
		}

		return valid(n.Left, min, n) && valid(n.Right, n, max)
	}

	return valid(root, nil, nil)
}

func isValidBSTv7(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var valid func(n, min, max *TreeNode) bool
	valid = func(n, min, max *TreeNode) bool {
		if n == nil {
			return true
		}

		if min != nil && min.Val >= n.Val {
			return false
		}

		if max != nil && max.Val <= n.Val {
			return false
		}

		return valid(n.Left, min, n) && valid(n.Right, n, max)
	}

	return valid(root, nil, nil)
}

func isValidBSTv6(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var valid func(n, min, max *TreeNode) bool
	valid = func(n, min, max *TreeNode) bool {
		if n == nil {
			return true
		}

		if min != nil && min.Val >= n.Val {
			return false
		}

		if max != nil && max.Val <= n.Val {
			return false
		}

		return valid(n.Left, min, n) && valid(n.Right, n, max)
	}

	return valid(root, nil, nil)
}

func isValidBSTv5(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var valid func(n, min, max *TreeNode) bool

	valid = func(n, min, max *TreeNode) bool {
		if n == nil {
			return true
		}

		if min != nil && min.Val >= n.Val {
			return false
		}

		if max != nil && max.Val <= n.Val {
			return false
		}

		return valid(n.Left, min, n) && valid(n.Right, n, max)
	}

	return valid(root, nil, nil)
}

func isValidBSTv4(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var valid func(n, min, max *TreeNode) bool

	valid = func(n, min, max *TreeNode) bool {
		if n == nil {
			return true
		}

		if min != nil && min.Val >= n.Val {
			return false
		}

		if max != nil && max.Val <= n.Val {
			return false
		}

		return valid(n.Left, min, n) && valid(n.Right, n, max)
	}

	return valid(root, nil, nil)
}

func isValidBSTv3(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var valid func(n, min, max *TreeNode) bool
	valid = func(n, min, max *TreeNode) bool {
		if n == nil {
			return true
		}

		if min != nil && min.Val >= n.Val {
			return false
		}

		if max != nil && max.Val <= n.Val {
			return false
		}

		return valid(n.Left, min, n) && valid(n.Right, n, max)
	}

	return valid(root, nil, nil)
}

func isValidBSTv2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var isValid func(n, min, max *TreeNode) bool
	isValid = func(n, min, max *TreeNode) bool {
		if n == nil {
			return true
		}

		if min != nil && min.Val >= n.Val {
			return false
		}

		if max != nil && max.Val <= n.Val {
			return false
		}

		return isValid(n.Left, min, n) && isValid(n.Right, n, max)
	}

	return isValid(root, nil, nil)
}

func isValidBSTv1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var isValid func(n, min, max *TreeNode) bool
	isValid = func(n, min, max *TreeNode) bool {
		if n == nil {
			return true
		}

		if min != nil && n.Val <= min.Val {
			return false
		}

		if max != nil && n.Val >= max.Val {
			return false
		}

		return isValid(n.Left, min, n) && isValid(n.Right, n, max)
	}

	return isValid(root, nil, nil)
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var vaild func(n, min, max *TreeNode) bool

	vaild = func(n, min, max *TreeNode) bool {
		if n == nil {
			return true
		}

		if min != nil && n.Val <= min.Val {
			return false
		}

		if max != nil && n.Val >= max.Val {
			return false
		}

		return vaild(n.Left, min, n) && vaild(n.Right, n, max)
	}

	return vaild(root, nil, nil)
}
