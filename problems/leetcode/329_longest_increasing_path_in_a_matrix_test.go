package b75

import "testing"

func TestLongestIncreasingPath(t *testing.T) {
	testfuncs := []func([][]int) int{
		longestIncreasingPathv12,
		longestIncreasingPathv11,
		longestIncreasingPathv10,
		longestIncreasingPathv9,
		longestIncreasingPathv8,
		longestIncreasingPathv7,
		longestIncreasingPathv6,
		longestIncreasingPathv5,
		longestIncreasingPathv4,
		longestIncreasingPathv3,
		longestIncreasingPathv2,
		longestIncreasingPathv1,
	}

	testcases := []struct {
		matrix [][]int
		want   int
	}{
		{
			matrix: [][]int{
				{9, 9, 4},
				{6, 6, 8},
				{2, 1, 1},
			},
			want: 4,
		},
		{
			matrix: [][]int{
				{3, 4, 5},
				{3, 2, 6},
				{2, 2, 1},
			},
			want: 4,
		},
		{
			matrix: [][]int{
				{1},
			},
			want: 1,
		},
	}

	for i, f := range testfuncs {
		t.Logf("start testing function[%d]...", i)
		for j, tt := range testcases {
			ans := f(tt.matrix)
			if ans != tt.want {
				t.Fatalf("case[%d]: it should be %d. but got %d", j, tt.want, ans)
			}
		}
		t.Logf("... passed")
	}
}
