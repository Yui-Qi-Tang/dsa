package b75

import "testing"

func TestIsSameTree(t *testing.T) {

	testfunc := []func(*TreeNode, *TreeNode) bool{
		isSameTreev27,
		isSameTreev26,
		isSameTreev25,
		isSameTreev24,
		isSameTreev23,
		isSameTreev22,
		isSameTreev21,
		isSameTreev20,
		isSameTreev19,
		isSameTreev18,
		isSameTreev17,
		isSameTreev16,
		isSameTreev15,
		isSameTreev14,
		isSameTreev13,
		isSameTreev12,
		isSameTreev11,
		isSameTreev10,
		isSameTreev9,
		isSameTreev8,
		isSameTreev7,
		isSameTreev6,
		isSameTreev5,
		isSameTreev4,
		isSameTreev3,
		isSameTreev2,
		isSameTreev1,
	}

	testcases := []struct {
		p, q []int
		want bool
	}{
		{
			p:    []int{1, 2, 3},
			q:    []int{1, 2, 3},
			want: true,
		},
		{
			p: []int{1, 2},
			q: []int{1, null, 2},
		},
		{
			p: []int{1, 2, 1},
			q: []int{1, 1, 2},
		},
	}

	for i, f := range testfunc {
		t.Logf("test function %d", i)
		for j, tt := range testcases {
			p := buildTree(tt.p)
			q := buildTree(tt.q)
			ans := f(p, q)
			if ans != tt.want {
				t.Fatalf("case[%d]: it should be %v, but got %v", j, tt.want, ans)
			}
		}
		t.Logf("function[%d] is passed", i)
	}

}
