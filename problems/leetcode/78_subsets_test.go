package b75

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	testfunc := []func([]int) [][]int{
		subsetsv11,
		subsetsv10,
		subsetsv9,
		subsetsv8,
		subsetsv7,
		subsetsv6,
		subsetsv5,
		subsetsv4,
		subsetsv3,
		subsetsv2,
		subsetsv1,
	}

	testcases := []struct {
		in   []int
		want [][]int
	}{
		{
			in: []int{1, 2, 3},
			want: [][]int{
				{1, 2, 3}, {1, 2}, {1, 3},
				{1}, {2, 3}, {2}, {3}, {},
			},
		},
		{
			in:   []int{0},
			want: [][]int{{0}, {}},
		},
	}

	for i, f := range testfunc {
		t.Run(fmt.Sprintf("test function %d", i), func(t *testing.T) {
			for j, tt := range testcases {
				ans := f(tt.in)
				if !reflect.DeepEqual(ans, tt.want) {
					t.Fatalf("case[%d]: it should be: %v, but got: %v", j, tt.want, ans)
				}
			}
		})
	}
}
