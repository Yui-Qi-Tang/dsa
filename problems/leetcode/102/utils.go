package b75

const null int = 1001

func buildTree(in []int) *TreeNode {
	if len(in) == 0 {
		return nil
	}

	root := &TreeNode{Val: in[0]}
	q := make([]*TreeNode, 0)
	q = append(q, root)

	i := 1
	for i < len(in) {
		n := q[0]
		q = q[1:]

		if in[i] != null {
			n.Left = &TreeNode{Val: in[i]}
			q = append(q, n.Left)
		}
		i++

		if i < len(in) && in[i] != null {
			n.Right = &TreeNode{Val: in[i]}
			q = append(q, n.Right)
		}
		i++
	}

	return root
}
