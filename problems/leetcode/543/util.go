package b75

// bfs way
func buildTree(in []int) *TreeNode {

	if len(in) == 0 {
		return nil
	}

	root := &TreeNode{Val: in[0]}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	i := 1

	for i < len(in) {

		// pop node from queue
		n := q[0]
		q = q[1:]

		// append left node
		n.Left = &TreeNode{Val: in[i]}
		q = append(q, n.Left)
		i++

		// append right node
		if i < len(in) {
			n.Right = &TreeNode{Val: in[i]}
			q = append(q, n.Right)
		}
		i++
	}

	return root
}
