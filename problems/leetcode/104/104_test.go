package b75

import (
	"fmt"
	"testing"
)

func TestBuildTree(t *testing.T) {

	testfuncs := []func(*TreeNode) int{
		maxDepthv28,
		maxDepthv27,
		maxDepthv26,
		maxDepthv25,
		maxDepthv24,
		maxDepthv23,
		maxDepthv22,
		maxDepthv21,
		maxDepthv20,
		maxDepthv19,
		maxDepthv18,
		maxDepthv17,
		maxDepthv16,
		maxDepthv15,
		maxDepthv14,
		maxDepthv13,
		maxDepthv12,
		maxDepthv11,
		maxDepthv10,
		maxDepthv9,
		maxDepthv8,
		maxDepthv7,
		maxDepthv6,
		maxDepthv5,
		maxDepthv4,
		maxDepthv3,
		maxDepthv2,
		maxDepthv1,
	}

	testcases := []struct {
		in   *TreeNode
		want int
	}{
		{
			in:   buildTree([]int{3, 9, 20, -101, -101, 15, 7}), // -101 denotes null
			want: 3,
		},
		{
			in:   buildTree([]int{1, -101, 2}), // -101 denotes null
			want: 2,
		},
	}

	for i, f := range testfuncs {
		t.Logf("test function %d", i)
		for j, tt := range testcases {
			ans := f(tt.in)
			if ans != tt.want {
				t.Fatalf("case[%d]: it should be %d, but got %d", j, tt.want, ans)
			}
		}
		t.Logf("function[%d] is passed", i)
	}

}

func dump(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Val)
	dump(root.Left)
	dump(root.Right)
}
