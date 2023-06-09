package b75

/*
105. Construct Binary Tree from Preorder and Inorder Traversal
Medium
12.7K
371
Companies
Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.



Example 1:


Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
Example 2:

Input: preorder = [-1], inorder = [-1]
Output: [-1]


Constraints:

1 <= preorder.length <= 3000
inorder.length == preorder.length
-3000 <= preorder[i], inorder[i] <= 3000
preorder and inorder consist of unique values.
Each value of inorder also appears in preorder.
preorder is guaranteed to be the preorder traversal of the tree.
inorder is guaranteed to be the inorder traversal of the tree.
Accepted
958.7K
Submissions
1.6M
Acceptance Rate
61.7%
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

func buildTreev30(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	findMid := func(in []int, t int) int {
		for i, v := range in {
			if v == t {
				return i
			}
		}
		return -1
	}
	mid := findMid(inorder, root.Val)
	root.Left = buildTreev30(preorder[1:mid+1], inorder[:mid+1])
	root.Right = buildTreev30(preorder[mid+1:], inorder[mid+1:])
	return root
}

func buildTreev29(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	findMid := func(in []int, t int) int {
		for i, v := range in {
			if v == t {
				return i
			}
		}
		return -1
	}
	mid := findMid(inorder, root.Val)
	root.Left = buildTreev29(preorder[1:mid+1], inorder[:mid+1])
	root.Right = buildTreev29(preorder[mid+1:], inorder[mid+1:])

	return root
}

func buildTreev28(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	findMid := func(in []int, t int) int {
		for i, v := range in {
			if v == t {
				return i
			}
		}
		return -1
	}
	mid := findMid(inorder, root.Val)
	root.Left = buildTreev28(preorder[1:mid+1], inorder[:mid+1])
	root.Right = buildTreev28(preorder[mid+1:], inorder[mid+1:])

	return root
}

func buildTreev27(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	findMid := func(in []int, t int) int {
		for i, v := range in {
			if v == t {
				return i
			}
		}
		return -1
	}
	mid := findMid(inorder, root.Val)
	root.Left = buildTreev27(preorder[1:mid+1], inorder[:mid+1])
	root.Right = buildTreev27(preorder[mid+1:], inorder[mid+1:])

	return root
}

func buildTreev26(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	findMid := func(in []int, t int) int {
		for i, v := range in {
			if v == t {
				return i
			}
		}
		return -1
	}
	mid := findMid(inorder, root.Val)
	root.Left = buildTreev26(preorder[1:mid+1], inorder[:mid+1])
	root.Right = buildTreev26(preorder[mid+1:], inorder[mid+1:])
	return root
}

func buildTreev25(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	findMid := func(in []int, t int) int {
		for i, v := range in {
			if v == t {
				return i
			}
		}
		return -1
	}
	mid := findMid(inorder, root.Val)
	root.Left = buildTreev25(preorder[1:mid+1], inorder[:mid+1])
	root.Right = buildTreev25(preorder[mid+1:], inorder[mid+1:])

	return root
}

func buildTreev24(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	findMid := func(in []int, t int) int {
		for i, v := range in {
			if v == t {
				return i
			}
		}
		return -1
	}
	mid := findMid(inorder, root.Val)
	root.Left = buildTreev24(preorder[1:mid+1], inorder[:mid+1])
	root.Right = buildTreev24(preorder[mid+1:], inorder[mid+1:])
	return root
}

func buildTreev23(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	findMid := func(in []int, t int) int {
		for i, v := range in {
			if v == t {
				return i
			}
		}
		return -1
	}
	mid := findMid(inorder, root.Val)
	root.Left = buildTreev23(preorder[1:mid+1], inorder[:mid+1])
	root.Right = buildTreev23(preorder[mid+1:], inorder[mid+1:])
	return root
}

func buildTreev22(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}

	findMid := func(in []int, t int) int {

		for i, v := range in {
			if v == t {
				return i
			}
		}
		return -1
	}

	mid := findMid(inorder, root.Val)
	root.Left = buildTreev22(preorder[1:mid+1], inorder[:mid+1])
	root.Right = buildTreev22(preorder[mid+1:], inorder[mid+1:])

	return root
}

func buildTreev21(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	findMid := func(in []int, t int) int { // t: target

		for i, v := range in {
			if v == t {
				return i
			}
		}

		return -1
	}
	mid := findMid(inorder, root.Val)

	root.Left = buildTreev21(preorder[1:mid+1], inorder[:mid+1])
	root.Right = buildTreev21(preorder[mid+1:], inorder[mid+1:])

	return root
}

func buildTreev20(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	findMid := func(in []int, target int) int {

		for i, v := range in {
			if v == target {
				return i
			}
		}
		return -1
	}
	mid := findMid(inorder, root.Val)

	root.Left = buildTreev20(preorder[1:mid+1], inorder[:mid+1])
	root.Right = buildTreev20(preorder[mid+1:], inorder[mid+1:])

	return root
}

func buildTreev19(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	findMid := func(in []int, target int) int {
		for i, v := range in {
			if v == target {
				return i
			}
		}
		return -1
	}
	mid := findMid(inorder, root.Val)

	root.Left = buildTreev19(preorder[1:mid+1], inorder[:mid+1])
	root.Right = buildTreev19(preorder[mid+1:], inorder[mid+1:])

	return root
}

func buildTreev18(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	findMid := func(in []int, target int) int {
		for i, v := range in {
			if v == target {
				return i
			}
		}
		return -1
	}

	mid := findMid(inorder, root.Val)
	root.Left = buildTreev18(preorder[1:mid+1], inorder[:mid+1])
	root.Right = buildTreev18(preorder[mid+1:], inorder[mid+1:])
	return root
}

func buildTreev17(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	findMid := func(in []int, target int) int {
		for i, v := range in {
			if v == target {
				return i
			}
		}
		return -1
	}
	mid := findMid(inorder, root.Val)

	root.Left = buildTreev17(preorder[1:mid+1], inorder[:mid+1])
	root.Right = buildTreev17(preorder[mid+1:], inorder[mid+1:])

	return root
}

func buildTreev16(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}

	findMid := func(in []int, target int) int {
		for i, v := range in {
			if v == target {
				return i
			}
		}
		return -1
	}
	mid := findMid(inorder, root.Val)

	root.Left = buildTreev16(preorder[1:mid+1], inorder[:mid+1])
	root.Right = buildTreev16(preorder[mid+1:], inorder[mid+1:])

	return root
}

func buildTreev15(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	findMid := func(in []int, target int) int {
		for i, v := range in {
			if v == target {
				return i
			}
		}
		return -1
	}
	mid := findMid(inorder, root.Val)

	root.Left = buildTreev15(preorder[1:mid+1], inorder[:mid+1])
	root.Right = buildTreev15(preorder[mid+1:], inorder[mid+1:])

	return root
}

func buildTreev14(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	findMid := func(in []int, target int) int {
		for i, v := range in {
			if v == target {
				return i
			}
		}
		return -1
	}
	mid := findMid(inorder, root.Val)

	root.Left = buildTreev14(preorder[1:mid+1], inorder[:mid+1])
	root.Right = buildTreev14(preorder[mid+1:], inorder[mid+1:])

	return root
}

func buildTreev13(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	findMid := func(arr []int, target int) int {
		for i, v := range arr {
			if v == target {
				return i
			}
		}
		return -1
	}
	mid := findMid(inorder, root.Val)

	root.Left = buildTreev13(preorder[1:mid+1], inorder[:mid+1])
	root.Right = buildTreev13(preorder[mid+1:], inorder[mid+1:])

	return root
}

func buildTreev12(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	findMid := func(arr []int, target int) int {
		for i, v := range arr {
			if v == target {
				return i
			}
		}
		return -1
	}
	mid := findMid(inorder, root.Val)

	root.Left = buildTreev12(preorder[1:mid+1], inorder[:mid+1])
	root.Right = buildTreev12(preorder[mid+1:], inorder[mid+1:])

	return root
}

func buildTreev11(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	findMid := func(arr []int, target int) int {

		for i, v := range arr {
			if v == target {
				return i
			}
		}

		return -1
	}
	mid := findMid(inorder, root.Val)

	root.Left = buildTreev11(preorder[1:mid+1], inorder[:mid+1])
	root.Right = buildTreev11(preorder[mid+1:], inorder[mid+1:])

	return root
}

func buildTreev10(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	findMid := func(arr []int, target int) int {

		for i, v := range arr {
			if v == target {
				return i
			}
		}

		return -1
	}
	mid := findMid(inorder, root.Val)

	root.Left = buildTreev10(preorder[1:mid+1], inorder[:mid+1])
	root.Right = buildTreev10(preorder[mid+1:], inorder[mid+1:])

	return root
}

func buildTreev9(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	findMid := func(arr []int, target int) int {
		for i, v := range arr {
			if v == target {
				return i
			}
		}

		return -1
	}

	mid := findMid(inorder, root.Val)

	root.Left = buildTreev9(preorder[1:mid+1], inorder[:mid+1])
	root.Right = buildTreev9(preorder[mid+1:], inorder[mid+1:])
	return root
}

func buildTreev8(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	findMid := func(arr []int, target int) int {

		for i, v := range arr {
			if v == target {
				return i
			}
		}
		return -1
	}

	root := &TreeNode{Val: preorder[0]}
	mid := findMid(inorder, root.Val)
	root.Left = buildTreev8(preorder[1:mid+1], inorder[:mid+1])
	root.Right = buildTreev8(preorder[mid+1:], inorder[mid+1:])
	return root
}

func buildTreev7(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	findMidIdx := func(arr []int, val int) int {

		for i, v := range arr {
			if v == val {
				return i
			}
		}

		return -1
	}
	root := &TreeNode{Val: preorder[0]}
	mid := findMidIdx(inorder, root.Val)
	root.Left = buildTreev7(preorder[1:mid+1], inorder[:mid+1])
	root.Right = buildTreev7(preorder[mid+1:], inorder[mid+1:])

	return root
}

func buildTreev6(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	findMid := func(arr []int, val int) int {
		for i, v := range arr {
			if v == val {
				return i
			}
		}
		return -1
	}

	// root first in preorder
	root := &TreeNode{Val: preorder[0]}
	mid := findMid(inorder, root.Val)
	root.Left = buildTreev6(preorder[1:mid+1], inorder[:mid+1])
	root.Right = buildTreev6(preorder[mid+1:], inorder[mid+1:])
	return root
}

func buildTreev5(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	findMidIdx := func(in []int, target int) int {
		for i, v := range in {
			if v == target {
				return i
			}
		}
		return -1
	}
	root := &TreeNode{Val: preorder[0]}
	mid := findMidIdx(inorder, root.Val)
	root.Left = buildTreev5(preorder[1:mid+1], inorder[:mid+1])
	root.Right = buildTreev5(preorder[mid+1:], inorder[mid+1:])
	return root
}

func buildTreev4(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]} // preorder: root first
	finMidIndex := func(arr []int, val int) int {
		for i, v := range arr {
			if v == val {
				return i
			}
		}
		return -1
	}
	mid := finMidIndex(inorder, root.Val)
	root.Left = buildTreev4(preorder[1:mid+1], inorder[:mid+1])
	root.Right = buildTreev4(preorder[mid+1:], inorder[mid+1:])
	return root
}

func buildTreev3(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	index := func(arr []int, val int) int {
		for i, v := range arr {
			if v == val {
				return i
			}
		}
		return -1
	}
	mid := index(inorder, root.Val)
	root.Left = buildTreev3(preorder[1:mid+1], inorder[:mid+1])
	root.Right = buildTreev3(preorder[mid+1:], inorder[mid+1:])
	return root
}

func buildTreev2(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	index := func(arr []int, val int) int {
		for i, v := range arr {
			if v == val {
				return i
			}
		}
		return -1
	}

	root := &TreeNode{Val: preorder[0]}
	mid := index(inorder, root.Val)
	root.Left = buildTreev2(preorder[1:mid+1], inorder[:mid+1])
	root.Right = buildTreev2(preorder[mid+1:], inorder[mid+1:])
	return root
}

func buildTreev1(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	index := func(arr []int, val int) int {
		for i, v := range arr {
			if v == val {
				return i
			}
		}
		return -1
	}

	root := &TreeNode{Val: preorder[0]}
	mid := index(inorder, preorder[0])
	root.Left = buildTreev1(preorder[1:mid+1], inorder[:mid])
	root.Right = buildTreev1(preorder[mid+1:], inorder[mid+1:])
	return root
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	index := func(arr []int, val int) int {
		for i, v := range arr {
			if v == val {
				return i
			}
		}
		return -1
	}

	root := &TreeNode{Val: preorder[0]}
	mid := index(inorder, preorder[0])
	root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])
	return root
}
