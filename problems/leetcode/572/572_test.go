package b75

import "testing"

func TestIsSubtree(t *testing.T) {

	testfuncs := []func(p, q *TreeNode) bool{
		isSubtreev22,
		isSubtreev21,
		isSubtreev20,
		isSubTreev19,
		isSubtreev18,
		isSubtreev17,
		isSubtreev16,
		isSubtreev15,
		isSubtreev14,
		isSubtreev13,
		isSubtreev12,
		isSubtreev11,
		isSubtreev10,
		isSubtreev9,
		isSubtreev8,
		isSubtreev7,
		isSubtreev6,
		isSubtreev5,
		isSubtreev4,
		isSubtreev3,
		isSubtreev2,
		isSubtreev1,
	}

	testcases := []struct {
		p, q []int
		want bool
	}{
		{
			p:    []int{3, 4, 5, 1, 2},
			q:    []int{4, 1, 2},
			want: true,
		},
		{
			p:    []int{3, 4, 5, 1, 2, null, null, null, null, 0},
			q:    []int{4, 1, 2},
			want: true,
		},
	}

	for i, f := range testfuncs {
		t.Logf("test function %d", i)
		for j, tt := range testcases {
			ans := f(buildTree(tt.p), buildTree(tt.p))
			if ans != tt.want {
				t.Fatalf("case[%d]: it should be %v, but got %v", j, tt.want, ans)
			}
		}
		t.Logf("function[%d] is passed", i)
	}

}
