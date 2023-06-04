package b75

import "testing"

func TestCopyRandomList(t *testing.T) {

	testfuncs := []func(*Node) *Node{
		copyRandomListv30,
		copyRandomListv29,
		copyRandomListv28,
		copyRandomListv27,
		copyRandomListv26,
		copyRandomListv25,
		copyRandomListv24,
		copyRandomListv23,
		copyRandomListv22,
		copyRandomListv21,
		copyRandomListv20,
		copyRandomListv19,
		copyRandomListv18,
		copyRandomListv17,
		copyRandomListv16,
		copyRandomListv15,
		copyRandomListv14,
		copyRandomListv13,
		copyRandomListv12,
		copyRandomListv11,
		copyRandomListv10,
		copyRandomListv9,
		copyRandomListv8,
		copyRandomListv7,
		copyRandomListv6,
		copyRandomListv5,
		copyRandomListv4,
		copyRandomListv3,
		copyRandomListv2,
		copyRandomListv1,
	}

	testcases := []struct {
		in [][2]int
	}{
		{
			in: [][2]int{{7, -1}, {13, 0}, {11, 2}},
		},

		{
			in: [][2]int{{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0}},
		},
		{
			in: [][2]int{{1, 1}, {2, 1}},
		},
		{
			in: [][2]int{{3, -1}, {3, 0}, {3, -1}},
		},
	}

	for i, f := range testfuncs {
		t.Logf("test function %d", i)
		for j, tt := range testcases {
			in := createList(tt.in)
			ans := f(in)
			//dump(t, ans)

			if !same(in, ans) {
				t.Fatalf("case[%d]: it should be equal, but not", j)
			}
		}
		t.Logf("function[%d] is passed", i)
	}

}

func dump(t *testing.T, head *Node) {
	if head == nil {
		t.Log("empty lise")
		return
	}
	t.Log("==> dump nodes of the list")
	p := head
	for p != nil {
		t.Logf("=> node: %p, value: %d, next: %p, random %p", p, p.Val, p.Next, p.Random)
		p = p.Next
	}
	t.Log("")
}

func same(l1, l2 *Node) bool {
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
