package b75

/*

230. Kth Smallest Element in a BST
Medium
9.6K
173
Companies
Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of all the values of the nodes in the tree.



Example 1:


Input: root = [3,1,4,null,2], k = 1
Output: 1
Example 2:


Input: root = [5,3,6,2,4,null,null,1], k = 3
Output: 3


Constraints:

The number of nodes in the tree is n.
1 <= k <= n <= 104
0 <= Node.val <= 104


Follow up: If the BST is modified often (i.e., we can do insert and delete operations) and you need to find the kth smallest frequently, how would you optimize?

Accepted
1.1M
Submissions
1.5M
Acceptance Rate
70.3%

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

func kthSmallestv25(root *TreeNode, k int) int {
	stack := []*TreeNode{}

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		k--
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if k == 0 {
			return n.Val
		}

		root = n.Right
	}

	return -1
}

func kthSmallestv24(root *TreeNode, k int) int {
	stack := []*TreeNode{}

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		k--
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if k == 0 {
			return n.Val
		}

		root = n.Right
	}

	return -1
}

func kthSmallestv23(root *TreeNode, k int) int {
	stack := []*TreeNode{}

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		k--
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if k == 0 {
			return n.Val
		}

		root = n.Right
	}

	return -1
}

func kthSmallestv22(root *TreeNode, k int) int {
	stack := []*TreeNode{}

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		k--
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if k == 0 {
			return n.Val
		}

		root = n.Right
	}

	return -1
}

func kthSmallestv21(root *TreeNode, k int) int {
	stack := []*TreeNode{}

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		k--
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if k == 0 {
			return n.Val
		}

		root = n.Right
	}

	return -1
}

func kthSmallestv20(root *TreeNode, k int) int {
	stack := []*TreeNode{root}

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		k--
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if k == 0 {
			return n.Val
		}

		root = n.Right

	}

	return -1
}

func kthSmallestv19(root *TreeNode, k int) int {
	stack := []*TreeNode{root}

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return n.Val
		}

		root = n.Right
	}

	return -1
}

func kthSmallestv18(root *TreeNode, k int) int {
	stack := []*TreeNode{}

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return n.Val
		}

		root = n.Right
	}

	return -1
}

func kthSmallestv17(root *TreeNode, k int) int {
	stack := []*TreeNode{}

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return n.Val
		}

		root = n.Right
	}

	return -1
}

func kthSmallestv16(root *TreeNode, k int) int {
	stack := []*TreeNode{root}

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return n.Val
		}

		root = n.Right
	}

	return -1
}

func kthSmallestv15(root *TreeNode, k int) int {
	stack := []*TreeNode{}

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return n.Val
		}

		root = n.Right
	}

	return -1
}

func kthSmallestv14(root *TreeNode, k int) int {
	stack := []*TreeNode{}

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return n.Val
		}

		root = n.Right
	}

	return -1
}

func kthSmallestv13(root *TreeNode, k int) int {
	stack := []*TreeNode{}

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return n.Val
		}

		root = n.Right
	}

	return -1
}

func kthSmallestv12(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	stack := []*TreeNode{}

	for len(stack) > 0 || root != nil {

		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return n.Val
		}

		root = n.Right
	}

	return -1
}

func kthSmallestv11(root *TreeNode, k int) int {
	stack := []*TreeNode{}

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return n.Val
		}

		root = n.Right
	}

	return -1
}

func kthSmallestv10(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	stack := []*TreeNode{}

	// inorder
	for root != nil || len(stack) > 0 {

		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return n.Val
		}
		root = n.Right
	}

	return -1
}

func kthSmallestv9(root *TreeNode, k int) int {
	stack := []*TreeNode{}

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return n.Val
		}
		root = n.Right
	}

	return -1
}

func kthSmallestv8(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	stack := []*TreeNode{}
	for len(stack) > 0 || root != nil {

		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return n.Val
		}
		root = n.Right
	}
	return -1
}

func kthSmallestv7(root *TreeNode, k int) int {

	if root == nil {
		return -1
	}

	stack := []*TreeNode{}

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return n.Val
		}

		root = n.Right
	}

	return -1
}

func kthSmallestv6(root *TreeNode, k int) int {

	// inorder -> minimal first...
	// stack save the previous node
	// if left is done, next is the right, final is root
	// because it's the valid bst
	stack := []*TreeNode{}

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		// the current minimal
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return n.Val
		}

		// switch to the right
		root = n.Right
	}

	return -1
}

func kthSmallestv5(root *TreeNode, k int) int {
	// use dfs to reach the smallest node
	// back to the k nodes to satisfy the kth small node
	// HINT inorder traversal

	stack := make([]*TreeNode, 0)

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return n.Val
		}

		root = n.Right
	}

	return -1
}

func kthSmallestv4(root *TreeNode, k int) int {

	stack := []*TreeNode{root}

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return n.Val
		}

		root = n.Right
	}

	return -1
}

func kthSmallestv3(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0)

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return n.Val
		}
		root = n.Right
	}

	return -1
}

func kthSmallestv2(root *TreeNode, k int) int {

	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--

		if k == 0 {
			return n.Val
		}
		root = n.Right
	}

	return -1
}

func kthSmallestv1(root *TreeNode, k int) int {
	// inorder
	stack := make([]*TreeNode, 0)
	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		// pop the previous one
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
}

func kthSmallest(root *TreeNode, k int) int {

	n := 0
	stack := make([]*TreeNode, 0)
	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		n++
		if n == k {
			return root.Val
		}
		root = root.Right
	}
}

func kthSmallestRec(root *TreeNode, k int) int {

	result := 0
	cnt := 0

	var inorder func(n *TreeNode)
	inorder = func(n *TreeNode) {
		if n == nil {
			return
		}

		inorder(n.Left)
		cnt++
		if cnt == k {
			result = n.Val
		}
		inorder(n.Right)
	}
	inorder(root)

	return result
}
