package b75

import (
	"fmt"
	"testing"
)

func TestMaxProduct(t *testing.T) {
	testfuncs := []func([]int) int{
		maxProductv12,
		maxProductv11,
		maxProductv10,
		maxProductv9,
		maxProductv8,
		maxProductv7,
		maxProductv6,
		maxProductv5,
		maxProductv4,
		maxProductv3,
		maxProductv2,
		maxProductv1,
	}

	testcases := []struct {
		in   []int
		want int
	}{
		{
			in:   []int{-2, 0, -1},
			want: 0,
		},
		{
			in:   []int{2, 3, -2, 4},
			want: 6,
		},
		{
			in:   []int{2, 0, -1},
			want: 2,
		},
	}

	for i, f := range testfuncs {
		t.Run(fmt.Sprintf("test function %d", i), func(t *testing.T) {
			for j, tt := range testcases {
				ans := f(tt.in)
				if ans != tt.want {
					t.Fatalf("case[%d]: it should be %d, but got %d", j, tt.want, ans)
				}
			}
		})
	}

}
