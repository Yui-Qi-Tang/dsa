package b75

import "testing"

func TestKthSmallest(t *testing.T) {

	testfuncs := []func(*TreeNode, int) int{
		kthSmallestv16,
		kthSmallestv15,
		kthSmallestv14,
		kthSmallestv13,
		kthSmallestv12,
		kthSmallestv11,
		kthSmallestv10,
		kthSmallestv9,
		kthSmallestv8,
		kthSmallestv7,
		kthSmallestv6,
		kthSmallestv5,
		kthSmallestv4,
		kthSmallestv3,
		kthSmallestv2,
		kthSmallestv1,
		kthSmallestRec,
		kthSmallest,
	}

	testcases := []struct {
		in   []int
		k    int
		want int
	}{
		{
			in: []int{3, 1, 4, null, 2}, k: 1,
			want: 1,
		},
		{
			in: []int{5, 3, 6, 2, 4, null, null, 1}, k: 3,
			want: 3,
		},
	}

	for i, f := range testfuncs {
		t.Logf("test function %d", i)
		for j, tt := range testcases {
			ans := f(buildTree(tt.in), tt.k)
			if ans != tt.want {
				t.Fatalf("case[%d]: it should be %d, but got %d", j, tt.want, ans)
			}
		}
		t.Logf("function[%d] is passed", i)
	}

}
