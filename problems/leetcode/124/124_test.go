package b75

import "testing"

func TestMaxPathSum(t *testing.T) {
	testfuncs := []func(*TreeNode) int{
		maxPathSumv10,
		maxPathSumv9,
		maxPathSumv8,
		maxPathSumv7,
		maxPathSumv6,
		maxPathSumv5,
		maxPathSumv4,
		maxPathSumv3,
		maxPathSumv2,
		maxPathSumv1,
		maxPathSum,
	}

	testcases := []struct {
		in   []int
		want int
	}{
		{
			in:   []int{1, 2, 3},
			want: 6,
		},
		{
			in:   []int{-10, 9, 20, null, null, 15, 7},
			want: 42,
		},
	}

	for i, f := range testfuncs {
		t.Logf("test function %d", i)
		for j, tt := range testcases {
			tree := buildTree(tt.in)
			// inorderPrint(tree) // debug
			ans := f(tree)
			if ans != tt.want {
				t.Fatalf("case[%d]: it should be %d, but got %d", j, tt.want, ans)
			}
		}
		t.Logf("function[%d] is passed", i)
	}

}