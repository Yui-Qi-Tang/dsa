package b75

import "testing"

func TestMinPathSum(t *testing.T) {

	testFuncs := []func(grid [][]int) int{
		minPathSumv13,
		minPathSumv12,
		minPathSumv11,
		minPathSumv10,
		minPathSumv9,
		minPathSumv8,
		minPathSumv7,
		minPathSumv6,
		minPathSumv5,
		minPathSumv4,
		minPathSumv3,
		minPathSumv2,
		minPathSumv1,
	}

	testcases := []struct {
		in   [][]int
		want int
	}{
		{
			in: [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
			want: 7,
		},

		{
			in: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			want: 12,
		},
	}

	for i, f := range testFuncs {
		t.Logf("test function %d...", i)
		for j, tt := range testcases {
			ans := f(tt.in)
			if ans != tt.want {
				t.Fatalf("=> case[%d]: it should be %d, but got %d", j, tt.want, ans)
			}
		}
		t.Logf("... function[%d] is passed", i)
	}

}
