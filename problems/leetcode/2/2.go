package b75

/*

2. Add Two Numbers
Medium
25.1K
4.9K
Companies
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.



Example 1:


Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
Example 2:

Input: l1 = [0], l2 = [0]
Output: [0]
Example 3:

Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]


Constraints:

The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros.

*/

/*
Definition for singly-linked list.
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbersv29(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	p := result

	carry := 0

	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{Val: (carry + sum) % 10}
		p = p.Next

		carry = (carry + sum) / 10
	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}

	return result.Next
}

func addTwoNumbersv28(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	p := result
	carry := 0

	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{Val: (sum + carry) % 10}
		p = p.Next
		carry = (sum + carry) / 10
	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}

	return result.Next
}

func addTwoNumbersv27(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	p := result
	carry := 0
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{Val: (sum + carry) % 10}
		p = p.Next
		carry = (sum + carry) / 10
	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}

	return result.Next
}

func addTwoNumbersv26(l1, l2 *ListNode) *ListNode {
	result := new(ListNode)
	p := result

	carry := 0
	for l1 != nil || l2 != nil {

		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{Val: (sum + carry) % 10}
		p = p.Next

		carry = (sum + carry) / 10

	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}

	return result.Next
}

func addTwoNumbersv25(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	p := result
	carry := 0

	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{Val: (sum + carry) % 10}
		p = p.Next

		carry = (sum + carry) / 10
	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}

	return result.Next
}

func addTwoNumbersv24(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	p := result
	carry := 0

	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{Val: (sum + carry) % 10}
		p = p.Next

		carry = (sum + carry) / 10
	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}

	return result.Next
}

func addTwoNumbersv23(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	p := result
	carry := 0

	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{Val: (sum + carry) % 10}
		p = p.Next
		carry = (sum + carry) / 10
	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}

	return result.Next
}

func addTwoNumbersv22(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	p := result
	carry := 0

	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{Val: (sum + carry) % 10}
		p = p.Next

		carry = (sum + carry) / 10
	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}
	return result.Next
}

func addTwoNumbersv21(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	p := result
	carry := 0

	for l1 != nil || l2 != nil {
		sum := 0

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{Val: (sum + carry) % 10}
		p = p.Next

		carry = (sum + carry) / 10

	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}

	return result.Next
}

func addTwoNumbersv20(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	p := result
	carry := 0

	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{Val: (sum + carry) % 10}
		p = p.Next
		carry = (sum + carry) / 10
	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}

	return result.Next
}

func addTwoNumbersv19(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	p := result
	carry := 0

	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{Val: (sum + carry) % 10}
		p = p.Next
		carry = (sum + carry) / 10
	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}

	return result.Next
}

func addTwoNumbersv18(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	p := result

	carry := 0
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{Val: (sum + carry) % 10}
		p = p.Next

		carry = (sum + carry) / 10

	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}

	return result.Next
}

func addTwoNumbersv17(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	p := result
	carry := 0

	for l1 != nil || l2 != nil {
		sum := 0

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{Val: (sum + carry) % 10}
		p = p.Next

		carry = (sum + carry) / 10
	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}

	return result.Next
}

func addTwoNumbersv16(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	p := result
	carry := 0

	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{Val: (sum + carry) % 10}
		p = p.Next

		carry = (sum + carry) / 10
	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}
	return result.Next
}

func addTwoNumbersv15(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	p := result
	carry := 0
	for l1 != nil || l2 != nil {
		sum := 0

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{Val: (sum + carry) % 10}
		p = p.Next

		carry = (sum + carry) / 10
	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}

	return result.Next
}

func addTwoNumbersv14(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	p := result
	carry := 0

	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{Val: (sum + carry) % 10}
		p = p.Next

		carry = (sum + carry) / 10
	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}

	return result.Next
}

func addTwoNumbersv13(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	p := result
	carry := 0
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{Val: (sum + carry) % 10}
		p = p.Next
		carry = (sum + carry) / 10
	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}

	return result.Next
}

func addTwoNumbersv12(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	p := result
	carry := 0

	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{Val: (sum + carry) % 10}
		p = p.Next
		carry = (sum + carry) / 10
	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}

	return result.Next
}

func addTwoNumbersv11(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	p := result
	carry := 0
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{Val: (sum + carry) % 10}
		p = p.Next
		carry = (sum + carry) / 10
	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}

	return result.Next
}

func addTwoNumbersv10(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	p := result
	carry := 0

	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{Val: (sum + carry) % 10}
		p = p.Next
		carry = (sum + carry) / 10
	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}

	return result.Next
}

func addTwoNumbersv9(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	p := result

	carry := 0
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{Val: (carry + sum) % 10}
		p = p.Next
		carry = (sum + carry) / 10
	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}

	return result.Next
}

func addTwoNumbersv8(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	p := result
	carry := 0

	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{Val: (sum + carry) % 10}
		p = p.Next

		carry = (sum + carry) / 10
	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}

	return result.Next
}

func addTwoNumbersv7(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	p := result
	carry := 0

	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{Val: (carry + sum) % 10}
		p = p.Next
		carry = (carry + sum) / 10
	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}

	return result.Next
}

func addTwoNumbersv6(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	p := result
	carry := 0
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{Val: (sum + carry) % 10}
		p = p.Next
		carry = (sum + carry) / 10
	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}

	return result.Next
}

func addTwoNumbersv5(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	p := result
	carry := 0
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{Val: (sum + carry) % 10}
		p = p.Next
		carry = (sum + carry) / 10
	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}

	return result.Next
}

func addTwoNumbersv4(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	p := head
	carry := 0
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{Val: (carry + sum) % 10}
		p = p.Next
		carry = (carry + sum) / 10
	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}

	return head.Next
}

func addTwoNumbersv3(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	p := head
	carry := 0

	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{Val: (sum + carry) % 10}
		p = p.Next
		carry = (carry + sum) / 10

	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}

	return head.Next
}

func addTwoNumbersv2(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	p := head
	carry := 0

	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{Val: (sum + carry) % 10}
		carry = (sum + carry) / 10
		p = p.Next
	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}

	return head.Next
}

func addTwoNumbersv1(l1 *ListNode, l2 *ListNode) *ListNode {

	head := &ListNode{}
	p := head
	carry := 0
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		node := &ListNode{Val: (sum + carry) % 10}
		carry = (sum + carry) / 10
		p.Next = node
		p = p.Next

	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}

	return head.Next
}

/*
// Neetcode version creates a 'next' node first, and caculate currnet value and put the carry number into the future node
// he uses two pointers tech. in this problem, I should study
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummy = new(ListNode)
	var l3 **ListNode = &dummy
	var carry int
	list1 := l1
	list2 := l2

	for (list1 != nil) || (list2 != nil) {
		l3 = &((*l3).Next)
		*l3 = new(ListNode)
		var sum int

		if list1 != nil {
			sum += list1.Val
			list1 = list1.Next
		}

		if list2 != nil {
			sum += list2.Val
			list2 = list2.Next
		}

		(*l3).Val = (sum + carry) % 10
		carry = (sum + carry) / 10
	}

	if carry > 0 {
		l3 = &((*l3).Next)
		*l3 = new(ListNode)
		(*l3).Val = carry
	}

	return dummy.Next
}
*/

/* simple solution but can't pass the big number case
func addTwoNumbersv1(l1 *ListNode, l2 *ListNode) *ListNode {

	reverse := func(l *ListNode) *ListNode {
		if l == nil {
			return nil
		}
		ptr := l
		var prev *ListNode
		for ptr != nil {
			tmp := ptr.Next
			ptr.Next = prev
			prev = ptr
			ptr = tmp
		}
		return prev
	}

	n1 := 0
	p := reverse(l1)
	for p != nil {
		n1 *= 10
		n1 += p.Val
		p = p.Next
	}

	n2 := 0
	p = reverse(l2)
	for p != nil {
		n2 *= 10
		n2 += p.Val
		p = p.Next
	}

	totalNum := n1 + n2
	if totalNum == 0 {
		return &ListNode{Val: 0}
	}

	head := &ListNode{}
	p = head
	for totalNum != 0 {
		v := totalNum % 10
		p.Next = &ListNode{Val: v}
		totalNum /= 10
		p = p.Next
	}

	return head.Next
}
*/
