package b75

const null int = 101

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
