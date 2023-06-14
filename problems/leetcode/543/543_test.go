package b75

import (
	"fmt"
	"testing"
)

func TestBuildTree(t *testing.T) {

	tree := buildTree([]int{1, 2, 3, 4, 5})
	dump(tree)

}

func TestDiameterOfBinaryTree(t *testing.T) {
	testfuncs := []func(*TreeNode) int{
		diameterOfBinaryTreev30,
		diameterOfBinaryTreev29,
		diameterOfBinaryTreev28,
		diameterOfBinaryTreev27,
		diameterOfBinaryTreev26,
		diameterOfBinaryTreev25,
		diameterOfBinaryTreev24,
		diameterOfBinaryTreev23,
		diameterOfBinaryTreev22,
		diameterOfBinaryTreev21,
		diameterOfBinaryTreev20,
		diameterOfBinaryTreev19,
		diameterOfBinaryTreev18,
		diameterOfBinaryTreev17,
		diameterOfBinaryTreev16,
		diameterOfBinaryTreev15,
		diameterOfBinaryTreev14,
		diameterOfBinaryTreev13,
		diameterOfBinaryTreev12,
		diameterOfBinaryTreev11,
		diameterOfBinaryTreev10,
		diameterOfBinaryTreev9,
		diameterOfBinaryTreev8,
		diameterOfBinaryTreev7,
		diameterOfBinaryTreev6,
		diameterOfBinaryTreev5,
		diameterOfBinaryTreev4,
		diameterOfBinaryTreev3,
		diameterOfBinaryTreev2,
		diameterOfBinaryTreev1,
	}

	testcases := []struct {
		in   []int
		want int
	}{
		{
			in:   []int{1, 2, 3, 4, 5},
			want: 3,
		},
		{
			in:   []int{1, 2},
			want: 1,
		},
		{
			in:   []int{},
			want: 0,
		},
	}

	for i, f := range testfuncs {
		t.Logf("test functino %d", i)
		for j, tt := range testcases {
			tree := buildTree(tt.in)
			ans := f(tree)
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
