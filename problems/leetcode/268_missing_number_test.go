package b75

import "testing"

func TestMissingNumber(t *testing.T) {

	testFunc := []func(nums []int) int{
		missingNumberv12,
		missingNumberv11,
		missingNumberv10,
		missingNumberv9,
		missingNumberv8,
		missingNumberv7,
		missingNumberv6,
		missingNumberv5,
		missingNumberv4,
		missingNumberv3,
		missingNumberv2,
		missingNumberv1,
	}

	testcases := []struct {
		in   []int
		want int
	}{
		{in: []int{3, 0, 1}, want: 2},
		{in: []int{0, 1}, want: 2},
		{in: []int{9, 6, 4, 2, 3, 5, 7, 0, 1}, want: 8},
		{in: []int{1, 2}, want: 0},
	}

	for i, f := range testFunc {

		t.Logf("test function %d...", i)
		for j, tt := range testcases {
			ans := f(tt.in)
			if ans != tt.want {
				t.Fatalf("=>case[%d]: it should be %d, but got %d", j, tt.want, ans)
			}
		}
		t.Logf("... function[%d] is passed", i)

	}

}
