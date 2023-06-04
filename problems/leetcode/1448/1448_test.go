package b75

import "testing"

func TestGoodNodes(t *testing.T) {
	tesrfuncs := []func(*TreeNode) int{
		goodNodesv14,
		goodNodesv13,
		goodNodesv12,
		goodNodesv11,
		goodNodesv10,
		goodNodesv9,
		goodNodesv8,
		goodNodesv7,
		goodNodesv6,
		goodNodesv5,
		goodNodesv4,
		goodNodesv3,
		goodNodesv2,
		goodNodesv1,
		goodNodes,
	}

	testcases := []struct {
		in   []int
		want int
	}{
		{
			in:   []int{3, 1, 4, 3, null, 1, 5},
			want: 4,
		},
		{
			in:   []int{3, 3, null, 4, 2},
			want: 3,
		},
		{
			in:   []int{1},
			want: 1,
		},
	}

	for i, f := range tesrfuncs {
		t.Logf("test function %d", i)
		for j, tt := range testcases {
			ans := f(buildTree(tt.in))
			if ans != tt.want {
				t.Fatalf("case[%d]: it should be %d, but got %d", j, tt.want, ans)
			}
		}
		t.Logf("function[%d] is passed", i)
	}
}
