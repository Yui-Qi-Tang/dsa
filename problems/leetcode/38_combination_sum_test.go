package b75

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	testfuncs := []func([]int, int) [][]int{
		combinationSumv6,
		combinationSumv5,
		combinationSumv4,
		combinationSumv3,
		combinationSumv2,
		combinationSumv1,
	}

	testcases := []struct {
		in     []int // candidates
		target int
		want   [][]int
	}{
		{
			in:     []int{2, 3, 6, 7},
			target: 7,
			want: [][]int{
				{2, 2, 3}, {7},
			},
		},
		{
			in:     []int{2, 3, 5},
			target: 8,
			want: [][]int{
				{2, 2, 2, 2}, {2, 3, 3}, {3, 5},
			},
		},
	}

	for i, f := range testfuncs {
		t.Run(fmt.Sprintf("test function %d", i), func(t *testing.T) {
			for j, tt := range testcases {
				ans := f(tt.in, tt.target)
				if !reflect.DeepEqual(ans, tt.want) {
					t.Fatalf("case[%d]: it should be %v, but got %v", j, tt.want, ans)
				}
			}

		})
	}
}
