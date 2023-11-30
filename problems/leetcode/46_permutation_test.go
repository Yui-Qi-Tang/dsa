package b75

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPermutation(t *testing.T) {
	testfuncs := []func([]int) [][]int{
		permutev4,
		permutev3,
		permutev2,
		permutev1,
	}

	testcases := []struct {
		in   []int
		want [][]int
	}{
		{
			in: []int{1, 2, 3},
			want: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
		{
			in: []int{0, 1},
			want: [][]int{
				{0, 1},
				{1, 0},
			},
		},
		{
			in: []int{5, 4, 6, 2},
			want: [][]int{
				{5, 4, 6, 2}, {5, 4, 2, 6}, {5, 6, 4, 2}, {5, 6, 2, 4}, {5, 2, 4, 6}, {5, 2, 6, 4}, {4, 5, 6, 2}, {4, 5, 2, 6},
				{4, 6, 5, 2}, {4, 6, 2, 5}, {4, 2, 5, 6}, {4, 2, 6, 5}, {6, 5, 4, 2}, {6, 5, 2, 4}, {6, 4, 5, 2}, {6, 4, 2, 5},
				{6, 2, 5, 4}, {6, 2, 4, 5}, {2, 5, 4, 6}, {2, 5, 6, 4}, {2, 4, 5, 6}, {2, 4, 6, 5}, {2, 6, 5, 4}, {2, 6, 4, 5},
			},
		},
	}

	for i, f := range testfuncs {
		t.Run(fmt.Sprintf("test function %d", i), func(t *testing.T) {
			for j, tt := range testcases {
				ans := f(tt.in)
				lackElem := lacks(tt.want, ans)
				if len(lackElem) != 0 {
					t.Fatalf("case[%d]: it should be %v, but got %v, diff :%v", j, tt.want, ans, lackElem)
				}
			}
		})
	}
}

func lacks(expected, value [][]int) [][]int {

	lacks := make([][]int, 0)
	for _, want := range expected {
		found := false
		for _, ans := range value {
			if reflect.DeepEqual(want, ans) {
				found = true
			}
		}

		if !found {
			lacks = append(lacks, want)
		}
	}

	return lacks
}
