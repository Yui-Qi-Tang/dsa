package b75

import (
	"testing"
)

func TestFindMinRotatedSortedArray(t *testing.T) {

	testfuncs := []func([]int) int{
		findMinRotatedSortedArrayv4,
		findMinRotatedSortedArrayv3,
		findMinRotatedSortedArrayv2,
		findMinRotatedSortedArrayv1,
	}

	testcases := []struct {
		in   []int
		want int
	}{
		{in: []int{3, 4, 5, 1, 2}, want: 1},
		{in: []int{4, 5, 6, 7, 0, 1, 2}, want: 0},
		{in: []int{11, 13, 15, 17}, want: 11},
	}

	for i, f := range testfuncs {
		t.Logf("test functions %d", i)
		for j, tt := range testcases {
			ans := f(tt.in)
			if ans != tt.want {
				t.Fatalf("case[%d]: it should be %d, but got %d", j, tt.want, ans)
			}
		}
		t.Logf("function %d is passed", i)
	}
}
