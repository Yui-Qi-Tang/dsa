package b75

import "testing"

func TestAddTwoNumbers(t *testing.T) {

	testfuncs := []func(*ListNode, *ListNode) *ListNode{
		addTwoNumbersv29,
		addTwoNumbersv28,
		addTwoNumbersv27,
		addTwoNumbersv26,
		addTwoNumbersv25,
		addTwoNumbersv24,
		addTwoNumbersv23,
		addTwoNumbersv22,
		addTwoNumbersv21,
		addTwoNumbersv20,
		addTwoNumbersv19,
		addTwoNumbersv18,
		addTwoNumbersv17,
		addTwoNumbersv16,
		addTwoNumbersv15,
		addTwoNumbersv14,
		addTwoNumbersv13,
		addTwoNumbersv12,
		addTwoNumbersv11,
		addTwoNumbersv10,
		addTwoNumbersv9,
		addTwoNumbersv8,
		addTwoNumbersv7,
		addTwoNumbersv6,
		addTwoNumbersv5,
		addTwoNumbersv4,
		addTwoNumbersv3,
		addTwoNumbersv2,
		addTwoNumbersv1,
	}

	testcases := []struct {
		l1, l2 []int
		want   []int
	}{
		{
			l1:   []int{2, 4, 3},
			l2:   []int{5, 6, 4},
			want: []int{7, 0, 8},
		},
		{
			l1:   []int{0},
			l2:   []int{0},
			want: []int{0},
		},
		{
			l1:   []int{2, 4, 9},
			l2:   []int{5, 6, 4, 9},
			want: []int{7, 0, 4, 0, 1},
		},
		{ // should consider the big number
			l1:   []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			l2:   []int{5, 6, 4},
			want: []int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		},
	}

	createList := func(nums ...int) *ListNode {
		if len(nums) == 0 {
			return nil
		}

		list := &ListNode{Val: nums[0]}
		p := list

		for i := 1; i < len(nums); i++ {
			p.Next = &ListNode{Val: nums[i]}
			p = p.Next
		}
		return list
	}

	var same func(*ListNode, *ListNode) bool
	same = func(l1 *ListNode, l2 *ListNode) bool {
		if l1 == nil && l2 == nil {
			return true
		} else if l1 == nil && l2 != nil {
			return false
		} else if l1 != nil && l2 == nil {
			return false
		} else if l1.Val != l2.Val {
			return false
		}

		return same(l1.Next, l2.Next)
	}

	dump := func(l *ListNode) {
		p := l
		for p != nil {
			t.Log(p.Val)
			p = p.Next
		}
	}

	for i, f := range testfuncs {
		t.Logf("test function %d", i)
		for j, tt := range testcases {
			l1, l2, want := createList(tt.l1...), createList(tt.l2...), createList(tt.want...)
			ans := f(l1, l2)
			if !same(ans, want) {
				t.Log("dump your answer => ")
				dump(ans)
				t.Log("dump expected answer => ")
				dump(want)
				t.Fatalf("case[%d]: ans list does not equal to want list", j)
			}
		}
		t.Logf("function[%d] is passed", i)
	}

}
