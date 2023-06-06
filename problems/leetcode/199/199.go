package b75

/*
199. Binary Tree Right Side View
Medium
9.8K
593
Companies
Given the root of a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.



Example 1:


Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]
Example 2:

Input: root = [1,null,3]
Output: [1,3]
Example 3:

Input: root = []
Output: []


Constraints:

The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100
Accepted
937.6K
Submissions
1.5M
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

func rightSideViewv18(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		result = append(result, queue[len(queue)-1].Val)
		qLen := len(queue)
		for qLen > 0 {
			qLen--
			n := queue[0]
			queue = queue[1:]
			if n.Left != nil {
				queue = append(queue, n.Left)
			}

			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}
	}

	return result
}

func rightSideViewv17(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		result = append(result, queue[len(queue)-1].Val)
		qLen := len(queue)

		for qLen > 0 {
			n := queue[0]
			queue = queue[1:]
			qLen--

			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}
	}

	return result
}

func rightSideViewv16(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		result = append(result, queue[len(queue)-1].Val)
		qLen := len(queue)

		for qLen > 0 {
			qLen--
			n := queue[0]
			queue = queue[1:]

			if n.Left != nil {
				queue = append(queue, n.Left)
			}

			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}
	}

	return result
}

func rightSideViewv15(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		result = append(result, queue[len(queue)-1].Val)
		qLen := len(queue)
		for qLen > 0 {
			qLen--
			n := queue[0]
			queue = queue[1:]
			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}
	}

	return result
}

func rightSideViewv14(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		result = append(result, queue[len(queue)-1].Val)
		qLen := len(queue)

		for qLen > 0 {
			qLen--

			n := queue[0]
			queue = queue[1:]

			if n.Left != nil {
				queue = append(queue, n.Left)
			}

			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}
	}

	return result
}

func rightSideViewv13(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		result = append(result, queue[len(queue)-1].Val)
		qLen := len(queue)
		for qLen > 0 {
			qLen--
			n := queue[0]
			queue = queue[1:]
			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}
	}

	return result
}

func rightSideViewv12(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		result = append(result, queue[len(queue)-1].Val)
		qLen := len(queue)
		for qLen > 0 {
			n := queue[0]
			queue = queue[1:]
			qLen--

			if n.Left != nil {
				queue = append(queue, n.Left)
			}

			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}
	}

	return result
}

func rightSideViewv11(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		result = append(result, queue[len(queue)-1].Val)
		qLen := len(queue)
		for qLen > 0 {
			qLen--
			n := queue[0]
			queue = queue[1:]
			if n.Left != nil {
				queue = append(queue, n.Left)
			}

			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}
	}

	return result
}

func rightSideViewv10(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		result = append(result, queue[len(queue)-1].Val)

		qLen := len(queue)
		for qLen > 0 {

			n := queue[0]
			queue = queue[1:]
			if n.Left != nil {
				queue = append(queue, n.Left)
			}

			if n.Right != nil {
				queue = append(queue, n.Right)
			}

			qLen--
		}
	}

	return result
}

func rightSideViewv9(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		result = append(result, queue[len(queue)-1].Val)
		qLen := len(queue)
		for qLen > 0 {
			n := queue[0]
			queue = queue[1:]
			qLen--

			if n.Left != nil {
				queue = append(queue, n.Left)
			}

			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}
	}

	return result
}

func rightSideViewv8(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		result = append(result, queue[len(queue)-1].Val)
		qLen := len(queue)
		for qLen > 0 {

			n := queue[0]
			queue = queue[1:]

			if n.Left != nil {
				queue = append(queue, n.Left)
			}

			if n.Right != nil {
				queue = append(queue, n.Right)
			}

			qLen--
		}

	}

	return result
}

func rightSideViewv7(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := []*TreeNode{root}
	result := make([]int, 0)

	for len(queue) > 0 {
		result = append(result, queue[len(queue)-1].Val)
		qLen := len(queue)
		for qLen > 0 {
			n := queue[0]
			queue = queue[1:]
			if n.Left != nil {
				queue = append(queue, n.Left)
			}

			if n.Right != nil {
				queue = append(queue, n.Right)
			}
			qLen--
		}

	}

	return result
}

func rightSideViewv6(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		result = append(result, queue[len(queue)-1].Val)

		qLen := len(queue)
		for qLen > 0 {
			n := queue[0]
			queue = queue[1:]

			if n.Left != nil {
				queue = append(queue, n.Left)
			}

			if n.Right != nil {
				queue = append(queue, n.Right)
			}
			qLen--
		}
	}

	return result
}

func rightSideViewv5(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		result = append(result, queue[len(queue)-1].Val)
		qLen := len(queue)
		for qLen > 0 {

			n := queue[0]
			queue = queue[1:]
			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
			qLen--
		}
	}

	return result
}

func rightSideViewv4(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		result = append(result, queue[len(queue)-1].Val)
		qLen := len(queue)
		for qLen > 0 {

			n := queue[0]
			queue = queue[1:]

			if n.Left != nil {
				queue = append(queue, n.Left)
			}

			if n.Right != nil {
				queue = append(queue, n.Right)
			}

			qLen--
		}

	}

	return result
}

func rightSideViewv3(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		result = append(result, queue[len(queue)-1].Val)
		qLen := len(queue)
		for qLen > 0 {

			n := queue[0]
			queue = queue[1:]
			if n.Left != nil {
				queue = append(queue, n.Left)
			}

			if n.Right != nil {
				queue = append(queue, n.Right)
			}

			qLen--
		}

	}

	return result
}

func rightSideViewv2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	q := make([]*TreeNode, 0)
	q = append(q, root)
	result := make([]int, 0)

	for len(q) > 0 {
		result = append(result, q[len(q)-1].Val)
		qLen := len(q)
		for qLen > 0 {
			n := q[0]
			q = q[1:]
			if n.Left != nil {
				q = append(q, n.Left)
			}

			if n.Right != nil {
				q = append(q, n.Right)
			}

			qLen--
		}
	}

	return result
}

func rightSideViewv1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0, 10) //10: because the max length of the input is 101 and we just collect the right side by binary tree, so, the max result is 10(log 101)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		result = append(result, queue[len(queue)-1].Val)
		qLen := len(queue)
		for qLen > 0 {
			n := queue[0]
			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
			queue = queue[1:]
			qLen--
		}
	}

	return result
}

func rightSideView(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}

	buffer := make([]*TreeNode, 0)
	buffer = append(buffer, root) // the right is put in the right of this buffer
	result := make([]int, 0)

	for len(buffer) > 0 {
		// append the right into the result
		result = append(result, buffer[len(buffer)-1].Val)
		// put the next level nodes into the buffer
		bufferLen := len(buffer)
		for bufferLen > 0 {
			n := buffer[0]
			if n.Left != nil {
				buffer = append(buffer, n.Left)
			}
			if n.Right != nil {
				buffer = append(buffer, n.Right)
			}
			bufferLen--
			buffer = buffer[1:]
		}
	}

	return result
}
