package b75

/*
1448. Count Good Nodes in Binary Tree
Medium
4.7K
124
Companies
Given a binary tree root, a node X in the tree is named good if in the path from root to X there are no nodes with a value greater than X.

Return the number of good nodes in the binary tree.



Example 1:



Input: root = [3,1,4,3,null,1,5]
Output: 4
Explanation: Nodes in blue are good.
Root Node (3) is always a good node.
Node 4 -> (3,4) is the maximum value in the path starting from the root.
Node 5 -> (3,4,5) is the maximum value in the path
Node 3 -> (3,1,3) is the maximum value in the path.
Example 2:



Input: root = [3,3,null,4,2]
Output: 3
Explanation: Node 2 -> (3, 3, 2) is not good, because "3" is higher than it.
Example 3:

Input: root = [1]
Output: 1
Explanation: Root is considered as good.


Constraints:

The number of nodes in the binary tree is in the range [1, 10^5].
Each node's value is between [-10^4, 10^4].
Accepted
315.1K
Submissions
424.3K
Acceptance Rate
74.3%
*/

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodesv14(root *TreeNode) int {

	var gn func(n *TreeNode, preMax int) int

	gn = func(n *TreeNode, preMax int) int {
		if n == nil {
			return 0
		}

		result := 0
		if n.Val >= preMax {
			result++
			preMax = n.Val
		}

		return result + gn(n.Left, preMax) + gn(n.Right, preMax)
	}

	return gn(root, root.Val)
}

func goodNodesv13(root *TreeNode) int {

	var gn func(n *TreeNode, preMax int) int

	gn = func(n *TreeNode, preMax int) int {
		if n == nil {
			return 0
		}

		result := 0
		if n.Val >= preMax {
			preMax = n.Val
			result++
		}

		return result + gn(n.Left, preMax) + gn(n.Right, preMax)
	}

	return gn(root, root.Val)
}

func goodNodesv12(root *TreeNode) int {
	var gn func(n *TreeNode, preMax int) int

	gn = func(n *TreeNode, preMax int) int {
		if n == nil {
			return 0
		}

		result := 0
		if n.Val >= preMax {
			preMax = n.Val
			result++
		}

		return result + gn(n.Left, preMax) + gn(n.Right, preMax)
	}

	return gn(root, root.Val)
}

func goodNodesv11(root *TreeNode) int {
	var gn func(n *TreeNode, preMax int) int
	gn = func(n *TreeNode, preMax int) int {
		if n == nil {
			return 0
		}

		result := 0
		if n.Val >= preMax {
			result++
			preMax = n.Val
		}

		return result + gn(n.Left, preMax) + gn(n.Right, preMax)
	}
	return gn(root, root.Val)
}

func goodNodesv10(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var gn func(n *TreeNode, preMax int) int

	gn = func(n *TreeNode, preMax int) int {
		if n == nil {
			return 0
		}

		result := 0

		if n.Val >= preMax {
			preMax = n.Val
			result++
		}

		return result + gn(n.Left, preMax) + gn(n.Right, preMax)
	}

	return gn(root, root.Val)
}

func goodNodesv9(root *TreeNode) int {
	var gn func(n *TreeNode, preMax int) int

	gn = func(n *TreeNode, preMax int) int {
		if n == nil {
			return 0
		}

		result := 0
		if n.Val >= preMax {
			result++
			preMax = n.Val
		}

		return result + gn(n.Left, preMax) + gn(n.Right, preMax)
	}

	return gn(root, root.Val)
}

func goodNodesv8(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var gn func(n *TreeNode, preMax int) int

	gn = func(n *TreeNode, preMax int) int {
		if n == nil {
			return 0
		}

		result := 0
		if n.Val >= preMax {
			result++
			preMax = n.Val
		}

		return result + gn(n.Left, preMax) + gn(n.Right, preMax)
	}

	return gn(root, root.Val)
}

func goodNodesv7(root *TreeNode) int {

	var gn func(n *TreeNode, preMax int) int

	gn = func(n *TreeNode, preMax int) int {
		if n == nil {
			return 0
		}

		result := 0
		if n.Val >= preMax {
			result++
			preMax = n.Val
		}

		return result + gn(n.Left, preMax) + gn(n.Right, preMax)
	}

	return gn(root, root.Val)
}

func goodNodesv6(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var gn func(n *TreeNode, preMax int) int

	gn = func(n *TreeNode, preMax int) int {
		if n == nil {
			return 0
		}

		result := 0
		if n.Val >= preMax {
			result++
			preMax = n.Val
		}

		return result + gn(n.Left, preMax) + gn(n.Right, preMax)
	}

	return gn(root, root.Val)
}

func goodNodesv5(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var gn func(n *TreeNode, preMax int) int
	gn = func(n *TreeNode, preMax int) int {
		if n == nil {
			return 0
		}

		result := 0
		if n.Val >= preMax {
			result++
			preMax = n.Val
		}

		return result + gn(n.Left, preMax) + gn(n.Right, preMax)
	}

	return gn(root, root.Val)
}

func goodNodesv4(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var gn func(n *TreeNode, preMax int) int
	gn = func(n *TreeNode, preMax int) int {
		if n == nil {
			return 0
		}

		result := 0
		if n.Val >= preMax {
			preMax = n.Val
			result++
		}

		return result + gn(n.Left, preMax) + gn(n.Right, preMax)
	}

	return gn(root, root.Val)
}

func goodNodesv3(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var gn func(n *TreeNode, prexMax int) int

	gn = func(n *TreeNode, prexMax int) int {
		if n == nil {
			return 0
		}

		result := 0
		if n.Val >= prexMax {
			prexMax = n.Val
			result++
		}
		return result + gn(n.Left, prexMax) + gn(n.Right, prexMax)
	}

	return gn(root, root.Val)
}

func goodNodesv2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var gn func(n *TreeNode, prexMax int) int
	gn = func(n *TreeNode, prexMax int) int {

		if n == nil {
			return 0
		}

		result := 0
		if n.Val >= prexMax {
			prexMax = n.Val
			result++
		}

		return result + gn(n.Left, prexMax) + gn(n.Right, prexMax)

	}

	return 1 + gn(root.Left, root.Val) + gn(root.Right, root.Val)
}

func goodNodesv1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var gn func(n *TreeNode, prevMax int) int
	gn = func(n *TreeNode, prevMax int) int {
		if n == nil {
			return 0
		}

		if n.Val >= prevMax {
			prevMax = n.Val
		}

		l := gn(n.Left, prevMax)
		r := gn(n.Right, prevMax)

		if n.Val >= prevMax {
			return 1 + l + r
		}

		return l + r

	}

	l := gn(root.Left, root.Val)
	r := gn(root.Right, root.Val)

	return 1 + l + r
}

func goodNodes(root *TreeNode) int {

	if root == nil {
		return 0
	}

	// gn: goodnode
	var gn func(n *TreeNode, prevMax int) int
	gn = func(n *TreeNode, prevMax int) int {
		if n == nil {
			return 0
		}
		result := 0
		if n.Val >= prevMax {
			prevMax = n.Val
			result++
		}
		l := gn(n.Left, prevMax)
		r := gn(n.Right, prevMax)
		return result + l + r
	}

	l := gn(root.Left, root.Val)
	r := gn(root.Right, root.Val)

	return 1 + l + r
}
