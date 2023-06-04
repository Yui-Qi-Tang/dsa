package b75

import (
	"testing"
)

func TestHasCycle(t *testing.T) {
	testfuncs := []func(head *ListNode) bool{
		hasCyclev29,
		hasCyclev28,
		hasCyclev27,
		hasCyclev26,
		hasCyclev25,
		hasCyclev24,
		hasCyclev23,
		hasCyclev22,
		hasCyclev21,
		hasCyclev20,
		hasCyclev19,
		hasCyclev18,
		hasCyclev17,
		hasCyclev16,
		hasCyclev15,
		hasCyclev14,
		hasCyclev13,
		hasCyclev12,
		hasCyclev11,
		hasCyclev10,
		hasCyclev9,
		hasCyclev8,
		hasCyclev7,
		hasCyclev6,
		hasCyclev5,
		hasCyclev4,
		hasCyclev3,
		hasCyclev2,
		hasCyclev1,
	}

	testcases := []struct {
		in   []int
		pos  int
		want bool
	}{

		{
			in:   []int{3, 2, 0, -4},
			pos:  1,
			want: true,
		},
		{
			in:   []int{1, 2},
			pos:  0,
			want: true,
		},
		{
			in:   []int{1, 2},
			pos:  -1,
			want: false,
		},
	}

	buildList := func(in []int) *ListNode {
		if len(in) == 0 {
			return nil
		}

		result := &ListNode{}
		p := result
		for _, v := range in {
			p.Next = &ListNode{Val: v}
			p = p.Next
		}

		return result.Next
	}

	findNode := func(head *ListNode, pos int) *ListNode {
		if pos == 0 {
			return head
		}

		p := head

		for pos > 0 {
			p = p.Next
			pos--
		}
		return p
	}

	findTail := func(head *ListNode) *ListNode {

		p := head

		for p.Next != nil {
			p = p.Next
		}

		return p
	}

	for i, f := range testfuncs {
		t.Logf("test function %d...", i)
		for j, tt := range testcases {
			head := buildList(tt.in)

			if tt.pos > -1 {
				node := findNode(head, tt.pos)
				tail := findTail(head)
				tail.Next = node
			}
			ans := f(head)
			if ans != tt.want {
				t.Fatalf("=> case[%d]: it should be %v, but got %v", j, tt.want, ans)
			}
		}
		t.Logf("...function[%d] is passed", i)
	}

}
