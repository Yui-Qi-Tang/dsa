package b75

/*
141. Linked List Cycle
Easy

10757

920

Add to List

Share
Given head, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.



Example 1:


Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).
Example 2:


Input: head = [1,2], pos = 0
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.
Example 3:


Input: head = [1], pos = -1
Output: false
Explanation: There is no cycle in the linked list.


Constraints:

The number of the nodes in the list is in the range [0, 104].
-105 <= Node.val <= 105
pos is -1 or a valid index in the linked-list.


Follow up: Can you solve it using O(1) (i.e. constant) memory?
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCyclev30(head *ListNode) bool {
	s, f := head, head

	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
		if s == f {
			return true
		}
	}

	return false
}

func hasCyclev29(head *ListNode) bool {
	s, f := head, head
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
		if s == f {
			return true
		}
	}

	return false
}

func hasCyclev28(head *ListNode) bool {
	s, f := head, head

	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next

		if s == f {
			return true
		}
	}

	return false
}

func hasCyclev27(head *ListNode) bool {
	s, f := head, head
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next

		if s == f {
			return true
		}
	}

	return false
}

func hasCyclev26(head *ListNode) bool {
	s, f := head, head
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next

		if s == f {
			return true
		}
	}

	return false
}

func hasCyclev25(head *ListNode) bool {
	s, f := head, head
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
		if s == f {
			return true
		}
	}

	return false
}

func hasCyclev24(head *ListNode) bool {
	s, f := head, head

	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
		if s == f {
			return true
		}
	}

	return false
}

func hasCyclev23(head *ListNode) bool {
	s, f := head, head
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
		if s == f {
			return true
		}
	}

	return false
}

func hasCyclev22(head *ListNode) bool {
	s, f := head, head

	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
		if s == f {
			return true
		}
	}

	return false
}

func hasCyclev21(head *ListNode) bool {
	s, f := head, head
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
		if s == f {
			return true
		}
	}

	return false
}

func hasCyclev20(head *ListNode) bool {
	if head == nil {
		return false
	}
	s, f := head, head.Next
	for f != nil && f.Next != nil {

		s = s.Next
		f = f.Next.Next
		if s == f {
			return true
		}

	}

	return false
}

func hasCyclev19(head *ListNode) bool {
	s, f := head, head.Next
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next

		if s == f {
			return true
		}
	}

	return false
}

func hasCyclev18(head *ListNode) bool {
	if head == nil {
		return false
	}

	s, f := head, head.Next

	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
		if s == f {
			return true
		}
	}

	return false
}

func hasCyclev17(head *ListNode) bool {
	if head == nil {
		return false
	}

	s, f := head, head.Next
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next

		if s == f {
			return true
		}
	}

	return false
}

func hasCyclev16(head *ListNode) bool {

	if head == nil {
		return false
	}

	s, f := head, head
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next

		if s == f {
			return true
		}

	}

	return false
}

func hasCyclev15(head *ListNode) bool {

	s, f := head, head
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next

		if s == f {
			return true
		}
	}

	return false
}

func hasCyclev14(head *ListNode) bool {
	s, f := head, head.Next
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
		if s == f {
			return true
		}
	}

	return false
}

func hasCyclev13(head *ListNode) bool {
	if head == nil {
		return false
	}
	s, f := head, head.Next
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
		if s == f {
			return true
		}
	}
	return false
}

func hasCyclev12(head *ListNode) bool {
	if head == nil {
		return false
	}

	s, f := head, head.Next
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
		if s == f {
			return true
		}
	}

	return false
}

func hasCyclev11(head *ListNode) bool {
	s, f := head, head.Next
	if f.Next == nil {
		return false
	}

	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next

		if s == f {
			return true
		}
	}
	return false
}

func hasCyclev10(head *ListNode) bool {
	s, f := head, head.Next
	if f == nil {
		return false
	}

	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
		if s == f {
			return true
		}

	}

	return false
}

func hasCyclev9(head *ListNode) bool {
	s, f := head, head.Next
	if f == nil {
		return false
	}

	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next

		if s == f {
			return true
		}
	}

	return false
}

func hasCyclev8(head *ListNode) bool {
	if head == nil {
		return false
	}
	s, f := head, head.Next
	if f == nil {
		return false
	}

	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next

		if s == f {
			return true
		}
	}

	return false
}

func hasCyclev7(head *ListNode) bool {
	s, f := head, head.Next
	if f == nil {
		return false
	}

	for f != nil && f.Next != nil {
		if s == f {
			return true
		}

		s = s.Next
		f = f.Next.Next
	}

	return false
}

func hasCyclev6(head *ListNode) bool {
	s, f := head, head.Next
	if f == nil {
		return false
	}

	for f != nil && f.Next != nil {
		if s == f {
			return true
		}
		s = s.Next
		f = f.Next.Next
	}

	return false
}

func hasCyclev5(head *ListNode) bool {

	if head == nil {
		return false
	}

	s, f := head, head.Next
	for f != nil && f.Next != nil {

		if s == f {
			return true
		}

		s = s.Next
		f = f.Next.Next
	}

	return false
}

func hasCyclev4(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head.Next
	if fast == nil || fast.Next == nil {
		return false
	}

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}

func hasCyclev3(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next
	for fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}

func hasCyclev2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {

		if slow == fast {
			return true
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}

func hasCyclev1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {

		if slow == fast {
			return true
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}

/*
// my env

func hasCycle(head *ListNode) bool {

	m := make(map[*ListNode]int) // addr:value
	for head != nil {

		if _, exist := m[head]; exist {
			return true
		}
		m[head] = head.Val
		head = head.Next
	}

	return false
}

func hasCycle2Pointers(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}

	p1 := head
	p2 := head.Next

	for p1 != nil && p2 != nil && p2.Next != nil {

		if p1 == p2 {
			return true
		}
		p1 = p1.Next
		p2 = p2.Next.Next

	}

	return false
}

func hasCyclev1(head *ListNode) bool {

	if head == nil {
		return false
	}

	slow, fast := head, head.Next

	for slow != nil && fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}

func hasCyclev2(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head.Next

	for slow != nil && fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}
*/
