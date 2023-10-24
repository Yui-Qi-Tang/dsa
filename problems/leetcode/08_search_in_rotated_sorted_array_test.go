package b75

import "testing"

func TestSearchInRotatedArray(t *testing.T) {

	testfuncs := []func([]int, int) int{
		SearchInRotatedArrayv22,
		SearchInRotatedArrayv21,
		SearchInRotatedArrayv20,
		SearchInRotatedArrayv19,
		SearchInRotatedArrayv18,
		SearchInRotatedArrayv17,
		SearchInRotatedArrayv16,
		SearchInRotatedArrayv15,
		SearchInRotatedArrayv14,
		SearchInRotatedArrayv13,
		SearchInRotatedArrayv12,
		SearchInRotatedArrayv11,
		SearchInRotatedArrayv10,
		SearchInRotatedArrayv9,
		SearchInRotatedArrayv8,
		SearchInRotatedArrayv7,
		SearchInRotatedArrayv6,
		SearchInRotatedArrayv5,
		SearchInRotatedArrayv4,
		SearchInRotatedArrayv3,
		SearchInRotatedArrayv2,
		SearchInRotatedArrayv1,
	}

	testcases := []struct {
		in     []int
		target int
		want   int
	}{
		{
			in:     []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
		{
			in:     []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			want:   -1, // 3 does not exist in the array
		},
		{
			in:     []int{4, 5, 6, 7, 0, 1, 2},
			target: 2,
			want:   6,
		},
		{
			in:     []int{1},
			target: 0,
			want:   -1,
		},
	}

	for i, f := range testfuncs {
		t.Logf("test function[%d]\n", i)
		for j, tt := range testcases {
			ans := f(tt.in, tt.target)
			if ans != tt.want {
				t.Fatalf("case[%d]: it should be %d, but got %d", j, tt.want, ans)
			}
		}
		t.Logf("function[%d] is passed\n", i)
	}

}
