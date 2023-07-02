package b75

import "fmt"

// buildTree creates binary tree (for now, only work on binary tree)
func buildTree(in []int) *TreeNode {
	if len(in) == 0 {
		return nil
	}

	root := &TreeNode{Val: in[0]}
	p := root
	for _, v := range in[1:] {

		for p.Left != nil && p.Right != nil {
			if v <= p.Val {
				p = p.Left
			} else {
				p = p.Right
			}
		}

		if v <= p.Val {
			p.Left = &TreeNode{Val: v}
		} else {
			p.Right = &TreeNode{Val: v}
		}

		p = root // back to root
	}

	return root
}

func postOrder(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Val)
	postOrder(root.Left)
	postOrder(root.Right)
}
