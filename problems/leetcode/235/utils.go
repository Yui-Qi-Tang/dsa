package b75

const null int = 10e9 + 1

func buildTree(in []int) *TreeNode {
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

func findRecusively(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if root.Val < val {
		return findRecusively(root.Right, val)
	}
	return findRecusively(root.Left, val)
}

func findIteration(root *TreeNode, val int) *TreeNode {
	p := root
	for p != nil {
		if p.Val == val {
			return p
		}

		if p.Val < val {
			p = p.Right
		} else {
			p = p.Left
		}
	}
	return nil
}
