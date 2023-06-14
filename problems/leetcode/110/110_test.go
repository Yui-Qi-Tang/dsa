package b75

import "testing"

func TestBuildTree(t *testing.T) {
	root := buildTree([]int{3, 9, 20, null, null, 15, 7})
	dump(root)
}

func TestIsBalanced(t *testing.T) {

	testfuncs := []func(*TreeNode) bool{
		isBalancedv29,
		isBalancedv28,
		isBalancedv27,
		isBalancedv26,
		isBalancedv25,
		isBalancedv24,
		isBalancedv23,
		isBalancedv22,
		isBalancedv21,
		isBalancedv20,
		isBalancedv19,
		isBalancedv18,
		isBalancedv17,
		isBalancedv16,
		isBalancedv15,
		isBalancedv14,
		isBalancedv13,
		isBalancedv12,
		isBalancedv11,
		isBalancedv10,
		isBalancedv9,
		isBalancedv8,
		isBalancedv7,
		isBalancedv6,
		isBalancedv5,
		isBalancedv4,
		isBalancedv3,
		isBalancedv2,
		isBalancedv1,
	}

	testcases := []struct {
		in   []int
		want bool
	}{
		{
			in:   []int{3, 9, 20, null, null, 15, 7},
			want: true,
		},
		{
			in:   []int{1, 2, 2, 3, 3, null, null, 4, 4},
			want: false,
		},
		{
			in:   []int{},
			want: true,
		},
		{
			in:   []int{1, 2, 2, 3, null, null, 3, 4, null, null, 4},
			want: false,
		},
	}

	for i, f := range testfuncs {
		t.Logf("test function %d", i)
		for j, tt := range testcases {
			ans := f(buildTree(tt.in))
			if ans != tt.want {
				t.Fatalf("case[%d]: it should be %v, but got %v", j, tt.want, ans)
			}
		}
		t.Logf("function %d is passed", i)
	}
}
