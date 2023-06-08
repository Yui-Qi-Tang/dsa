package b75

import (
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	testfuncs := []func(*TreeNode) []int{
		rightSideViewv20,
		rightSideViewv19,
		rightSideViewv18,
		rightSideViewv17,
		rightSideViewv16,
		rightSideViewv15,
		rightSideViewv14,
		rightSideViewv13,
		rightSideViewv12,
		rightSideViewv11,
		rightSideViewv10,
		rightSideViewv9,
		rightSideViewv8,
		rightSideViewv7,
		rightSideViewv6,
		rightSideViewv5,
		rightSideViewv4,
		rightSideViewv3,
		rightSideViewv2,
		rightSideViewv1,
		rightSideView,
	}

	testcases := []struct {
		in   []int
		want []int
	}{
		{
			in:   []int{1, 2, 3, null, 5, null, 4},
			want: []int{1, 3, 4},
		},
		{
			in:   []int{1, null, 3},
			want: []int{1, 3},
		},
		{
			in:   []int{},
			want: []int{},
		},
		{
			in:   []int{1, 2, 3, null, 5, null, 4},
			want: []int{1, 3, 4},
		},
		{
			in:   []int{1, 2, null, 3, null},
			want: []int{1, 2, 3},
		},
		{
			in:   []int{1, 2, 3, 4},
			want: []int{1, 3, 4},
		},
	}

	for i, f := range testfuncs {
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
