package b75

import (
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	in := []int{4, 2, 7, 1, 3, 6, 9}

	root := buildTree(in)
	postOrder(root)

	in2 := []int{2, 1, 3}

	root2 := buildTree(in2)
	postOrder(root2)
	/*
	   t.Log("-=======-")
	   in2 := []int{1, 2, 3}
	   root2 := buildTree(in2)
	   postOrder(root2)

	   t.Log("-=======-")
	   in3 := []int{3, 2, 1}
	   root3 := buildTree(in3)
	   postOrder(root3)
	*/
}

func TestInvertTree(t *testing.T) {

	testfuncs := []func(*TreeNode) *TreeNode{
		invertTreev30,
		invertTreev29,
		invertTreev28,
		invertTreev27,
		invertTreev26,
		invertTreev25,
		invertTreev24,
		invertTreev23,
		invertTreev22,
		invertTreev21,
		invertTreev20,
		invertTreev19,
		invertTreev18,
		invertTreev17,
		invertTreev16,
		invertTreev15,
		invertTreev14,
		invertTreev13,
		invertTreev12,
		invertTreev11,
		invertTreev10,
		invertTreev9,
		invertTreev8,
		invertTreev7,
		invertTreev6,
		invertTreev5,
		invertTreev4,
		invertTreev3,
		invertTreev2,
		invertTreev1,
	}

	testcases := []struct {
		in   []int
		want []int
	}{
		{
			in:   []int{4, 2, 7, 1, 3, 6, 9},
			want: []int{4, 7, 2, 9, 6, 3, 1},
		},
		{
			in:   []int{2, 1, 3},
			want: []int{2, 3, 1},
		},
		{
			in:   []int{},
			want: []int{},
		},
	}

	for i, f := range testfuncs {
		t.Logf("test function %d", i)
		for j, tt := range testcases {
			root := f(buildTree(tt.in))
			ans := make([]int, 0, len(tt.want))

			var appendAns func(r *TreeNode)
			appendAns = func(r *TreeNode) {

				if r == nil {
					return
				}

				// append root
				if root == r {
					ans = append(ans, r.Val)
				}

				if r.Left != nil {
					ans = append(ans, r.Left.Val)
				}

				if r.Right != nil {
					ans = append(ans, r.Right.Val)
				}
				appendAns(r.Left)
				appendAns(r.Right)

			}
			appendAns(root)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Fatalf("case[%d]: it should be %v, but got %v", j, tt.want, ans)
			}

		}
		t.Logf("function[%d] is passed", i)
	}

}
