package b75

import (
	"fmt"
	"testing"
)

func TestMinPairSum(t *testing.T) {

	testfuncs := []func([]int) int{
		minPairSumv7,
		minPairSumv6,
		minPairSumv5,
		minPairSumv4,
		minPairSumv3,
		minPairSumv2,
		minPairSumv1,
	}

	testcases := []struct {
		in   []int
		want int
	}{
		{
			in:   []int{3, 5, 2, 3},
			want: 7,
		},
		{
			in:   []int{3, 5, 4, 2, 4, 6},
			want: 8,
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
