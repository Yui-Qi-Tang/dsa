package b75

const null int = 3001

func build(in []int) *TreeNode {
	if len(in) == 0 {
		return nil
	}

	root := &TreeNode{Val: in[0]}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	i := 1
	for i < len(in) {
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

// print x
// go x.left
// go x.right
func printPreorder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	result := make([]int, 0)

	for len(stack) > 0 {
		n := stack[len(stack)-1]
		result = append(result, n.Val)
		stack = stack[:len(stack)-1]

		if n.Right != nil {
			stack = append(stack, n.Right)
		}

		if n.Left != nil {
			stack = append(stack, n.Left)
		}
	}

	return result
}

// go x.left
// print x
// go x.right
func printInorder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := make([]*TreeNode, 0)
	result := make([]int, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		n := stack[len(stack)-1]
		result = append(result, n.Val)
		stack = stack[:len(stack)-1]
		root = n.Right
	}

	return result
}

// go x.left
// go x.right
// print x
func printPostorder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	result := make([]int, 0)

	for len(stack) > 0 {
		n := stack[len(stack)-1]
		result = append(result, n.Val)
		stack = stack[:len(stack)-1]

		if n.Left != nil {
			stack = append(stack, n.Left)
		}

		if n.Right != nil {
			stack = append(stack, n.Right)
		}
	}

	// reverse the result to get postorder traversal
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}
