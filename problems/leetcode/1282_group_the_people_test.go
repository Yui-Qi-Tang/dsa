package b75

import (
	"reflect"
	"testing"
)

func TestGroupthepeople(t *testing.T) {
	testfuncs := []func([]int) [][]int{
		groupthepeoplev12,
		groupthepeoplev11,
		groupthepeoplev10,
		groupthepeoplev9,
		groupthepeoplev8,
		groupthepeoplev7,
		groupthepeoplev6,
		groupthepeoplev5,
		groupthepeoplev4,
		groupthepeoplev3,
		groupthepeoplev2,
		groupthepeoplev1,
	}

	testcases := []struct {
		in   []int
		want [][]int
	}{
		{
			in: []int{3, 3, 3, 3, 3, 1, 3},
			want: [][]int{
				{0, 1, 2},
				{3, 4, 6},
				{5},
			},
		},
		{
			in: []int{2, 1, 3, 3, 3, 2},
			want: [][]int{
				{0, 5},
				{1},
				{2, 3, 4},
			},
		},
	}

	for i, f := range testfuncs {
		t.Logf("test function[%d]...", i)

		for j, tt := range testcases {
			ans := f(tt.in)

			for _, want := range tt.want {
				hit := false
				for _, a := range ans {
					if reflect.DeepEqual(want, a) {
						hit = true
						break
					}
				}
				if !hit {
					t.Fatalf("=> case[%d]: it should be %v, but got %v", j, tt.want, ans)
				}
			}

		}

		t.Logf("... Passed")
	}
}
