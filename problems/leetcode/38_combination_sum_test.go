package b75

import (
	"fmt"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	testfuncs := []func([]int, int) [][]int{
		combinationSumv43,
		combinationSumv42,
		combinationSumv41,
		combinationSumv40,
		combinationSumv39,
		combinationSumv38,
		combinationSumv37,
		combinationSumv36,
		combinationSumv35,
		combinationSumv34,
		combinationSumv33,
		combinationSumv32,
		combinationSumv31,
		combinationSumv30,
		combinationSumv29,
		combinationSumv28,
		combinationSumv27,
		combinationSumv26,
		combinationSumv25,
		combinationSumv24,
		combinationSumv23,
		combinationSumv22,
		combinationSumv21,
		combinationSumv20,
		combinationSumv19,
		combinationSumv18,
		combinationSumv17,
		combinationSumv16,
		combinationSumv15,
		combinationSumv14,
		combinationSumv13,
		combinationSumv12,
		combinationSumv11,
		combinationSumv10,
		combinationSumv9,
		combinationSumv8,
		combinationSumv7,
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
				lackElem := lacks(tt.want, ans)
				if len(lackElem) != 0 {
					t.Fatalf("case[%d]: it should be %v, but got %v, diff :%v", j, tt.want, ans, lackElem)
				}
			}
		})
	}
}
