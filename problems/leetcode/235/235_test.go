package b75

import "testing"

func TestBuildTreeAndFind(t *testing.T) {
	testcases := []struct {
		in []int
	}{
		{
			in: []int{6, 2, 8, 0, 4, 7, 9, null, null, 3, 5},
		},
	}

	for _, tt := range testcases {
		root := buildTree(tt.in)
		for _, val := range tt.in {
			if val != null {
				if node := findIteration(root, val); node == nil {
					t.Fatal(val, "dosen't not exist in tree")
				}
				if node := findRecusively(root, val); node == nil {
					t.Fatal(val, "dosen't not exist in tree")
				}
			}
		}
	}

}

func TestLowestCommonAncestor(t *testing.T) {

	testfuncs := []func(root, p, q *TreeNode) *TreeNode{
		lowestCommonAncestorv16,
		lowestCommonAncestorv15,
		lowestCommonAncestorv14,
		lowestCommonAncestorv13,
		lowestCommonAncestorv12,
		lowestCommonAncestorv11,
		lowestCommonAncestorv10,
		lowestCommonAncestorv9,
		lowestCommonAncestorv8,
		lowestCommonAncestorv7,
		lowestCommonAncestorv6,
		lowestCommonAncestorv5,
		lowestCommonAncestorv4,
		lowestCommonAncestorv3,
		lowestCommonAncestorv2,
		lowestCommonAncestorv1,
		lowestCommonAncestor,
	}

	testcases := []struct {
		in   []int
		p, q int
		want int
	}{
		{
			in:   []int{6, 2, 8, 0, 4, 7, 9, null, null, 3, 5},
			p:    2,
			q:    8,
			want: 6,
		},
		{
			in:   []int{6, 2, 8, 0, 4, 7, 9, null, null, 3, 5},
			p:    2,
			q:    4,
			want: 2,
		},
		{
			in:   []int{2, 1},
			p:    2,
			q:    1,
			want: 2,
		},
	}

	for i, f := range testfuncs {
		t.Logf("test function %d", i)
		for j, tt := range testcases {
			root := buildTree(tt.in)
			p := findIteration(root, tt.p)
			q := findIteration(root, tt.q)
			want := findIteration(root, tt.want)
			ans := f(root, p, q)
			if ans != want {
				t.Logf("case[%d]: it should be %p, but got %p", j, want, ans)
			}
		}
		t.Logf("function[%d] is passed", i)
	}

}
