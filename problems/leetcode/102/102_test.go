package b75

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {

	testfunc := []func(*TreeNode) [][]int{
		levelOrderv19,
		levelOrderv18,
		levelOrderv17,
		levelOrderv16,
		levelOrderv15,
		levelOrderv14,
		levelOrderv13,
		levelOrderv12,
		levelOrderv11,
		levelOrderv10,
		levelOrderv9,
		levelOrderv8,
		levelOrderv7,
		levelOrderv6,
		levelOrderv5,
		levelOrderv4,
		levelOrderv3,
		levelOrderv2,
		levelOrderv1,
		levelOrder,
	}

	testcases := []struct {
		in   []int
		want [][]int
	}{
		{
			in: []int{3, 9, 20, null, null, 15, 7},
			want: [][]int{
				{3},
				{9, 20},
				{15, 7},
			},
		},
		{
			in: []int{1},
			want: [][]int{
				{1},
			},
		},
		{
			in:   []int{},
			want: [][]int{},
		},
	}

	for i, f := range testfunc {
		t.Logf("test function %d", i)
		for j, tt := range testcases {
			ans := f(buildTree(tt.in))
			if !reflect.DeepEqual(ans, tt.want) {
				t.Fatalf("case[%d]: it should be %v, but got %v", j, tt.want, ans)
			}
		}
		t.Logf("function[%d] is passed", i)
	}

}
