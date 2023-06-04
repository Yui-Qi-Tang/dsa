package b75

/*
138. Copy List with Random Pointer
Medium
10.9K
1.1K
Companies
A linked list of length n is given such that each node contains an additional random pointer, which could point to any node in the list, or null.

Construct a deep copy of the list. The deep copy should consist of exactly n brand new nodes, where each new node has its value set to the value of its corresponding original node. Both the next and random pointer of the new nodes should point to new nodes in the copied list such that the pointers in the original list and copied list represent the same list state. None of the pointers in the new list should point to nodes in the original list.

For example, if there are two nodes X and Y in the original list, where X.random --> Y, then for the corresponding two nodes x and y in the copied list, x.random --> y.

Return the head of the copied linked list.

The linked list is represented in the input/output as a list of n nodes. Each node is represented as a pair of [val, random_index] where:

val: an integer representing Node.val
random_index: the index of the node (range from 0 to n-1) that the random pointer points to, or null if it does not point to any node.
Your code will only be given the head of the original linked list.



Example 1:


Input: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
Output: [[7,null],[13,0],[11,4],[10,2],[1,0]]
Example 2:


Input: head = [[1,1],[2,1]]
Output: [[1,1],[2,1]]
Example 3:



Input: head = [[3,null],[3,0],[3,null]]
Output: [[3,null],[3,0],[3,null]]


Constraints:

0 <= n <= 1000
-104 <= Node.val <= 104
Node.random is null or is pointing to some node in the linked list.
*/

// Node Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomListv29(head *Node) *Node {
	m := map[*Node]*Node{}

	p := head
	for p != nil {
		m[p] = &Node{Val: p.Val}
		p = p.Next
	}

	p = head
	for p != nil {
		clone := m[p]
		clone.Next = m[p.Next]
		clone.Random = m[p.Random]
		p = p.Next
	}

	return m[head]
}

func copyRandomListv28(head *Node) *Node {
	m := map[*Node]*Node{}

	p := head
	for p != nil {
		m[p] = &Node{Val: p.Val}
		p = p.Next
	}

	p = head

	for p != nil {
		clone := m[p]
		clone.Next = m[p.Next]
		clone.Random = m[p.Random]
		p = p.Next
	}

	return m[head]
}

func copyRandomListv27(head *Node) *Node {
	m := map[*Node]*Node{}

	p := head
	for p != nil {
		m[p] = &Node{Val: p.Val}
		p = p.Next
	}

	p = head
	for p != nil {
		clone := m[p]
		clone.Next = m[p.Next]
		clone.Random = m[p.Random]
		p = p.Next
	}

	return m[head]
}

func copyRandomListv26(head *Node) *Node {
	m := map[*Node]*Node{}

	p := head
	for p != nil {
		m[p] = &Node{Val: p.Val}
		p = p.Next
	}

	p = head

	for p != nil {
		clone := m[p]
		clone.Next = m[p.Next]
		clone.Random = m[p.Random]
		p = p.Next
	}

	return m[head]
}

func copyRandomListv25(head *Node) *Node {
	m := make(map[*Node]*Node)

	p := head
	for p != nil {
		m[p] = &Node{Val: p.Val}
		p = p.Next
	}

	p = head
	for p != nil {
		clone := m[p]
		clone.Next = m[p.Next]
		clone.Random = m[p.Random]
		p = p.Next
	}

	return m[head]
}

func copyRandomListv24(head *Node) *Node {
	m := map[*Node]*Node{}

	p := head
	for p != nil {
		m[p] = &Node{Val: p.Val}
		p = p.Next
	}

	p = head

	for p != nil {
		clone := m[p]
		clone.Next = m[p.Next]
		clone.Random = m[p.Random]
		p = p.Next
	}

	return m[head]
}

func copyRandomListv23(head *Node) *Node {
	m := map[*Node]*Node{}

	p := head
	for p != nil {
		m[p] = &Node{Val: p.Val}
		p = p.Next
	}

	p = head
	for p != nil {
		clone := m[p]
		clone.Next = m[p.Next]
		clone.Random = m[p.Random]
		p = p.Next
	}

	return m[head]
}

func copyRandomListv22(head *Node) *Node {
	m := map[*Node]*Node{}

	p := head
	for p != nil {
		m[p] = &Node{Val: p.Val}
		p = p.Next
	}

	p = head
	for p != nil {
		clone := m[p]
		clone.Next = m[p.Next]
		clone.Random = m[p.Random]
		p = p.Next
	}

	return m[head]
}

func copyRandomListv21(head *Node) *Node {
	m := make(map[*Node]*Node, 0)

	p := head
	for p != nil {
		m[p] = &Node{Val: p.Val}
		p = p.Next
	}

	p = head
	for p != nil {
		clone := m[p]
		clone.Next = m[p.Next]
		clone.Random = m[p.Random]
		p = p.Next
	}

	return m[head]
}

func copyRandomListv20(head *Node) *Node {
	m := make(map[*Node]*Node)

	p := head
	for p != nil {
		m[p] = &Node{Val: p.Val}
		p = p.Next
	}

	p = head
	for p != nil {
		clone := m[p]
		clone.Next = m[p.Next]
		clone.Random = m[p.Random]
		p = p.Next
	}

	return m[head]
}

func copyRandomListv19(head *Node) *Node {
	m := make(map[*Node]*Node)

	p := head
	for p != nil {
		m[p] = &Node{Val: p.Val}
		p = p.Next
	}

	p = head
	for p != nil {
		clone := m[p]
		clone.Next = m[p.Next]
		clone.Random = m[p.Random]
		p = p.Next
	}

	return m[head]
}

func copyRandomListv18(head *Node) *Node {
	m := make(map[*Node]*Node)

	p := head
	for p != nil {
		m[p] = &Node{Val: p.Val}
		p = p.Next
	}

	p = head
	for p != nil {
		clone := m[p]
		clone.Next = m[p.Next]
		clone.Random = m[p.Random]
		p = p.Next
	}

	return m[head]
}

func copyRandomListv17(head *Node) *Node {
	m := make(map[*Node]*Node)

	p := head

	for p != nil {
		m[p] = &Node{Val: p.Val}
		p = p.Next
	}

	p = head
	for p != nil {
		clone := m[p]
		clone.Next = m[p.Next]
		clone.Random = m[p.Random]
		p = p.Next
	}

	return m[head]
}

func copyRandomListv16(head *Node) *Node {
	m := make(map[*Node]*Node)
	p := head

	for p != nil {
		m[p] = &Node{Val: p.Val}
		p = p.Next
	}

	p = head

	for p != nil {
		clone := m[p]
		clone.Next = m[p.Next]
		clone.Random = m[p.Random]
		p = p.Next
	}

	return m[head]
}

func copyRandomListv15(head *Node) *Node {
	m := make(map[*Node]*Node)
	p := head

	for p != nil {
		m[p] = &Node{Val: p.Val}
		p = p.Next
	}

	p = head
	for p != nil {
		clone := m[p]
		clone.Next = m[p.Next]
		clone.Random = m[p.Random]
		p = p.Next
	}

	return m[head]
}

func copyRandomListv14(head *Node) *Node {
	m := make(map[*Node]*Node)
	p := head

	for p != nil {
		m[p] = &Node{Val: p.Val}
		p = p.Next
	}

	p = head
	for p != nil {
		clone := m[p]
		clone.Next = m[p.Next]
		clone.Random = m[p.Random]
		p = p.Next
	}

	return m[head]
}

func copyRandomListv13(head *Node) *Node {
	m := make(map[*Node]*Node)

	p := head
	for p != nil {
		m[p] = &Node{Val: p.Val}
		p = p.Next
	}

	p = head
	for p != nil {
		clone := m[p]
		clone.Next = m[p.Next]
		clone.Random = m[p.Random]
		p = p.Next
	}

	return m[head]
}

func copyRandomListv12(head *Node) *Node {
	m := make(map[*Node]*Node)

	p := head
	for p != nil {
		m[p] = &Node{Val: p.Val}
		p = p.Next
	}

	p = head
	for p != nil {
		clone := m[p]
		clone.Next = m[p.Next]
		clone.Random = m[p.Random]
		p = p.Next
	}

	return m[head]
}

func copyRandomListv11(head *Node) *Node {
	m := make(map[*Node]*Node, 0)

	p := head
	for p != nil {
		m[p] = &Node{Val: p.Val}
		p = p.Next
	}

	p = head

	for p != nil {
		clone := m[p]
		clone.Next = m[p.Next]
		clone.Random = m[p.Random]
		p = p.Next
	}

	return m[head]
}

func copyRandomListv10(head *Node) *Node {
	m := make(map[*Node]*Node)

	p := head
	for p != nil {
		m[p] = &Node{Val: p.Val}
		p = p.Next
	}

	p = head
	for p != nil {
		clone := m[p]
		clone.Next = m[p.Next]
		clone.Random = m[p.Random]
		p = p.Next
	}

	return m[head]
}

func copyRandomListv9(head *Node) *Node {
	m := make(map[*Node]*Node) // old: clone
	p := head
	for p != nil {
		m[p] = &Node{Val: p.Val}
		p = p.Next
	}

	p = head
	for p != nil {
		clone := m[p]
		clone.Next = m[p.Next]     // next of clone
		clone.Random = m[p.Random] // random of clone
		p = p.Next
	}

	return m[head]
}

func copyRandomListv8(head *Node) *Node {
	m := make(map[*Node]*Node)

	p := head
	for p != nil {
		m[p] = &Node{Val: p.Val}
		p = p.Next
	}

	p = head
	for p != nil {
		clone := m[p]
		clone.Next = m[p.Next]
		clone.Random = m[p.Random]
		p = p.Next
	}

	return m[head]
}

func copyRandomListv7(head *Node) *Node {
	m := make(map[*Node]*Node, 0)

	p := head
	for p != nil {
		m[p] = &Node{Val: p.Val}
		p = p.Next
	}

	p = head
	for p != nil {
		clone := m[p]
		clone.Next = m[p.Next]
		clone.Random = m[p.Random]
		p = p.Next
	}

	return m[head]
}

func copyRandomListv6(head *Node) *Node {
	m := make(map[*Node]*Node)
	p := head

	for p != nil {
		m[p] = &Node{Val: p.Val}
		p = p.Next
	}

	p = head

	for p != nil {
		clone := m[p]
		clone.Next = m[p.Next]
		clone.Random = m[p.Random]
		p = p.Next
	}

	return m[head]
}

func copyRandomListv5(head *Node) *Node {
	m := make(map[*Node]*Node)
	// first: copy current node
	p := head
	for p != nil {
		m[p] = &Node{Val: p.Val}
		p = p.Next
	}

	p = head

	for p != nil {
		cNode := m[p]
		cNode.Next = m[p.Next]
		cNode.Random = m[p.Random]
		p = p.Next
	}

	return m[head]
}

func copyRandomListv4(head *Node) *Node {
	m := make(map[*Node]*Node)

	p := head
	for p != nil {
		m[p] = &Node{Val: p.Val}
		p = p.Next
	}

	p = head
	for p != nil {
		np := m[p]
		np.Next = m[p.Next]
		np.Random = m[p.Random]
		p = p.Next
	}

	return m[head]
}

func copyRandomListv3(head *Node) *Node {
	m := make(map[*Node]*Node) // orig: copy

	p := head
	for p != nil {
		clone := &Node{Val: p.Val}
		m[p] = clone
		p = p.Next
	}

	p = head

	for p != nil {
		m[p].Next = m[p.Next]
		m[p].Random = m[p.Random]
		p = p.Next
	}

	return m[head]
}

// using two pass algorithm
func copyRandomListv2(head *Node) *Node {
	cache := make(map[*Node]*Node) // old: copy

	p := head

	for p != nil {
		clone := &Node{Val: p.Val}
		cache[p] = clone
		p = p.Next
	}

	p = head
	for p != nil {
		cache[p].Next = cache[p.Next]
		cache[p].Random = cache[p.Random]
		p = p.Next
	}

	return cache[head]
}

func copyRandomListv1(head *Node) *Node {
	if head == nil {
		return nil
	}

	m := make(map[*Node]int)
	p := head

	idx := 0
	for p != nil {
		m[p] = idx
		idx++
		p = p.Next
	}

	in := make([][2]int, 0)

	p = head

	for p != nil {
		data := [2]int{}
		data[0] = p.Val
		if p.Random != nil {
			data[1] = m[p.Random]
		} else {
			data[1] = -1
		}
		in = append(in, data)
		p = p.Next
	}

	// fmt.Println(in)

	return createList(in)
}

func createList(in [][2]int) *Node {
	if len(in) == 0 {
		return nil
	}

	head := &Node{Val: in[0][0]}
	p := head
	m := make(map[int]*Node)
	m[0] = p // a store for saving the index:node pointer

	for i := 1; i < len(in); i++ {
		p.Next = &Node{Val: in[i][0]}
		p = p.Next
		m[i] = p
	}

	p = head
	for i := 0; i < len(in); i++ {
		p.Random = m[in[i][1]]
		p = p.Next
	}

	return head
}
