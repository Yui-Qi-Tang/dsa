package b75

import "testing"

func TestBuildTree(t *testing.T) {

	testfuncs := []func([]int, []int) *TreeNode{
		buildTreev26,
		buildTreev25,
		buildTreev24,
		buildTreev23,
		buildTreev22,
		buildTreev21,
		buildTreev20,
		buildTreev19,
		buildTreev18,
		buildTreev17,
		buildTreev16,
		buildTreev15,
		buildTreev14,
		buildTreev13,
		buildTreev12,
		buildTreev11,
		buildTreev10,
		buildTreev9,
		buildTreev8,
		buildTreev7,
		buildTreev6,
		buildTreev5,
		buildTreev4,
		buildTreev3,
		buildTreev2,
		buildTreev1,
		buildTree,
	}

	testcases := []struct {
		prorder, inorder []int
		want             []int
	}{
		{
			prorder: []int{3, 9, 20, 15, 7},
			inorder: []int{9, 3, 15, 20, 7},
			want:    []int{3, 9, 20, null, null, 15, 7},
		},

		{
			prorder: []int{-1},
			inorder: []int{-1},
			want:    []int{-1},
		},
	}

	var same func(p, q *TreeNode) bool
	same = func(p *TreeNode, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}

		if p.Val != q.Val {
			return false
		}

		return same(p.Left, q.Left) && same(p.Right, q.Right)
	}

	for i, f := range testfuncs {
		t.Logf("test function %d", i)
		for j, tt := range testcases {
			want := build(tt.want)
			ans := f(tt.prorder, tt.inorder)

			if !same(want, ans) {
				t.Fatalf("case[%d]: it should be %v, but got %v", j, printPreorder(want), printPreorder(ans))
			}
		}

		t.Logf("function[%d] is passed", i)
	}

}
