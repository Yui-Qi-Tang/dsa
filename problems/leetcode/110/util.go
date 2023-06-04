package b75

import "fmt"

var null int = -101

func buildTree(in []int) *TreeNode {

	if len(in) == 0 {
		return nil
	}

	root := &TreeNode{Val: in[0]}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	i := 1
	for i < len(in) {
		node := queue[0]
		queue = queue[1:]

		if in[i] != null {
			node.Left = &TreeNode{Val: in[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(in) && in[i] != null {
			node.Right = &TreeNode{Val: in[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func dump(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Val)
	dump(root.Left)
	dump(root.Right)
}
