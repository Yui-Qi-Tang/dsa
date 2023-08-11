package b75

import "testing"

func TestMaximumSubarray(t *testing.T) {

	testfuncs := []func([]int) int{
		MaximumSubarrayv38,
		MaximumSubarrayv37,
		MaximumSubarrayv36,
		MaximumSubarrayv35,
		MaximumSubarrayv34,
		MaximumSubarrayv33,
		MaximumSubarrayv32,
		MaximumSubarrayv31,
		MaximumSubarrayv30,
		MaximumSubarrayv29,
		MaximumSubarrayv28,
		MaximumSubarrayv27,
		MaximumSubarrayv26,
		MaximumSubarrayv25,
		MaximumSubarrayv24,
		MaximumSubarrayv23,
		MaximumSubarrayv22,
		MaximumSubarrayv21,
		MaximumSubarrayv20,
		MaximumSubarrayv19,
		MaximumSubarrayv18,
		MaximumSubarrayv17,
		MaximumSubarrayv16,
		MaximumSubarrayv15,
		MaximumSubarrayv14,
		MaximumSubarrayv13,
		MaximumSubarrayv12,
		MaximumSubarrayv11,
		MaximumSubarrayv10,
		MaximumSubarrayv9,
		MaximumSubarrayv8,
		MaximumSubarrayv7,
		MaximumSubarrayv6,
		MaximumSubarrayv5,
		MaximumSubarrayv4,
		MaximumSubarrayv3,
		MaximumSubarrayv2,
		MaximumSubarrayv1,
	}

	testcases := []struct {
		in   []int
		want int
	}{
		{in: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, want: 6},
		{in: []int{1}, want: 1},
		{in: []int{5, 4, -1, 7, 8}, want: 23},
		{in: []int{1, 2, 3, -2, 5}, want: 9},
		{in: []int{-1, -2, -3, -4}, want: -1},
	}

	for i, f := range testfuncs {
		t.Logf("test function %d", i)
		for j, tt := range testcases {
			ans := f(tt.in)
			if ans != tt.want {
				t.Fatalf("case[%d]: it should be %d, but got %d", j, tt.want, ans)
			}
		}

		t.Logf("function[%d] is passed", i)
	}

}
