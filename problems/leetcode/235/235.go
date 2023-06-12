package b75

/*
235. Lowest Common Ancestor of a Binary Search Tree
Medium
9.2K
261
Companies
Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”



Example 1:


Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
Output: 6
Explanation: The LCA of nodes 2 and 8 is 6.
Example 2:


Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
Output: 2
Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself according to the LCA definition.
Example 3:

Input: root = [2,1], p = 2, q = 1
Output: 2


Constraints:

The number of nodes in the tree is in the range [2, 105].
-109 <= Node.val <= 109
All Node.val are unique.
p != q
p and q will exist in the BST.
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

func lowestCommonAncestorv23(root, p, q *TreeNode) *TreeNode {

	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestorv23(root.Right, p, q)
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestorv23(root.Left, p, q)
	}

	return root
}

func lowestCommonAncestorv22(root, p, q *TreeNode) *TreeNode {

	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestorv22(root.Right, p, q)
	}

	if q.Val < root.Val && p.Val < root.Val {
		return lowestCommonAncestorv22(root.Left, p, q)
	}

	return root
}

func lowestCommonAncestorv21(root, p, q *TreeNode) *TreeNode {

	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestorv21(root.Right, p, q)
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestorv21(root.Left, p, q)
	}

	return root
}

func lowestCommonAncestorv20(root, p, q *TreeNode) *TreeNode {

	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestorv20(root.Right, p, q)
	}

	if q.Val < root.Val && p.Val < root.Val {
		return lowestCommonAncestorv20(root.Left, p, q)
	}

	return root
}

func lowestCommonAncestorv19(root, p, q *TreeNode) *TreeNode {

	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestorv19(root.Right, p, q)
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestorv19(root.Left, p, q)
	}

	return root
}

func lowestCommonAncestorv18(root, p, q *TreeNode) *TreeNode {

	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestorv18(root.Right, p, q)
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestorv18(root.Left, p, q)
	}

	return root
}

func lowestCommonAncestorv17(root, p, q *TreeNode) *TreeNode {

	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestorv17(root.Right, p, q)
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestorv17(root.Left, p, q)
	}

	return root
}

func lowestCommonAncestorv16(root, p, q *TreeNode) *TreeNode {

	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestorv16(root.Right, p, q)
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestorv16(root.Left, p, q)
	}

	return root
}

func lowestCommonAncestorv15(root, p, q *TreeNode) *TreeNode {

	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestorv15(root.Right, p, q)
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestorv15(root.Left, p, q)
	}

	return root
}

func lowestCommonAncestorv14(root, p, q *TreeNode) *TreeNode {

	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestorv14(root.Left, p, q)
	}

	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestorv14(root.Right, p, q)
	}

	return root
}

func lowestCommonAncestorv13(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestorv13(root.Right, p, q)
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestorv13(root.Left, p, q)
	}

	return root
}

func lowestCommonAncestorv12(root, p, q *TreeNode) *TreeNode {

	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestorv12(root.Right, p, q)
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestorv12(root.Left, p, q)
	}

	return root
}

func lowestCommonAncestorv11(root, p, q *TreeNode) *TreeNode {

	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestorv11(root.Right, p, q)
	}

	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestorv11(root.Left, p, q)
	}

	return root
}

func lowestCommonAncestorv10(root, p, q *TreeNode) *TreeNode {

	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestorv10(root.Right, p, q)
	}

	if p.Val < root.Val && q.Val < root.Left.Val {
		return lowestCommonAncestorv10(root.Left, p, q)
	}

	return root
}

func lowestCommonAncestorv9(root, p, q *TreeNode) *TreeNode {

	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestorv9(root.Right, p, q)
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestorv9(root.Left, p, q)
	}

	return root
}

func lowestCommonAncestorv8(root, p, q *TreeNode) *TreeNode {

	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestorv8(root.Right, p, q)
	}

	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestorv8(root.Left, p, q)
	}

	return root
}

func lowestCommonAncestorv7(root, p, q *TreeNode) *TreeNode {

	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestorv7(root.Right, p, q)
	}

	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestorv7(root.Left, p, q)
	}

	return root
}

func lowestCommonAncestorv6(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestorv6(root.Right, p, q)
	}

	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestorv6(root.Left, p, q)
	}

	return root
}

func lowestCommonAncestorv5(root, p, q *TreeNode) *TreeNode {

	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestorv5(root.Right, p, q)
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestorv5(root.Left, p, q)
	}

	return root
}

func lowestCommonAncestorv4(root, p, q *TreeNode) *TreeNode {

	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestorv4(root.Right, p, q)
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestorv4(root.Left, p, q)
	}

	return root
}

func lowestCommonAncestorv3(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestorv3(root.Right, p, q)
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestorv3(root.Left, p, q)
	}

	return root
}

func lowestCommonAncestorv2(root, p, q *TreeNode) *TreeNode {

	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestorv2(root.Right, p, q)
	}

	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestorv2(root.Left, p, q)
	}

	return root
}

func lowestCommonAncestorv1(root, p, q *TreeNode) *TreeNode {

	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	return root
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if (root.Left == p && root.Right == q) || (root.Left == q && root.Left == p) {
		return root
	}

	if p.Left == q || p.Right == q {
		return p
	}

	if q.Left == p || q.Right == p {
		return q
	}

	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	return root
}
