package b75

func buildTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	root := &TreeNode{Val: nums[0]}
	queue := []*TreeNode{root}
	i := 1

	for i < len(nums) {
		node := queue[0]
		queue = queue[1:]

		if nums[i] != -101 {
			node.Left = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(nums) && nums[i] != -101 {
			node.Right = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// buggy
func heapArrayToTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if nums[0] == -101 {
		return nil
	}
	root := &TreeNode{Val: nums[0]}
	if nums[1] != -101 {
		root.Left = heapArrayToTree(nums[1:])
	}
	if nums[2] != -101 {
		root.Right = heapArrayToTree(nums[len(nums)/2+1:])
	}
	return root
}
