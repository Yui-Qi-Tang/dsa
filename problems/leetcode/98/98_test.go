package b75

import "testing"

func TestIsValidBST(t *testing.T) {

	testfuncs := []func(*TreeNode) bool{
		isValidBSTv28,
		isValidBSTv27,
		isValidBSTv26,
		isValidBSTv25,
		isValidBSTv24,
		isValidBSTv23,
		isValidBSTv22,
		isValidBSTv21,
		isValidBSTv20,
		isValidBSTv19,
		isValidBSTv18,
		isValidBSTv17,
		isValidBSTv16,
		isValidBSTv15,
		isValidBSTv14,
		isValidBSTv13,
		isValidBSTv12,
		isValidBSTv11,
		isValidBSTv10,
		isValidBSTv9,
		isValidBSTv8,
		isValidBSTv7,
		isValidBSTv6,
		isValidBSTv5,
		isValidBSTv4,
		isValidBSTv3,
		isValidBSTv2,
		isValidBSTv1,
		isValidBST,
	}

	testcases := []struct {
		in   []int
		want bool
	}{
		{
			in:   []int{2, 1, 3},
			want: true,
		},
		{
			in:   []int{5, 1, 4, null, null, 3, 6},
			want: false,
		},
		{
			in:   []int{2, 2, 2},
			want: false,
		},
		{
			in:   []int{1, null, 1},
			want: false,
		},
		{ // we should consider the min, max
			in:   []int{5, 4, 6, null, null, 3, 7},
			want: false,
		},
		{
			in:   []int{3, 1, 5, 0, 2, 4, 6},
			want: true,
		},
	}

	for i, f := range testfuncs {
		t.Logf("test function %d", i)
		for j, tt := range testcases {
			ans := f(buildTree(tt.in))
			if ans != tt.want {
				t.Fatalf("cases[%d]: it should be %v, but got %v", j, tt.want, ans)
			}
		}
		t.Logf("function[%d] is passed", i)
	}

}
