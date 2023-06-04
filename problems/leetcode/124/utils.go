package b75

import "fmt"

const null int = 1001

func buildTree(in []int) *TreeNode {

	if len(in) == 0 {
		return nil
	}

	root := &TreeNode{Val: in[0]}
	queue := []*TreeNode{root}

	for i := 1; i < len(in); {

		n := queue[0]
		queue = queue[1:]
		if in[i] != null {
			n.Left = &TreeNode{Val: in[i]}
			queue = append(queue, n.Left)
		}
		i++

		if i < len(in) && in[i] != null {
			n.Right = &TreeNode{Val: in[i]}
			queue = append(queue, n.Right)
		}

		i++

	}

	return root
}

func inorderPrint(root *TreeNode) {
	if root == nil {
		fmt.Println("the tree is empty")
		return
	}

	stack := []*TreeNode{}

	result := []int{}
	for len(stack) > 0 || root != nil {

		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, n.Val)

		root = n.Right
	}

	fmt.Println(result)
}
