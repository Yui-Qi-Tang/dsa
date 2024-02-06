package b75

import "testing"

func TestMinPathSum(t *testing.T) {

	testFuncs := []func(grid [][]int) int{
		minPathSumv31,
		minPathSumv30,
		minPathSumv29,
		minPathSumv28,
		minPathSumv27,
		minPathSumv26,
		minPathSumv25,
		minPathSumv24,
		minPathSumv23,
		minPathSumv22,
		minPathSumv21,
		minPathSumv20,
		minPathSumv19,
		minPathSumv18,
		minPathSumv17,
		minPathSumv16,
		minPathSumv15,
		minPathSumv14,
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
			// HINT: copy 2-d slice directly,
			// the copy func. will just copy the pointer
			cp := make([][]int, len(tt.in))
			for i := range cp {
				cp[i] = make([]int, len(tt.in[i]))
				copy(cp[i], tt.in[i])
			}
			ans := f(cp)
			if ans != tt.want {
				t.Fatalf("=> case[%d]: it should be %d, but got %d", j, tt.want, ans)
			}
		}
		t.Logf("... function[%d] is passed", i)
	}

}
